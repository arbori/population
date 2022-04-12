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
	neighborhood := environment.NeighborhoodValues(position.X[0], position.X[1])

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
		space.NewPoint(0, 0),
		space.NewPoint(-1, 0),
		space.NewPoint(0, 1),
		space.NewPoint(1, 0),
		space.NewPoint(0, -1),
	},
}

func constructMooreNeighborhoodMotion() space.NeighborhoodMotion {
	motion := space.MakeNeighborhoodMotion(9, 2)

	motion.Directions[0] = space.NewPoint(0, 0)
	motion.Directions[1] = space.NewPoint(-1, 0)
	motion.Directions[2] = space.NewPoint(-1, 1)
	motion.Directions[3] = space.NewPoint(0, 1)
	motion.Directions[4] = space.NewPoint(1, 1)
	motion.Directions[5] = space.NewPoint(1, 0)
	motion.Directions[6] = space.NewPoint(1, -1)
	motion.Directions[7] = space.NewPoint(0, -1)
	motion.Directions[8] = space.NewPoint(-1, -1)

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
			if environment.Cells[x][y].Content != nil && environment.Cells[x][y].Content.(*agent.MobileAgent).Foodstuffs > .01 {
				fmt.Printf("%.2f\t", environment.Cells[x][y].Content.(*agent.MobileAgent).Foodstuffs) //"x\t")
			} else {
				fmt.Printf("_\t")
			}
		}

		fmt.Print("\n")
	}
}

func main() {
	motion := vonNeumannNeighborhoodMotion
	environment := constructEnvironment(&motion)

	spreadRule := rule.AverageRuleVonNeumann{}

	size := 3

	agents := make([]*agent.MobileAgent, size, size)
	agents[0] = &agent.MobileAgent{
		IsAvailable: true,
		Position:    space.NewPoint(1, 1),
		Foodstuffs:  7.0,
		MotionRule:  motionRuleDefinition,
	}
	agents[1] = &agent.MobileAgent{
		IsAvailable: true,
		Position:    space.NewPoint(3, 3),
		Foodstuffs:  5.0,
		MotionRule:  motionRuleDefinition,
	}
	agents[2] = &agent.MobileAgent{
		IsAvailable: true,
		Position:    space.NewPoint(0, 4),
		Foodstuffs:  6.0,
		MotionRule:  motionRuleDefinition,
	}

	dead := make([]*agent.MobileAgent, size, size)

	for a := 0; a < len(agents); a += 1 {
		environment.Cells[agents[a].Position.X[0]][agents[a].Position.X[1]].Content = agents[a]
		environment.Cells[agents[a].Position.X[0]][agents[a].Position.X[1]].Value = agents[a].Foodstuffs
	}

	for t := 0; len(agents) > 0; t += 1 {
		fmt.Printf("%d\n", t)
		showAgentsInEnvironment(environment)
		fmt.Print("\n\n")

		environment.ApplyRule(spreadRule)
		for a := 0; a < len(agents); a += 1 {
			agents[a].Walk(&environment)

			agents[a].Foodstuffs *= (1 - environment.Inertia)

			if agents[a].Foodstuffs <= .01 {
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
}
