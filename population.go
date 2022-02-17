package main

import "fmt"

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

type Rule interface {
	transition(neighborhood []float32) float32
}

type SpreadRuleVonNeumann struct {
	decay float32
}

func (rule SpreadRuleVonNeumann) transition(neighborhood []float32) float32 {
	var result float32 = 0.0

	size := len(neighborhood)

	if size != 5 {
		return result
	}

	result = (1-rule.decay)*neighborhood[4] + (rule.decay/float32(size))*(neighborhood[0]+neighborhood[1]+neighborhood[2]+neighborhood[3])

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

func (s *Space) ApplyRule(rule Rule, motion NeighborhoodMotion) {
	for y := 0; y < s.Y; y += 1 {
		for x := 0; x < s.X; x += 1 {
			s.Cells[x][y] = rule.transition(s.Neighborhood(motion, x, y))
		}
	}
}

func main() {
	fmt.Println("Hello, world.")

	space := MakeSpace(5, 5)
	motion := MakeNeighborhoodMotion(5, 2)

	space.Cells[3][2] = 1
	space.Cells[2][1] = 2
	space.Cells[1][2] = 3
	space.Cells[2][3] = 4
	space.Cells[2][2] = 5

	motion.Motion[0][0] = +1
	motion.Motion[0][1] = 0
	motion.Motion[1][0] = 0
	motion.Motion[1][1] = -1
	motion.Motion[2][0] = -1
	motion.Motion[2][1] = 0
	motion.Motion[3][0] = 0
	motion.Motion[3][1] = +1
	motion.Motion[4][0] = 0
	motion.Motion[4][0] = 0

	rule := SpreadRuleVonNeumann{
		decay: .15,
	}

	for t := 0; t < 5; t += 1 {
		for y := 0; y < space.Y; y += 1 {
			for x := 0; x < space.X; x += 1 {
				fmt.Print(space.Cells[x][y])
				fmt.Print("\t")
			}

			fmt.Print("\n")
		}

		fmt.Print("\n\n")

		space.ApplyRule(rule, motion)
	}
}
