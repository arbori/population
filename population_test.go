package main

import (
	"log"
	"testing"

	//"github.com/arbori/population.git/population/rule"
	"github.com/arbori/population.git/population/space"
)

func TestAplyRuleSimulation(t *testing.T) {
	log.Println("TestAplyRuleSimulation - Start test")
	motion := vonNeumannNeighborhoodMotion
	environment := constructEnvironment(&motion)

	// rule := rule.AverageRuleVonNeumann{
	// }

	average := func(environment *space.Environment, x int, y int) float32 {
		return (environment.Cells[x][y].Value + environment.Cells[x-1][y].Value + environment.Cells[x][y+1].Value + environment.Cells[x+1][y].Value + environment.Cells[x][y-1].Value) / 5
	}

	x := 2
	y := 2
	environment.Cells[x][y].Value = 7

	center := average(&environment, x, y)
	x = 1
	y = 2
	left := average(&environment, x, y)
	x = 2
	y = 3
	button := average(&environment, x, y)
	x = 3
	y = 2
	right := average(&environment, x, y)
	x = 2
	y = 1
	top := average(&environment, x, y)

	//environment.ApplyRule(rule)

	x = 2
	y = 2
	if center != environment.Cells[x][y].Value {
		t.Errorf("The expected value for cell (%d, %d) was %f, but the actual value is %f\n", x, y, center, environment.Cells[x][y])
	}
	x = 1
	y = 2
	if left != environment.Cells[x][y].Value {
		t.Errorf("The expected value for cell (%d, %d) was %f, but the actual value is %f\n", x, y, left, environment.Cells[x][y])
	}
	x = 2
	y = 3
	if button != environment.Cells[x][y].Value {
		t.Errorf("The expected value for cell (%d, %d) was %f, but the actual value is %f\n", x, y, button, environment.Cells[x][y])
	}
	x = 3
	y = 2
	if right != environment.Cells[x][y].Value {
		t.Errorf("The expected value for cell (%d, %d) was %f, but the actual value is %f\n", x, y, right, environment.Cells[x][y])
	}
	x = 2
	y = 1
	if top != environment.Cells[x][y].Value {
		t.Errorf("The expected value for cell (%d, %d) was %f, but the actual value is %f\n", x, y, top, environment.Cells[x][y])
	}

	log.Println("TestAplyRuleSimulation - Test has been concluded.")
}
