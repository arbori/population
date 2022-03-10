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
	motion := space.MakeNeighborhoodMotion(9, 2)

	motion.Motion[0][0] = 0
	motion.Motion[0][1] = 0

	motion.Motion[1][0] = -1
	motion.Motion[1][1] = 0

	motion.Motion[2][0] = -1
	motion.Motion[2][1] = 1

	motion.Motion[3][0] = 0
	motion.Motion[3][1] = 1

	motion.Motion[4][0] = 1
	motion.Motion[4][1] = 1

	motion.Motion[5][0] = 1
	motion.Motion[5][1] = 0

	motion.Motion[6][0] = 1
	motion.Motion[6][1] = -1

	motion.Motion[7][0] = 0
	motion.Motion[7][1] = -1

	motion.Motion[8][0] = -1
	motion.Motion[8][1] = -1

	return motion
}

func constructEnvironment(motion *space.NeighborhoodMotion) space.Environment {
	environment := space.MakeEnvironment(5, 5, motion, .05)

	return environment
}

func main() {
	motion := constructNeighborhoodMotion()
	environment := constructEnvironment(&motion)

	spreadRule := rule.SpreadRuleMoore{
		Decay: .15,
	}

	agents := make([]agent.MobileAgent, 2)
	agents[0] = agent.MobileAgent{
		Position: space.Point{
			X: 1,
			Y: 1,
		},
		Foodstuffs: 7.0,
		MotionRule: motionRuleDefinition,
	}
	agents[1] = agent.MobileAgent{
		Position: space.Point{
			X: 3,
			Y: 3,
		},
		Foodstuffs: 5.0,
		MotionRule: motionRuleDefinition,
	}

	for a := 0; a < len(agents); a += 1 {
		environment.Cells[agents[a].Position.X][agents[a].Position.Y] = agents[a].Foodstuffs
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
		for a := 0; a < len(agents); a += 1 {
			//	agents[a].Walk(&environment)
			agents[a].Foodstuffs *= (1 - environment.Inertia)
			environment.Cells[agents[a].Position.X][agents[a].Position.Y] = agents[a].Foodstuffs
		}
	}

	for y := 0; y < environment.Y; y += 1 {
		for x := 0; x < environment.X; x += 1 {
			fmt.Printf("%.2f", environment.Cells[x][y])
			fmt.Print("\t")
		}

		fmt.Print("\n")
	}

	fmt.Print("\n\n")
}
