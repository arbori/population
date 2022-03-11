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
	neighborhoodMotion.Motion[0] = NewPoint(1, 0)
	neighborhoodMotion.Motion[1] = NewPoint(0, -1)
	neighborhoodMotion.Motion[2] = NewPoint(-1, 0)
	neighborhoodMotion.Motion[3] = NewPoint(0, +1)
	neighborhoodMotion.Motion[4] = NewPoint(0, 0)

	XEnvironment = 15
	YEnvironment = 17
	environment = MakeEnvironment(XEnvironment, YEnvironment, &neighborhoodMotion, .1)
}

func TestPointAssignDiferentDimentions(t *testing.T) {
	p1 := NewPoint(0, 1)
	p2 := NewPoint(0, 1, 2)

	err := p1.Assign(&p2)

	if err == nil {
		t.Error("Two points with different dimentions can not be assign.")
	}
}

func TestPointExAddDiferentDimentions(t *testing.T) {
	p1 := NewPoint(0, 1)
	p2 := NewPoint(0, 1, 2)

	err := p1.Add(&p2)

	if err == nil {
		t.Error("Two points with different dimentions can not be add.")
	}
}

func TestPointExAdd(t *testing.T) {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(2, 1)
	expected := NewPoint(3, 3)

	err := p1.Add(&p2)

	if err != nil {
		t.Error(err)
	}

	if p1.X[0] != expected.X[0] || p1.X[1] != expected.X[1] {
		t.Error("Error add two points")
	}
}

func TestNeighborhoodMotionCreation(t *testing.T) {
	fmt.Printf("neighborhoodMotion.Size supose to be %d\n", size)
	if neighborhoodMotion.Size != size {
		t.Errorf("neighborhoodMotion.Size actual value: %d\n", neighborhoodMotion.Size)
	}

	fmt.Printf("neighborhoodMotion.Dimention supose to be %d\n", size)
	if neighborhoodMotion.Motion[0].Dim != dimention {
		t.Errorf("neighborhoodMotion.Dimention actual value: %d\n", neighborhoodMotion.Motion[0].Dim)
	}

	fmt.Printf("neighborhoodMotion.Motion must folow length Size and Dimention\n")
	if len(neighborhoodMotion.Motion) != size || neighborhoodMotion.Motion[0].Dim != dimention {
		t.Errorf("neighborhoodMotion.Motion have size and dimention, respectively %d and %d\n", len(neighborhoodMotion.Motion), neighborhoodMotion.Motion[0].Dim)
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

	environment.Cells[1][1] = 0
	environment.Cells[2][1] = 1
	environment.Cells[1][0] = 2
	environment.Cells[0][1] = 3
	environment.Cells[1][2] = 4

	neighborhood := environment.Neighborhood(1, 1)

	result := neighborhood[0] != environment.Cells[1][1] ||
		neighborhood[1] != environment.Cells[2][1] ||
		neighborhood[2] != environment.Cells[1][0] ||
		neighborhood[3] != environment.Cells[0][1] ||
		neighborhood[4] != environment.Cells[1][2]

	if !result {
		t.Errorf("Expected neighborhood is (%d, %d, %d, %d, %d), bus was (%f, %f, %f, %f, %f)\n", 0, 1, 2, 3, 4, neighborhood[0], neighborhood[1], neighborhood[2], neighborhood[3], neighborhood[4])
	}
}

func TestEnvironmentNewPosition(t *testing.T) {
	point := NewPoint(1, 1)
	expectedPoint := NewPoint(0, 1)

	newPosition := environment.GetNewPosition(&point, 2)

	if expectedPoint.X[0] != newPosition.X[0] || expectedPoint.X[1] != newPosition.X[1] {
		t.Fatalf("The movimento to the wrong point (%d, %d). The expected position is (%d, %d).\n", newPosition.X[0], newPosition.X[1], expectedPoint.X[0], expectedPoint.X[1])
	}
}
