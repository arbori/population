package main

import (
	"fmt"
	"math/rand"

	"github.com/arbori/population.git/population/agent"
	"github.com/arbori/population.git/population/rule"
	"github.com/arbori/population.git/population/space"
)

func interationRuleDefinition(a1 *agent.MobileAgent, a2 *agent.MobileAgent, contribuitionProbability float32, exchangeRate float32) {
	if a1.Position.X[0] != a2.Position.X[0] || a1.Position.X[1] != a2.Position.X[1] {
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
	neighborhood := environment.Neighborhood(position.X[0], position.X[1])

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

func constructVonNeumannNeighborhoodMotion() space.NeighborhoodMotion {
	motion := space.MakeNeighborhoodMotion(5, 2)

	motion.Motion[0] = space.NewPoint(0, 0)
	motion.Motion[1] = space.NewPoint(-1, 0)
	motion.Motion[2] = space.NewPoint(0, 1)
	motion.Motion[3] = space.NewPoint(1, 0)
	motion.Motion[4] = space.NewPoint(0, -1)

	return motion
}

func constructMooreNeighborhoodMotion() space.NeighborhoodMotion {
	motion := space.MakeNeighborhoodMotion(9, 2)

	motion.Motion[0] = space.NewPoint(0, 0)
	motion.Motion[1] = space.NewPoint(-1, 0)
	motion.Motion[2] = space.NewPoint(-1, 1)
	motion.Motion[3] = space.NewPoint(0, 1)
	motion.Motion[4] = space.NewPoint(1, 1)
	motion.Motion[5] = space.NewPoint(1, 0)
	motion.Motion[6] = space.NewPoint(1, -1)
	motion.Motion[7] = space.NewPoint(0, -1)
	motion.Motion[8] = space.NewPoint(-1, -1)

	return motion
}

func constructEnvironment(motion *space.NeighborhoodMotion) space.Environment {
	environment := space.MakeEnvironment(5, 5, motion, .05)

	return environment
}

func main() {
	motion := constructVonNeumannNeighborhoodMotion()
	environment := constructEnvironment(&motion)

	spreadRule := rule.SpreadRuleVonNeumann{
		Decay: 1.0,
	}

	agents := make([]agent.MobileAgent, 2)
	agents[0] = agent.MobileAgent{
		Position:   space.NewPoint(1, 1),
		Foodstuffs: 7.0,
		MotionRule: motionRuleDefinition,
	}
	agents[1] = agent.MobileAgent{
		Position:   space.NewPoint(3, 3),
		Foodstuffs: 5.0,
		MotionRule: motionRuleDefinition,
	}

	for a := 0; a < len(agents); a += 1 {
		environment.Cells[agents[a].Position.X[0]][agents[a].Position.X[1]] = agents[a].Foodstuffs
	}

	for t := 0; t < 5; t -= 1 {
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
			agents[a].Walk(&environment)
			//agents[a].Foodstuffs *= (1 - environment.Inertia)
			environment.Cells[agents[a].Position.X[0]][agents[a].Position.X[1]] = agents[a].Foodstuffs
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
