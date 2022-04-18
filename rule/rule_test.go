package rule

import (
	"fmt"
	"testing"
)

func TestAverageRuleVonNeumannWithouNeighborhood(t *testing.T) {
	transitionValue := AverageRuleVonNeumann{}.Transition(nil)

	fmt.Println("The expected value for none neighborhood is 0.0")
	if transitionValue != 0 {
		t.Errorf("The expected value is 0.0, but the transition function return %f\n", transitionValue)
	}
}

func TestAverageRuleVonNeumannIncompliteNeighborhood(t *testing.T) {
	transitionValue := AverageRuleVonNeumann{}.Transition([]float32{2.3, 1.7})
	var averageExpected float32 = 0.0

	if transitionValue != averageExpected {
		t.Errorf("The expected value is %f, but the transition function return %f\n", averageExpected, transitionValue)
	}
}

func TestRule30D1R1(t *testing.T) {
	d := 1
	r := 1
	m := d*r*(r+1) + 1
	state := []float32{0, 1}
	ruleNumber := 30

	ruleStruct := MakeStateTransitionFunction(len(state), m, ruleNumber)

	transitionTableExpected := [][]float32{
		{0, 0, 0, 0},
		{0, 0, 1, 1},
		{0, 1, 0, 1},
		{0, 1, 1, 1},
		{1, 0, 0, 1},
		{1, 0, 1, 0},
		{1, 1, 0, 0},
		{1, 1, 1, 0},
	}

	height := len(transitionTableExpected)
	width := len(transitionTableExpected[0])

	if len(ruleStruct.transitionTable) != height {
		t.Errorf("The expected number of transitions was %d, but %d was obtained.", height, len(ruleStruct.transitionTable))
	}

	for i := 0; i < height; i += 1 {
		for j := 0; j < width; j += 1 {
			value := ruleStruct.transitionTable[i][j]

			if value != state[0] && value != state[1] {
				t.Errorf("The state %f does not belong to {%f, %f} state set.", value, state[0], state[1])
			}

			if value != transitionTableExpected[i][j] {
				t.Errorf("The expected %dº state in transition %d was %.0f, but %.0f was obtained.", j, i, transitionTableExpected[i][j], value)
			}
		}

		transitionTo := ruleStruct.Transition(transitionTableExpected[i][:m])

		if transitionTableExpected[i][m] != transitionTo {
			t.Errorf("The expected transition is %f, but %f was obtained.", transitionTableExpected[i][m], transitionTo)
		}
	}
}

func TestRule42001D1R1_5(t *testing.T) {
	d := 1
	r := float32(1.5)
	m := int(float32(d)*r*(r+1) + 1)
	state := []float32{0, 1}
	ruleNumber := 42001

	ruleStruct := MakeStateTransitionFunction(len(state), m, ruleNumber)

	transitionTableExpected := [][]float32{
		{0, 0, 0, 0, 1},
		{0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 1, 0},
		{0, 1, 0, 0, 1},
		{0, 1, 0, 1, 0},
		{0, 1, 1, 0, 0},
		{0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 1, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 0, 1, 1},
		{1, 1, 1, 0, 0},
		{1, 1, 1, 1, 1},
	}

	height := len(transitionTableExpected)
	width := len(transitionTableExpected[0])

	if len(ruleStruct.transitionTable) != height {
		t.Errorf("The expected number of transitions was %d, but %d was obtained.", height, len(ruleStruct.transitionTable))
	}

	for i := 0; i < height; i += 1 {
		for j := 0; j < width; j += 1 {
			value := ruleStruct.transitionTable[i][j]

			if value != state[0] && value != state[1] {
				t.Errorf("The state %f does not belong to {%.0f, %.0f} state set.", value, state[0], state[1])
			}

			if value != transitionTableExpected[i][j] {
				t.Errorf("The expected %dº state in transition %d was %.0f, but %.0f was obtained.", j, i, transitionTableExpected[i][j], value)
			}
		}

		transitionTo := ruleStruct.Transition(transitionTableExpected[i][:m])

		if transitionTableExpected[i][m] != transitionTo {
			t.Errorf("The expected transition is %f, but %f was obtained.", transitionTableExpected[i][m], transitionTo)
		}
	}
}
