package main

import (
	"fmt"
	"math/rand"

	"github.com/arbori/population.git/population/agent"
	"github.com/arbori/population.git/population/rule"
	"github.com/arbori/population.git/population/space"
)

func interationRuleDefinition(a1 *agent.MobileAgent, a2 *agent.MobileAgent, contribuitionProbability float32, exchangeRate float32) {
	if a1.Position.X != a2.Position.X || a1.Position.Y != a2.Position.Y {
		return
	}

	selfishness := rand.Float32()

	if a1.Foodstuffs > a2.Foodstuffs {
		if selfishness > contribuitionProbability {
			a1.Foodstuffs = a1.Foodstuffs + exchangeRate*a2.Foodstuffs
			a2.Foodstuffs = (1 - exchangeRate) * a2.Foodstuffs
		} else {
			a2.Foodstuffs = a2.Foodstuffs + exchangeRate*a1.Foodstuffs
			a1.Foodstuffs = (1 - exchangeRate) * a1.Foodstuffs
		}
	} else {
		if selfishness > contribuitionProbability {
			a2.Foodstuffs = a2.Foodstuffs + exchangeRate*a1.Foodstuffs
			a1.Foodstuffs = (1 - exchangeRate) * a1.Foodstuffs
		} else {
			a1.Foodstuffs = a1.Foodstuffs + exchangeRate*a2.Foodstuffs
			a2.Foodstuffs = (1 - exchangeRate) * a2.Foodstuffs
		}
	}
}

func motionRuleDefinition(environment *space.Environment, position *space.Point) space.Point {
	neighborhood := environment.Neighborhood(position.X, position.Y)

	var maxPosition int = 0
	var maxValue float32 = neighborhood[maxPosition]

	for i := 1; i < len(neighborhood); i += 1 {
		if neighborhood[i] > maxValue {
			maxPosition = i
			maxValue = neighborhood[i]
		}
	}

	return environment.GetNewPosition(position, maxPosition)
}
func constructNeighborhoodMotion() space.NeighborhoodMotion {
	motion := space.MakeNeighborhoodMotion(5, 2)

	motion.Motion[0][0] = 0
	motion.Motion[0][1] = 0

	motion.Motion[1][0] = -1
	motion.Motion[1][1] = 0

	motion.Motion[2][0] = 0
	motion.Motion[2][1] = 1

	motion.Motion[3][0] = +1
	motion.Motion[3][1] = 0

	motion.Motion[4][0] = 0
	motion.Motion[4][1] = -1

	return motion
}

func constructEnvironment(motion *space.NeighborhoodMotion) space.Environment {
	environment := space.MakeEnvironment(5, 5, motion)

	environment.Cells[2][2] = 7

	return environment
}

func main() {
	fmt.Println("Hello, world.")

	motion := constructNeighborhoodMotion()
	environment := constructEnvironment(&motion)

	spreadRule := rule.SpreadRuleVonNeumann{
		Decay: .15,
	}

	for t := 0; t < 5; t += 1 {
		for y := 0; y < environment.Y; y += 1 {
			for x := 0; x < environment.X; x += 1 {
				fmt.Printf("%.2f", environment.Cells[x][y])
				fmt.Print("\t")
			}

			fmt.Print("\n")
		}

		fmt.Print("\n\n")

		environment.ApplyRule(spreadRule)
	}
	for y := 0; y < environment.Y; y += 1 {
		for x := 0; x < environment.X; x += 1 {
			fmt.Printf("%.2f", environment.Cells[x][y])
			fmt.Print("\t")
		}

		fmt.Print("\n")
	}

	fmt.Print("\n\n")

	agent := agent.MobileAgent{
		Position: space.Point{
			X: 1,
			Y: 1,
		},
		MotionRule: motionRuleDefinition,
	}

	agent.Walk(&environment)
}
