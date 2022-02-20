package space

import (
	"github.com/arbori/population.git/population/rule"
)

type Point struct {
	X int
	Y int
}

func (p *Point) Add(point *Point) {
	p.X = point.X
	p.Y = point.Y
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
	neighborhoodMotion NeighborhoodMotion
	Cells              [][]float32
}

func MakeEnvironment(X int, Y int, neighborhoodMotion NeighborhoodMotion) Environment {
	environment := Environment{
		X:                  X,
		Y:                  Y,
		neighborhoodMotion: neighborhoodMotion,
		Cells:              make([][]float32, X),
	}

	for x := 0; x < environment.X; x += 1 {
		environment.Cells[x] = make([]float32, Y)
	}

	return environment
}

func (s *Environment) Neighborhood(x int, y int) []float32 {
	neighborhood := make([]float32, s.neighborhoodMotion.Size)

	if len(s.neighborhoodMotion.Motion) == 0 || len(s.neighborhoodMotion.Motion) != s.neighborhoodMotion.Size || len(s.neighborhoodMotion.Motion[0]) != s.neighborhoodMotion.Dimention {
		return neighborhood
	}

	var i int
	var j int

	for index := 0; index < s.neighborhoodMotion.Size; index += 1 {
		j = s.neighborhoodMotion.Motion[index][0] + x
		i = s.neighborhoodMotion.Motion[index][1] + y

		if j < 0 {
			j = s.X + j
		} else if j >= s.X {
			j = j - s.X
		}

		if i < 0 {
			i = s.Y + i
		} else if i >= s.Y {
			i = i - s.Y
		}

		neighborhood[index] = s.Cells[j][i]
	}

	return neighborhood
}

func (s *Environment) GetNewPosition(position *Point, directionChoosed int) Point {
	result := Point{
		X: s.neighborhoodMotion.Motion[directionChoosed][0] + position.X,
		Y: s.neighborhoodMotion.Motion[directionChoosed][1] + position.Y,
	}

	if result.X < 0 {
		result.X = s.X + result.X
	} else if result.X >= s.X {
		result.X = result.X - s.X
	}

	if result.Y < 0 {
		result.Y = s.Y + result.Y
	} else if result.Y >= s.Y {
		result.Y = result.Y - s.Y
	}

	return result
}

func (s *Environment) ApplyRule(r rule.Rule) {
	for y := 0; y < s.Y; y += 1 {
		for x := 0; x < s.X; x += 1 {
			s.Cells[x][y] = r.Transition(s.Neighborhood(x, y))
		}
	}
}
