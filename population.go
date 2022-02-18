package main

import (
	"fmt"

	"github.com/arbori/population.git/population/rule"
	"github.com/arbori/population.git/population/space"
)

func main() {
	fmt.Println("Hello, world.")

	s := space.MakeSpace(5, 5)
	motion := space.MakeNeighborhoodMotion(5, 2)

	s.Cells[3][2] = 1
	s.Cells[2][1] = 2
	s.Cells[1][2] = 3
	s.Cells[2][3] = 4
	s.Cells[2][2] = 5

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

	spreadRule := rule.SpreadRuleVonNeumann{
		Decay: .15,
	}

	for t := 0; t < 5; t += 1 {
		for y := 0; y < s.Y; y += 1 {
			for x := 0; x < s.X; x += 1 {
				fmt.Print(s.Cells[x][y])
				fmt.Print("\t")
			}

			fmt.Print("\n")
		}

		fmt.Print("\n\n")

		s.ApplyRule(spreadRule, motion)
	}
}
