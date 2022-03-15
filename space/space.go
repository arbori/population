package space

import (
	"github.com/arbori/population.git/population/rule"
)

type PointError struct {
	msg string
}

func (pe PointError) Error() string {
	return pe.msg
}

type Point struct {
	X   []int
	Dim int
}

func NewPoint(x ...int) Point {
	result := Point{}

	result.Dim = len(x)
	result.X = make([]int, result.Dim)

	for i := 0; i < result.Dim; i += 1 {
		result.X[i] = x[i]
	}

	return result
}

func (p *Point) Assign(point *Point) error {
	if len(p.X) != len(point.X) {
		return PointError{msg: "The length of these two points are different."}
	}

	for i := 0; i < len(point.X); i += 1 {
		p.X[i] = point.X[i]
	}

	return nil
}

func (p *Point) Add(point *Point) error {
	if len(p.X) != len(point.X) {
		return PointError{msg: "The length of these two points are different."}
	}

	for i := 0; i < len(point.X); i += 1 {
		p.X[i] += point.X[i]
	}

	return nil
}

type NeighborhoodMotion struct {
	Size   int
	Motion []Point
}

func MakeNeighborhoodMotion(size int, dimention int) NeighborhoodMotion {
	result := NeighborhoodMotion{
		Size:   size,
		Motion: make([]Point, size),
	}

	for s := 0; s < size; s += 1 {
		result.Motion[s] = Point{
			X:   make([]int, dimention),
			Dim: dimention,
		}
	}

	return result
}

type Cell struct {
	Value   float32
	Content interface{}
}

type Environment struct {
	X                  int
	Y                  int
	Cells              [][]Cell
	mirror             [][]Cell
	neighborhoodMotion NeighborhoodMotion
	Inertia            float32
}

func MakeEnvironment(X int, Y int, neighborhoodMotion *NeighborhoodMotion, inertia float32) Environment {
	environment := Environment{
		X:                  X,
		Y:                  Y,
		Cells:              make([][]Cell, X),
		mirror:             make([][]Cell, X),
		neighborhoodMotion: *neighborhoodMotion,
		Inertia:            inertia,
	}

	for x := 0; x < environment.X; x += 1 {
		environment.Cells[x] = make([]Cell, Y)
		environment.mirror[x] = make([]Cell, Y)

		for y := 0; y < environment.Y; y += 1 {
			environment.Cells[x][y] = Cell{Value: 0.0, Content: nil}
			environment.mirror[x][y] = Cell{Value: 0.0, Content: nil}
		}
	}

	return environment
}

func (e *Environment) Neighborhood(x int, y int) []float32 {
	neighborhood := make([]float32, e.neighborhoodMotion.Size)

	if len(e.neighborhoodMotion.Motion) == 0 || len(e.neighborhoodMotion.Motion) != e.neighborhoodMotion.Size {
		return neighborhood
	}

	var i int
	var j int

	for index := 0; index < e.neighborhoodMotion.Size; index += 1 {
		j = e.neighborhoodMotion.Motion[index].X[0] + x
		i = e.neighborhoodMotion.Motion[index].X[1] + y

		if j < 0 {
			j = e.X + j
		} else if j >= e.X {
			j = j - e.X
		}

		if i < 0 {
			i = e.Y + i
		} else if i >= e.Y {
			i = i - e.Y
		}

		neighborhood[index] = e.Cells[j][i].Value
	}

	return neighborhood
}

func (e *Environment) GetNewPosition(position *Point, directionChoosed int) Point {
	result := NewPoint(
		e.neighborhoodMotion.Motion[directionChoosed].X[0]+position.X[0],
		e.neighborhoodMotion.Motion[directionChoosed].X[1]+position.X[1])

	if result.X[0] < 0 {
		result.X[0] = e.X + result.X[0]
	} else if result.X[0] >= e.X {
		result.X[0] = result.X[0] - e.X
	}

	if result.X[1] < 0 {
		result.X[1] = e.Y + result.X[1]
	} else if result.X[1] >= e.Y {
		result.X[1] = result.X[1] - e.Y
	}

	return result
}

func (e *Environment) ApplyRule(r rule.Rule) {
	for y := 0; y < e.Y; y += 1 {
		for x := 0; x < e.X; x += 1 {
			e.mirror[x][y].Value = r.Transition(e.Neighborhood(x, y))
		}
	}

	for y := 0; y < e.Y; y += 1 {
		for x := 0; x < e.X; x += 1 {
			e.Cells[x][y].Value = e.mirror[x][y].Value
		}
	}
}
