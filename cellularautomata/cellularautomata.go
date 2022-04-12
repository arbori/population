package cellularautomata

import (
	"math"

	"github.com/arbori/population.git/population/lattice"
	"github.com/arbori/population.git/population/rule"
)

type Cellularautomata struct {
	env       lattice.Lattice
	mirror    lattice.Lattice
	states    []interface{}
	motion    [][]int
	rule      rule.Rule
	dimention int
}

func New(states []interface{}, motion [][]int, rule rule.Rule, dim ...int) (Cellularautomata, error) {
	var err error
	var env lattice.Lattice
	var mirror lattice.Lattice

	env, err = lattice.New(dim...)

	if err != nil {
		return Cellularautomata{}, err
	}

	mirror, err = lattice.New(dim...)

	if err != nil {
		return Cellularautomata{}, err
	}

	return Cellularautomata{
		env:       env,
		mirror:    mirror,
		dimention: len(dim),
	}, nil
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

type StateTransitionFunction struct {
	numberOfStates            int
	numberOfNeighborhoodCells int
	ruleNumber                int
	transitionTable           [][]float32
}

func (st StateTransitionFunction) Transition(neighborhood []float32) float32 {
	for i := 0; i < len(st.transitionTable); i += 1 {
		match := true

		for j := 0; match && j < len(st.transitionTable[i])-1; j += 1 {
			match = match && neighborhood[j] == st.transitionTable[i][j]
		}

		if match {
			return st.transitionTable[i][len(st.transitionTable[i])-1]
		}
	}

	return 0
}

func MakeStateTransitionFunction(statesSize int, m int, ruleNumber int) rule.Rule {
	base := float64(statesSize)
	cells := float64(m)

	rule := StateTransitionFunction{
		numberOfStates:            statesSize,
		numberOfNeighborhoodCells: m,
		ruleNumber:                ruleNumber,
	}

	var transitionsSize int = int(math.Pow(base, cells))

	rule.transitionTable = make([][]float32, transitionsSize)

	ruleconverted := numBaseConv(ruleNumber, statesSize)
	convPos := 0
	convSize := len(ruleconverted)

	value := float32(ruleconverted[convPos])
	convPos += 1

	rule.transitionTable[0] = nextTransition(nil, cells, value, float32(statesSize))

	for i := 1; i < transitionsSize; i += 1 {
		if convPos < convSize {
			value = float32(ruleconverted[convPos])
			convPos += 1
		} else {
			value = 0
		}

		rule.transitionTable[i] = nextTransition(rule.transitionTable[i-1], cells, value, float32(statesSize))
	}

	return rule
}

func nextTransition(current []float32, cells float64, value float32, max float32) []float32 {
	m := int(cells)

	next := make([]float32, m+1)

	for j := 0; j < m; j += 1 {
		if current == nil {
			next[j] = 0
		} else {
			next[j] = current[j]
		}
	}

	for j := m - 1; current != nil && j >= 0; j -= 1 {
		next[j] = current[j] + 1

		if next[j] < max {
			break
		}

		next[j] = 0
	}

	next[m] = value

	return next
}

func numBaseConv(n int, base int) []int {
	result := make([]int, 0)

	for n >= base {
		result = append(result, n%base)
		n = n / base
	}

	if n > 0 {
		result = append(result, n)
	}

	return result
}
