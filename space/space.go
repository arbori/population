package space

import (
	"github.com/arbori/population.git/population/rule"
)

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

type Space struct {
	X     int
	Y     int
	Cells [][]float32
}

func MakeSpace(X int, Y int) Space {
	space := Space{
		X:     X,
		Y:     Y,
		Cells: make([][]float32, X),
	}

	for x := 0; x < space.X; x += 1 {
		space.Cells[x] = make([]float32, Y)
	}

	return space
}

func (s *Space) Neighborhood(motion NeighborhoodMotion, x int, y int) []float32 {
	neighborhood := make([]float32, motion.Size)

	if len(motion.Motion) == 0 || len(motion.Motion) != motion.Size || len(motion.Motion[0]) != motion.Dimention {
		return neighborhood
	}

	var i int
	var j int

	for index := 0; index < motion.Size; index += 1 {
		j = motion.Motion[index][0] + x
		i = motion.Motion[index][1] + y

		if j < 0 || j >= s.X || i < 0 || i >= s.Y {
			neighborhood[index] = 0.0
		} else {
			neighborhood[index] = s.Cells[j][i]
		}
	}

	return neighborhood
}

func (s *Space) ApplyRule(r rule.Rule, motion NeighborhoodMotion) {
	for y := 0; y < s.Y; y += 1 {
		for x := 0; x < s.X; x += 1 {
			s.Cells[x][y] = r.Transition(s.Neighborhood(motion, x, y))
		}
	}
}
