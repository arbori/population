package agent

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/arbori/population.git/population/space"
)

func interationRuleDefinition(a1 *MobileAgent, a2 *MobileAgent, contribuitionProbability float32, exchangeRate float32) {
	if a1.Position.X[0] != a2.Position.X[0] && a1.Position.X[1] != a2.Position.X[1] {
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

func makeEnvironmentForTest() space.Environment {
	motion := space.MakeNeighborhoodMotion(5, 2)
	environment := space.MakeEnvironment(5, 5, &motion, .1)

	motion.Directions[0] = space.NewPoint(+1, 0)
	motion.Directions[1] = space.NewPoint(0, -1)
	motion.Directions[2] = space.NewPoint(-1, 0)
	motion.Directions[3] = space.NewPoint(0, +1)
	motion.Directions[4] = space.NewPoint(0, 0)

	environment.Cells[0][0].Value = 0.0030474192
	environment.Cells[1][0].Value = 0.047903188
	environment.Cells[2][0].Value = 0.27384293
	environment.Cells[3][0].Value = 0.06288487
	environment.Cells[4][0].Value = 0.006969613
	environment.Cells[0][1].Value = 0.042323112
	environment.Cells[1][1].Value = 0.46316734
	environment.Cells[2][1].Value = 1.8816189
	environment.Cells[3][1].Value = 0.5958209
	environment.Cells[4][1].Value = 0.0724816
	environment.Cells[0][2].Value = 0.1930319
	environment.Cells[1][2].Value = 1.4092783
	environment.Cells[2][2].Value = 3.0912015
	environment.Cells[3][2].Value = 2.2813718
	environment.Cells[4][2].Value = 0.3151865
	environment.Cells[0][3].Value = 0.034155533
	environment.Cells[1][3].Value = 0.28721535
	environment.Cells[2][3].Value = 0.9084532
	environment.Cells[3][3].Value = 0.4175537
	environment.Cells[4][3].Value = 0.06292355
	environment.Cells[0][4].Value = 0.103392154
	environment.Cells[1][4].Value = 0.03049182
	environment.Cells[2][4].Value = 0.10911071
	environment.Cells[3][4].Value = 0.044589132
	environment.Cells[4][4].Value = 0.0073037986

	for y := 0; y < environment.Y; y += 1 {
		for x := 0; x < environment.X; x += 1 {
			fmt.Print(environment.Cells[x][y])
			fmt.Print("\t")
		}

		fmt.Print("\n")
	}

	fmt.Print("\n\n")

	return environment
}

func TestAgentMotion(t *testing.T) {
	expectedOneStep := space.NewPoint(2, 1)
	expectedInside := space.NewPoint(0, 4)
	environment := makeEnvironmentForTest()
	agentPosition := space.NewPoint(1, 1)

	agent := MobileAgent{
		Position:    agentPosition,
		MotionRule:  motionRuleDefinition,
		IsAvailable: true,
	}

	agent.Walk(&environment)

	testResult := agent.Position.X[0] == expectedOneStep.X[0] && agent.Position.X[1] == expectedOneStep.X[1]

	fmt.Printf("The agent expected to be on site (%d, %d).\n", expectedOneStep.X[0], expectedOneStep.X[1])

	if !testResult {
		t.Fatalf("The agent moved to the wrong site! Expected site is (%d, %d), but the atual site is (%d, %d)",
			expectedOneStep.X[0], expectedOneStep.X[1], agent.Position.X[0], agent.Position.X[1])
	}

	agent.Position.X[0] = 0
	agent.Position.X[1] = 0
	agent.Walk(&environment)

	testResult = agent.Position.X[0] == expectedInside.X[0] && agent.Position.X[1] == expectedInside.X[1]

	fmt.Printf("The agent expected to be on site (%d, %d), inside the invironment.\n", expectedInside.X[0], expectedInside.X[1])

	if !testResult {
		t.Fatalf("The agent moved to the wrong site! Expected site is (%d, %d), but the atual site is (%d, %d)",
			expectedInside.X[0], expectedInside.X[1], agent.Position.X[0], agent.Position.X[1])
	}
}

func TestAgentInteration(t *testing.T) {
	exchange := Exchange{
		ContribuitionProbability: .5,
		ExchangeRate:             .25,
		interationRule:           interationRuleDefinition,
	}

	a1 := MobileAgent{
		Position:    space.NewPoint(1, 1),
		Foodstuffs:  100,
		MotionRule:  motionRuleDefinition,
		IsAvailable: true,
	}
	a1_expected := MobileAgent{
		Position:    space.NewPoint(1, 1),
		Foodstuffs:  100,
		MotionRule:  motionRuleDefinition,
		IsAvailable: true,
	}
	a2 := MobileAgent{
		Position:    space.NewPoint(0, 0),
		Foodstuffs:  50,
		MotionRule:  motionRuleDefinition,
		IsAvailable: true,
	}
	a2_expected := MobileAgent{
		Position:    space.NewPoint(0, 0),
		Foodstuffs:  50,
		MotionRule:  motionRuleDefinition,
		IsAvailable: true,
	}

	exchange.Interation(&a1, &a2)

	fmt.Printf("It is supose there is not interation between two agents in diferent place.\n")
	if a1.Foodstuffs != a1_expected.Foodstuffs || a2.Foodstuffs != a2_expected.Foodstuffs {
		t.Fatalf("There was interation between two agents in diferent place")
	}

	a2.Position.X[0] = a1.Position.X[0]
	a2.Position.X[1] = a1.Position.X[1]

	exchange.Interation(&a1, &a2)

	fmt.Printf("It is supose there is interation between two agents in diferent place.\n")
	if a1.Foodstuffs == a1_expected.Foodstuffs || a2.Foodstuffs == a2_expected.Foodstuffs {
		t.Fatalf("There was not interation between two agents in same place")
	}

	fmt.Printf("It is supose afeter interation between two agents, its foodstuffs amount has been change.\n")
	if a2.Foodstuffs == (a1.Foodstuffs+a1_expected.Foodstuffs)/exchange.ExchangeRate || a1.Foodstuffs == (a2.Foodstuffs+a2_expected.Foodstuffs)/exchange.ExchangeRate {
		t.Fatalf("Exchange Foodstuffs between two agents in the same site is ")
	}
}

func TestRetriveAgentsInTheNeighborhood(t *testing.T) {
	neighborhoodMotion := space.MakeNeighborhoodMotion(5, 2)
	neighborhoodMotion.Directions[0].X[0] = 0
	neighborhoodMotion.Directions[0].X[1] = 0
	neighborhoodMotion.Directions[1].X[0] = -1
	neighborhoodMotion.Directions[1].X[1] = 0
	neighborhoodMotion.Directions[2].X[0] = 0
	neighborhoodMotion.Directions[2].X[1] = 1
	neighborhoodMotion.Directions[3].X[0] = 1
	neighborhoodMotion.Directions[3].X[1] = 0
	neighborhoodMotion.Directions[4].X[0] = 0
	neighborhoodMotion.Directions[4].X[1] = -1

	a1 := MobileAgent{
		Position:    space.NewPoint(1, 1),
		Foodstuffs:  100,
		MotionRule:  motionRuleDefinition,
		IsAvailable: true,
	}
	a2 := MobileAgent{
		Position:    space.NewPoint(0, 1),
		Foodstuffs:  50,
		MotionRule:  motionRuleDefinition,
		IsAvailable: true,
	}

	env := space.MakeEnvironment(5, 5, &neighborhoodMotion, .1)
	env.Cells[a1.Position.X[0]][a1.Position.X[1]].Content = a1
	env.Cells[a2.Position.X[0]][a2.Position.X[1]].Content = a2
	env.Motion = neighborhoodMotion

	neighborhoodAgents := retriveAgentsInTheNeighborhood(&env, 1, 1)

	fmt.Printf("It is supose there are 2 agents in the neighborhood.\n")
	if len(neighborhoodAgents) != 2 {
		t.Fatalf("There are not only two agents in the neighborhood.")
	}
}

func TestAgentsExchangeResources(t *testing.T) {
	exchange := Exchange{
		ContribuitionProbability: .5,
		ExchangeRate:             .25,
		interationRule:           interationRuleDefinition,
	}

	neighborhoodMotion := space.MakeNeighborhoodMotion(5, 2)
	neighborhoodMotion.Directions[0].X[0] = 0
	neighborhoodMotion.Directions[0].X[1] = 0
	neighborhoodMotion.Directions[1].X[0] = -1
	neighborhoodMotion.Directions[1].X[1] = 0
	neighborhoodMotion.Directions[2].X[0] = 0
	neighborhoodMotion.Directions[2].X[1] = 1
	neighborhoodMotion.Directions[3].X[0] = 1
	neighborhoodMotion.Directions[3].X[1] = 0
	neighborhoodMotion.Directions[4].X[0] = 0
	neighborhoodMotion.Directions[4].X[1] = -1

	a1 := MobileAgent{
		Position:    space.NewPoint(1, 1),
		Foodstuffs:  100,
		MotionRule:  motionRuleDefinition,
		IsAvailable: true,
	}
	a1_expected := MobileAgent{
		Position:    space.NewPoint(1, 1),
		Foodstuffs:  100,
		MotionRule:  motionRuleDefinition,
		IsAvailable: true,
	}
	a2 := MobileAgent{
		Position:    space.NewPoint(0, 1),
		Foodstuffs:  50,
		MotionRule:  motionRuleDefinition,
		IsAvailable: true,
	}
	a2_expected := MobileAgent{
		Position:    space.NewPoint(0, 1),
		Foodstuffs:  50,
		MotionRule:  motionRuleDefinition,
		IsAvailable: true,
	}

	env := space.MakeEnvironment(5, 5, &neighborhoodMotion, .1)
	env.Cells[a1.Position.X[0]][a1.Position.X[1]].Content = a1
	env.Cells[a2.Position.X[0]][a2.Position.X[1]].Content = a2
	env.Motion = neighborhoodMotion

	AgentsExchangeResources(&env, &exchange, interationRuleDefinition)

	fmt.Printf("It is supose there is not interation between two agents in diferent place.\n")
	if a1.Foodstuffs != a1_expected.Foodstuffs || a2.Foodstuffs != a2_expected.Foodstuffs {
		t.Fatalf("There was interation between two agents in diferent place")
	}
}
