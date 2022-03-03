package space

import (
	"github.com/arbori/population.git/population/rule"
)

type Point struct {
	X int
	Y int
}

func (p *Point) Add(point *Point) {
	p.X += point.X
	p.Y += point.Y
}

type NeighborhoodMotion struct {
	Size      int
	Dimention int
	Motion    [][]int
}

func MakeNeighborhoodMotion(size int, dimention int) NeighborhoodMotion {
	result := NeighborhoodMotion{
		Size:      size,
		Dimention: dimention,
		Motion:    make([][]int, size),
	}

	for s := 0; s < size; s += 1 {
		result.Motion[s] = make([]int, dimention)
	}

	return result
}

type Environment struct {
	X                  int
	Y                  int
	Cells              [][]float32
	mirror             [][]float32
	neighborhoodMotion NeighborhoodMotion
}

func MakeEnvironment(X int, Y int, neighborhoodMotion *NeighborhoodMotion) Environment {
	environment := Environment{
		X:                  X,
		Y:                  Y,
		Cells:              make([][]float32, X),
		mirror:             make([][]float32, X),
		neighborhoodMotion: *neighborhoodMotion,
	}

	for x := 0; x < environment.X; x += 1 {
		environment.Cells[x] = make([]float32, Y)
		environment.mirror[x] = make([]float32, Y)
	}

	return environment
}

func (e *Environment) Neighborhood(x int, y int) []float32 {
	neighborhood := make([]float32, e.neighborhoodMotion.Size)

	if len(e.neighborhoodMotion.Motion) == 0 || len(e.neighborhoodMotion.Motion) != e.neighborhoodMotion.Size || len(e.neighborhoodMotion.Motion[0]) != e.neighborhoodMotion.Dimention {
		return neighborhood
	}

	var i int
	var j int

	for index := 0; index < e.neighborhoodMotion.Size; index += 1 {
		j = e.neighborhoodMotion.Motion[index][0] + x
		i = e.neighborhoodMotion.Motion[index][1] + y

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

		neighborhood[index] = e.Cells[j][i]
	}

	return neighborhood
}

func (e *Environment) GetNewPosition(position *Point, directionChoosed int) Point {
	result := Point{
		X: e.neighborhoodMotion.Motion[directionChoosed][0] + position.X,
		Y: e.neighborhoodMotion.Motion[directionChoosed][1] + position.Y,
	}

	if result.X < 0 {
		result.X = e.X + result.X
	} else if result.X >= e.X {
		result.X = result.X - e.X
	}

	if result.Y < 0 {
		result.Y = e.Y + result.Y
	} else if result.Y >= e.Y {
		result.Y = result.Y - e.Y
	}

	return result
}

func (e *Environment) ApplyRule(r rule.Rule) {
	for y := 0; y < e.Y; y += 1 {
		for x := 0; x < e.X; x += 1 {
			e.mirror[x][y] = r.Transition(e.Neighborhood(x, y))
		}
	}

	for y := 0; y < e.Y; y += 1 {
		for x := 0; x < e.X; x += 1 {
			e.Cells[x][y] = e.mirror[x][y]
		}
	}
}
