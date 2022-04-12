package cellularautomata

import (
	"math"
	"testing"
)

func TestNeighborhoodVonNeummanCreate(t *testing.T) {
	d := 1
	r := 1

	neighborhood := NeighborhoodMotionVonNeumman(d, r)

	expectedSize := d*r*(r+1) + 1
	actualSize := len(neighborhood)

	if actualSize != expectedSize {
		t.Errorf("The expected size for neighborhood is %d, but %d was obtained.", expectedSize, actualSize)
		return
	}

	var expectedNeighborhood [][]int = [][]int{
		{-1},
		{0},
		{1},
	}

	for i := 0; i < expectedSize; i += 1 {
		if len(neighborhood[i]) != d {
			t.Errorf("The expected size of motion is %d, but %d was obtained.", d, len(neighborhood[i]))
			return
		}

		if expectedNeighborhood[i][0] != neighborhood[i][0] {
			t.Errorf("The expected value for motion %d is %d, but %d was obtained.", i, expectedNeighborhood[i][0], neighborhood[i][0])
			return
		}
	}

	///////////////////////////////
	d = 3
	r = 1

	neighborhood = NeighborhoodMotionVonNeumman(d, r)

	expectedSize = d*r*(r+1) + 1
	actualSize = len(neighborhood)

	if actualSize != expectedSize {
		t.Errorf("The expected size for neighborhood is %d, but %d was obtained.", expectedSize, actualSize)
		return
	}

	expectedNeighborhood = [][]int{
		{-1, 0, 0},
		{0, -1, 0},
		{0, 0, -1},
		{0, 0, 0},
		{0, 0, 1},
		{0, 1, 0},
		{1, 0, 0},
	}

	for i := 0; i < expectedSize; i += 1 {
		if len(neighborhood[i]) != d {
			t.Errorf("The expected size of motion is %d, but %d was obtained.", d, len(neighborhood[i]))
			return
		}

		if expectedNeighborhood[i][0] != neighborhood[i][0] {
			t.Errorf("The expected value for motion %d is %d, but %d was obtained.", i, expectedNeighborhood[i][0], neighborhood[i][0])
			return
		}
	}

}

func TestNeighborhoodMooreCreate(t *testing.T) {
	d := 1
	r := 1

	neighborhood := NeighborhoodMotionMoore(d, r)

	expectedSize := int(math.Pow(float64(2*r+1), float64(d)))
	actualSize := len(neighborhood)

	if actualSize != expectedSize {
		t.Errorf("The expected size for neighborhood is %d, but %d was obtained.", expectedSize, actualSize)
		return
	}

	var expectedNeighborhood [][]int = [][]int{
		{-1},
		{0},
		{1},
	}

	for i := 0; i < expectedSize; i += 1 {
		if len(neighborhood[i]) != d {
			t.Errorf("The expected size of motion is %d, but %d was obtained.", d, len(neighborhood[i]))
			return
		}

		if expectedNeighborhood[i][0] != neighborhood[i][0] {
			t.Errorf("The expected value for motion %d is %d, but %d was obtained.", i, expectedNeighborhood[i][0], neighborhood[i][0])
			return
		}
	}

	////////////////
	d = 3
	r = 1

	neighborhood = NeighborhoodMotionMoore(d, r)

	expectedSize = int(math.Pow(float64(2*r+1), float64(d)))
	actualSize = len(neighborhood)

	if actualSize != expectedSize {
		t.Errorf("The expected size for neighborhood is %d, but %d was obtained.", expectedSize, actualSize)
		return
	}

	expectedNeighborhood = [][]int{
		{-1, -1, -1},
		{-1, -1, 0},
		{-1, -1, 1},
		{-1, 0, -1},
		{-1, 0, 0},
		{-1, 0, 1},
		{-1, 1, -1},
		{-1, 1, 0},
		{-1, 1, 1},
		{0, -1, -1},
		{0, -1, 0},
		{0, -1, 1},
		{0, 0, -1},
		{0, 0, 0},
		{0, 0, 1},
		{0, 1, -1},
		{0, 1, 0},
		{0, 1, 1},
		{1, -1, -1},
		{1, -1, 0},
		{1, -1, 1},
		{1, 0, -1},
		{1, 0, 0},
		{1, 0, 1},
		{1, 1, -1},
		{1, 1, 0},
		{1, 1, 1},
	}

	for i := 0; i < expectedSize; i += 1 {
		if len(neighborhood[i]) != d {
			t.Errorf("The expected size of motion is %d, but %d was obtained.", d, len(neighborhood[i]))
			return
		}

		if expectedNeighborhood[i][0] != neighborhood[i][0] {
			t.Errorf("The expected value for motion %d is %d, but %d was obtained.", i, expectedNeighborhood[i][0], neighborhood[i][0])
			return
		}
	}
}

func TestRuleCreation(t *testing.T) {
	d := 1
	r := 1
	m := d*r*(r+1) + 1
	state := []float32{0, 1}
	ruleNumber := 30

	rule := MakeStateTransitionFunction(len(state), m, ruleNumber)

	if rule == nil {
		t.Errorf("The rule expected to be not nil.")
	}

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

	if len(rule.(StateTransitionFunction).transitionTable) != height {
		t.Errorf("The expected number of transitions was %d, but %d was obtained.", height, len(rule.(StateTransitionFunction).transitionTable))
	}

	for i := 0; i < height; i += 1 {
		for j := 0; j < width; j += 1 {
			value := rule.(StateTransitionFunction).transitionTable[i][j]

			if value != state[0] && value != state[1] {
				t.Errorf("The state %f does not belong to {%f, %f} state set.", value, state[0], state[1])
			}

			if value != transitionTableExpected[i][j] {
				t.Errorf("The expected state was %f, but %f was obtained.", transitionTableExpected[i][j], value)
			}
		}

		transitionTo := rule.Transition(transitionTableExpected[i][:m])

		if transitionTableExpected[i][m] != transitionTo {
			t.Errorf("The expected transition is %f, but %f was obtained.", transitionTableExpected[i][m], transitionTo)
		}
	}
}
