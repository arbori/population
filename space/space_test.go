package space

import (
	"fmt"
	"testing"
)

var size int
var dimention int
var neighborhoodMotion NeighborhoodMotion

var XEnvironment int
var YEnvironment int
var environment Environment

func TestSetup(t *testing.T) {
	size = 5
	dimention = 2

	neighborhoodMotion = MakeNeighborhoodMotion(size, dimention)
	neighborhoodMotion.Directions[0] = []int{1, 0}
	neighborhoodMotion.Directions[1] = []int{0, -1}
	neighborhoodMotion.Directions[2] = []int{-1, 0}
	neighborhoodMotion.Directions[3] = []int{0, +1}
	neighborhoodMotion.Directions[4] = []int{0, 0}

	XEnvironment = 15
	YEnvironment = 17
	environment = MakeEnvironment(XEnvironment, YEnvironment, &neighborhoodMotion, .1)
}

func TestPointAssignDiferentDimentions(t *testing.T) {
	p1 := Point([]int{0, 1})
	p2 := Point([]int{0, 1, 2})

	err := p1.Assign(&p2)

	if err == nil {
		t.Error("Two points with different dimentions can not be assign.")
	}
}

func TestPointExAddDiferentDimentions(t *testing.T) {
	p1 := Point([]int{0, 1})
	p2 := Point([]int{0, 1, 2})

	err := p1.Add(&p2)

	if err == nil {
		t.Error("Two points with different dimentions can not be add.")
	}
}

func TestPointExAdd(t *testing.T) {
	p1 := Point([]int{1, 2})
	p2 := Point([]int{2, 1})
	expected := Point([]int{3, 3})

	err := p1.Add(&p2)

	if err != nil {
		t.Error(err)
	}

	if p1[0] != expected[0] || p1[1] != expected[1] {
		t.Error("Error add two points")
	}
}

func TestNeighborhoodMotionCreation(t *testing.T) {
	fmt.Printf("neighborhoodMotion.Size supose to be %d\n", size)
	if neighborhoodMotion.Size != size {
		t.Errorf("neighborhoodMotion.Size actual value: %d\n", neighborhoodMotion.Size)
	}

	fmt.Printf("neighborhoodMotion.Dimention supose to be %d\n", size)
	if len(neighborhoodMotion.Directions[0]) != dimention {
		t.Errorf("neighborhoodMotion.Dimention actual value: %d\n", len(neighborhoodMotion.Directions[0]))
	}

	fmt.Printf("neighborhoodMotion.Motion must folow length Size and Dimention\n")
	if len(neighborhoodMotion.Directions) != size || len(neighborhoodMotion.Directions[0]) != dimention {
		t.Errorf("neighborhoodMotion.Motion have size and dimention, respectively %d and %d\n", len(neighborhoodMotion.Directions), len(neighborhoodMotion.Directions[0]))
	}
}

func TestMakeEnvironment(t *testing.T) {
	fmt.Printf("environment.X supose to be %d\n", XEnvironment)
	if environment.X != XEnvironment {
		t.Errorf("environment.X actual value: %d\n", environment.X)
	}

	fmt.Printf("environment.Y supose to be %d\n", YEnvironment)
	if environment.Y != YEnvironment {
		t.Errorf("environment.Y actual value: %d\n", environment.Y)
	}

	fmt.Printf("Size of environment.Cells supose to be %d\n", XEnvironment)
	if len(environment.Cells) != XEnvironment {
		t.Errorf("environment.Cells actual size: %d\n", len(environment.Cells))
	}

	fmt.Printf("Size of environment.Cells[0] supose to be %d\n", YEnvironment)
	if len(environment.Cells[0]) != YEnvironment {
		t.Errorf("environment.Cells[0] actual size: %d\n", len(environment.Cells[0]))
	}
}

func TestNeighborhood(t *testing.T) {
	fmt.Println("TestNeighborhood...")

	environment.Cells[1][1].Value = 0
	environment.Cells[2][1].Value = 1
	environment.Cells[1][0].Value = 2
	environment.Cells[0][1].Value = 3
	environment.Cells[1][2].Value = 4

	neighborhood := environment.NeighborhoodValues(1, 1)

	result := neighborhood[0] != environment.Cells[1][1].Value ||
		neighborhood[1] != environment.Cells[2][1].Value ||
		neighborhood[2] != environment.Cells[1][0].Value ||
		neighborhood[3] != environment.Cells[0][1].Value ||
		neighborhood[4] != environment.Cells[1][2].Value

	if !result {
		t.Errorf("Expected neighborhood is (%d, %d, %d, %d, %d), bus was (%f, %f, %f, %f, %f)\n", 0, 1, 2, 3, 4, neighborhood[0], neighborhood[1], neighborhood[2], neighborhood[3], neighborhood[4])
	}
}

func TestEnvironmentNewPosition(t *testing.T) {
	point := Point([]int{1, 1})
	expectedPoint := Point([]int{0, 1})

	newPosition := environment.GetNewPosition(&point, 2)

	if expectedPoint[0] != newPosition[0] || expectedPoint[1] != newPosition[1] {
		t.Fatalf("The movimento to the wrong point (%d, %d). The expected position is (%d, %d).\n", newPosition[0], newPosition[1], expectedPoint[0], expectedPoint[1])
	}
}
