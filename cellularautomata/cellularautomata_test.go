package cellularautomata

import (
	"fmt"
	"math"
	"testing"

	"github.com/arbori/population.git/population/lattice"
	"github.com/arbori/population.git/population/rule"
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

func TestCANeighborhood(t *testing.T) {
	d := 2
	r := 1
	m := d*r*(r+1) + 1
	ruleNumber := 30

	states := []float32{0, 1}
	motion := NeighborhoodMotionVonNeumman(d, r)
	ruleStruct := rule.MakeStateTransitionFunction(len(states), m, ruleNumber)

	ca, err := New(states, motion, ruleStruct, 7, 5)

	if err != nil {
		t.Error(err)
	}

	ca.Set(0, 1, 1)
	ca.Set(1, 2, 1)
	ca.Set(1, 3, 1)

	ca.Set(1, 1, 2)
	ca.Set(0, 2, 2)
	ca.Set(1, 3, 2)

	ca.Set(1, 1, 3)
	ca.Set(1, 2, 3)
	ca.Set(0, 3, 3)

	neighborhood := ca.NeighborhoodValues(2, 2)

	if neighborhood == nil {
		t.Error("Neighborhood vector supose to be not nil")
	}

	for i := 0; i < len(neighborhood); i += 1 {
		cellValue := ca.Get(2+motion[i][0], 2+motion[i][1])

		if neighborhood[i] != cellValue {
			t.Errorf("neighborhood[%d] == %f, but the correct value is %f.", i, neighborhood[i], cellValue)
		}
	}
}

func TestTemporalEvolution(t *testing.T) {
	d := 1
	r := 1
	m := d*r*(r+1) + 1
	states := []float32{0, 1}
	ruleNumber := 30
	length := 41
	time := 3

	motion := NeighborhoodMotionVonNeumman(d, r)
	rule := rule.MakeStateTransitionFunction(len(states), m, ruleNumber)

	expectedDimention := 1
	expectedLimit := []int{length, time}
	expectedTemporalEvolution, latticeErr := lattice.NewWithValue(float32(0), expectedLimit...)

	if latticeErr != nil {
		t.Error(latticeErr)
	}

	expectedTemporalEvolution.Set(float32(1), 21, 0)
	expectedTemporalEvolution.Set(float32(1), 20, 1)
	expectedTemporalEvolution.Set(float32(1), 21, 1)
	expectedTemporalEvolution.Set(float32(1), 22, 1)
	expectedTemporalEvolution.Set(float32(1), 19, 2)
	expectedTemporalEvolution.Set(float32(1), 20, 2)
	expectedTemporalEvolution.Set(float32(1), 23, 2)

	ca, err := New(states, motion, rule, expectedLimit[0])

	if err != nil {
		t.Error(err)
	}

	limit := ca.Limits()

	if len(limit) != expectedDimention || limit[0] != expectedLimit[0] {
		t.Error("The dimention and limits suppose to be 1 and 41, respectively.")
	}

	ca.Set(1, 21)

	for y := 0; y < 3; y += 1 {
		for x := 0; x < limit[0]; x += 1 {
			if ca.Get(x) != expectedTemporalEvolution.At(x, y).(float32) {
				te := ""
				teExpec := ""

				for j := 0; j < limit[0]; j += 1 {
					te += fmt.Sprintf("%.0f", ca.Get(j))
					teExpec += fmt.Sprintf("%.0f", expectedTemporalEvolution.At(j, y).(float32))
				}

				t.Errorf("Temporal evolution is wrong.\nExpected: %s\nCurrent:  %s", teExpec, te)

				break
			}
		}

		ca.Evolve()
	}
}
