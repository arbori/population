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
	neighborhoodMotion.Motion[0][0] = +1
	neighborhoodMotion.Motion[0][1] = 0
	neighborhoodMotion.Motion[1][0] = 0
	neighborhoodMotion.Motion[1][1] = -1
	neighborhoodMotion.Motion[2][0] = -1
	neighborhoodMotion.Motion[2][1] = 0
	neighborhoodMotion.Motion[3][0] = 0
	neighborhoodMotion.Motion[3][1] = +1
	neighborhoodMotion.Motion[4][0] = 0
	neighborhoodMotion.Motion[4][0] = 0

	XEnvironment = 15
	YEnvironment = 17
	environment = MakeEnvironment(XEnvironment, YEnvironment, &neighborhoodMotion, .1)

}

func TestPointAdd(t *testing.T) {
	p1 := Point{
		X: 3,
		Y: 2,
	}

	p2 := Point{
		X: 2,
		Y: 3,
	}

	p2.Add(&p1)

	fmt.Println("Expected point cordinate are (5, 5).")

	if p2.X != 5 || p2.Y != 5 {
		t.Fatalf("Point.Add is wrong. Actual coordinate are (%d, %d)\n", p2.X, p2.Y)
	}
}

func TestNeighborhoodMotionCreation(t *testing.T) {
	fmt.Printf("neighborhoodMotion.Size supose to be %d\n", size)
	if neighborhoodMotion.Size != size {
		t.Errorf("neighborhoodMotion.Size actual value: %d\n", neighborhoodMotion.Size)
	}

	fmt.Printf("neighborhoodMotion.Dimention supose to be %d\n", size)
	if neighborhoodMotion.Dimention != dimention {
		t.Errorf("neighborhoodMotion.Dimention actual value: %d\n", neighborhoodMotion.Dimention)
	}

	fmt.Printf("neighborhoodMotion.Motion must folow length Size and Dimention\n")
	if len(neighborhoodMotion.Motion) != size || len(neighborhoodMotion.Motion[0]) != dimention {
		t.Errorf("neighborhoodMotion.Motion have size and dimention, respectively %d and %d\n", len(neighborhoodMotion.Motion), len(neighborhoodMotion.Motion[0]))
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
	point := Point{
		X: 1,
		Y: 1,
	}
	expectedPoint := Point{
		X: 0,
		Y: 1,
	}

	newPosition := environment.GetNewPosition(&point, 2)

	if expectedPoint.X != newPosition.X || expectedPoint.Y != newPosition.Y {
		t.Fatalf("The movimento to the wrong point (%d, %d). The expected position is (%d, %d).\n", newPosition.X, newPosition.Y, expectedPoint.X, expectedPoint.Y)
	}
}
