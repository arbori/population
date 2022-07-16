package main

import (
	"fmt"
	"math/rand"

	"github.com/arbori/population.git/population/agent"
	"github.com/arbori/population.git/population/space"
)

func interationRuleDefinition(a1 *agent.MobileAgent, a2 *agent.MobileAgent, contribuitionProbability float32, exchangeRate float32) {
	if a1.Position[0] != a2.Position[0] || a1.Position[1] != a2.Position[1] {
		return
	}

	selfishness := rand.Float32()

	if a1.Resources > a2.Resources {
		if selfishness > contribuitionProbability {
			a1.Resources = a1.Resources + exchangeRate*a2.Resources
			a2.Resources = (1 - exchangeRate) * a2.Resources
		} else {
			a2.Resources = a2.Resources + exchangeRate*a1.Resources
			a1.Resources = (1 - exchangeRate) * a1.Resources
		}
	} else {
		if selfishness > contribuitionProbability {
			a2.Resources = a2.Resources + exchangeRate*a1.Resources
			a1.Resources = (1 - exchangeRate) * a1.Resources
		} else {
			a1.Resources = a1.Resources + exchangeRate*a2.Resources
			a2.Resources = (1 - exchangeRate) * a2.Resources
		}
	}
}

func motionRuleDefinition(environment *space.Environment, position *space.Point) space.Point {
	neighborhood := environment.NeighborhoodValues((*position)[0], (*position)[1])

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

var vonNeumannNeighborhoodMotion = space.NeighborhoodMotion{
	Size: 5,
	Directions: []space.Point{
		[]int{0, 0},
		[]int{-1, 0},
		[]int{0, 1},
		[]int{1, 0},
		[]int{0, -1},
	},
}

func constructMooreNeighborhoodMotion() space.NeighborhoodMotion {
	motion := space.MakeNeighborhoodMotion(9, 2)

	motion.Directions[0] = []int{0, 0}
	motion.Directions[1] = []int{-1, 0}
	motion.Directions[2] = []int{-1, 1}
	motion.Directions[3] = []int{0, 1}
	motion.Directions[4] = []int{1, 1}
	motion.Directions[5] = []int{1, 0}
	motion.Directions[6] = []int{1, -1}
	motion.Directions[7] = []int{0, -1}
	motion.Directions[8] = []int{-1, -1}

	return motion
}

func constructEnvironment(motion *space.NeighborhoodMotion) space.Environment {
	environment := space.MakeEnvironment(5, 5, motion, .05)

	return environment
}

func exchangeBetweenAgents(agents []agent.MobileAgent, env *space.Environment) {

}

func showAgentsInEnvironment(environment space.Environment) {
	for y := 0; y < environment.Y; y += 1 {
		for x := 0; x < environment.X; x += 1 {
			if environment.Cells[x][y].Content != nil && environment.Cells[x][y].Content.(*agent.MobileAgent).Resources > .01 {
				fmt.Printf("%.2f\t", environment.Cells[x][y].Content.(*agent.MobileAgent).Resources) //"x\t")
			} else {
				fmt.Printf("_\t")
			}
		}

		fmt.Print("\n")
	}
}

func main() {
	/*
		motion := vonNeumannNeighborhoodMotion
		environment := constructEnvironment(&motion)

		spreadRule := rule.AverageRuleVonNeumann{}

		size := 3

		agents := make([]*agent.MobileAgent, size, size)
		agents[0] = &agent.MobileAgent{
			IsAvailable: true,
			Position:    []int {1, 1},
			Resources:   7.0,
			MotionRule:  motionRuleDefinition,
		}
		agents[1] = &agent.MobileAgent{
			IsAvailable: true,
			Position:    []int {3, 3},
			Resources:   5.0,
			MotionRule:  motionRuleDefinition,
		}
		agents[2] = &agent.MobileAgent{
			IsAvailable: true,
			Position:    []int {0, 4},
			Resources:   6.0,
			MotionRule:  motionRuleDefinition,
		}

		dead := make([]*agent.MobileAgent, size, size)

		for a := 0; a < len(agents); a += 1 {
			environment.Cells[agents[a].Position[0]][agents[a].Position[1]].Content = agents[a]
			environment.Cells[agents[a].Position[0]][agents[a].Position[1]].Value = agents[a].Resources
		}

		for t := 0; len(agents) > 0; t += 1 {
			fmt.Printf("%d\n", t)
			showAgentsInEnvironment(environment)
			fmt.Print("\n\n")

			environment.ApplyRule(spreadRule)

			for a := 0; a < len(agents); a += 1 {
				agents[a].Walk(&environment)

				agents[a].Resources *= (1 - environment.Inertia)

				if agents[a].Resources <= .01 {
					dead = append(dead, agents[a])
					agents = append(agents[:a], agents[a+1:]...)
				}
			}
		}

		for y := 0; y < environment.Y; y += 1 {
			for x := 0; x < environment.X; x += 1 {
				fmt.Printf("%.2f", environment.Cells[x][y].Value)
				fmt.Print("\t")
			}

			fmt.Print("\n")
		}

		fmt.Print("\n\n")
	*/
}
