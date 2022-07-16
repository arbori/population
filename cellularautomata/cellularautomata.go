package cellularautomata

import (
	"math"

	"github.com/arbori/population.git/population/lattice"
	"github.com/arbori/population.git/population/rule"
	"github.com/arbori/population.git/population/space"
)

type Cellularautomata struct {
	env       lattice.Lattice
	mirror    lattice.Lattice
	states    []float32
	motion    [][]int
	rule      rule.Rule
	dimention int
}

func New(states []float32, motion [][]int, rule rule.Rule, dim ...int) (Cellularautomata, error) {
	var err error
	var env lattice.Lattice
	var mirror lattice.Lattice

	env, err = lattice.NewWithValue(float32(0), dim...)

	if err != nil {
		return Cellularautomata{}, err
	}

	mirror, err = lattice.NewWithValue(float32(0), dim...)

	if err != nil {
		return Cellularautomata{}, err
	}

	return Cellularautomata{
		env:       env,
		mirror:    mirror,
		states:    states,
		motion:    motion,
		rule:      rule,
		dimention: len(dim),
	}, nil
}

func (ca *Cellularautomata) NeighborhoodValues(X ...int) []float32 {
	size := len(ca.motion)
	dimention := len(ca.motion[0])

	if dimention != len(X) {
		return make([]float32, 0)
	}

	neighborhood := make([]float32, size)
	point := make([]int, dimention)

	for n := 0; n < size; n += 1 {
		for c := 0; c < dimention; c += 1 {
			point[c] = ca.motion[n][c] + X[c]

			if point[c] < 0 {
				point[c] = ca.env.Limits[c] + point[c]
			} else if point[c] >= ca.env.Limits[c] {
				point[c] = point[c] - ca.env.Limits[c]
			}
		}

		neighborhood[n] = ca.env.At(point...).(float32)
	}

	return neighborhood
}

func (ca *Cellularautomata) Get(x ...int) float32 {
	var result float32 = 0

	cell := ca.env.At(x...)

	if cell != nil {
		result = cell.(float32)
	}

	return result
}

func (ca *Cellularautomata) Set(value float32, x ...int) {
	for i := 0; ca.states[i] != value; i += 1 {
	}

	ca.env.Set(value, x...)
}

func (ca *Cellularautomata) Dimention() int {
	return ca.env.Dimention
}

func (ca *Cellularautomata) Limits() []int {
	return ca.env.Limits
}

func (ca *Cellularautomata) Evolve() {
	point := make([]int, ca.env.Dimention)
	position := 0
	overflowed := 1

	point[0] = -1

	for inc(&point, &ca.env.Limits, position, overflowed) {
		neighborhood := ca.NeighborhoodValues(point...)
		state := ca.rule.Transition(neighborhood)

		ca.mirror.Set(state, point...)
	}

	for i := 0; i < len(point); i += 1 {
		point[i] = 0
	}
	position = 0
	overflowed = 1

	point[0] = -1

	for inc(&point, &ca.env.Limits, position, overflowed) {
		ca.env.Set(ca.mirror.At(point...), point...)
	}
}

func inc(point *[]int, limits *space.Point, position int, overflowed int) bool {
	if position >= len(*point) || len(*point) != len(*limits) || (*point)[position] >= (*limits)[position] {
		return false
	}

	(*point)[position] += overflowed

	if (*point)[position] >= (*limits)[position] {
		(*point)[position] = 0
		overflowed = 1
		position += 1

		return inc(point, limits, position+1, overflowed)
	}

	return true
}

func NeighborhoodMotionVonNeumman(d int, r int) [][]int {
	var size int = d*r*(r+1) + 1
	var sum float64

	result := make([][]int, size)

	point := nextPointer(nil, d, r)

	for i := 0; i < size; {
		sum = 0

		for j := d - 1; j >= 0; j -= 1 {
			sum += math.Abs(float64(point[j]))
		}

		if sum <= 1 {
			result[i] = point
			i += 1
		}

		point = nextPointer(point, d, r)
	}

	return result
}

func NeighborhoodMotionMoore(d int, r int) [][]int {
	var size int = int(math.Pow(float64(2*r+1), float64(d)))

	result := make([][]int, size)

	result[0] = nextPointer(nil, d, r)

	for i := 1; i < size; i += 1 {
		result[i] = nextPointer(result[i-1], d, r)
	}

	return result
}

func nextPointer(current []int, d int, r int) []int {
	next := make([]int, d)

	for j := 0; j < d; j += 1 {
		if current == nil {
			next[j] = -r
		} else {
			next[j] = current[j]
		}
	}

	for j := d - 1; current != nil && j >= 0; j -= 1 {
		next[j] = current[j] + 1

		if next[j] <= r {
			break
		}

		next[j] = -r
	}

	return next
}
