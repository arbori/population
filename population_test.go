package main

import (
	"log"
	"testing"

	"github.com/arbori/population.git/population/rule"
)

func TestAplyRuleSimulation(t *testing.T) {
	log.Println("TestAplyRuleSimulation - Start test")
	motion := constructVonNeumannNeighborhoodMotion()
	environment := constructEnvironment(&motion)

	rule := rule.SpreadRuleVonNeumann{
		Decay: .15,
	}

	x := 2
	y := 2
	environment.Cells[x][y].Value = 7

	center := (1-rule.Decay)*environment.Cells[x][y].Value + (rule.Decay/5.0)*(environment.Cells[x-1][y].Value+environment.Cells[x][y+1].Value+environment.Cells[x+1][y].Value+environment.Cells[x][y-1].Value)
	x = 1
	y = 2
	left := (1-rule.Decay)*environment.Cells[x][y].Value + (rule.Decay/5.0)*(environment.Cells[x-1][y].Value+environment.Cells[x][y+1].Value+environment.Cells[x+1][y].Value+environment.Cells[x][y-1].Value)
	x = 2
	y = 3
	button := (1-rule.Decay)*environment.Cells[x][y].Value + (rule.Decay/5.0)*(environment.Cells[x-1][y].Value+environment.Cells[x][y+1].Value+environment.Cells[x+1][y].Value+environment.Cells[x][y-1].Value)
	x = 3
	y = 2
	right := (1-rule.Decay)*environment.Cells[x][y].Value + (rule.Decay/5.0)*(environment.Cells[x-1][y].Value+environment.Cells[x][y+1].Value+environment.Cells[x+1][y].Value+environment.Cells[x][y-1].Value)
	x = 2
	y = 1
	top := (1-rule.Decay)*environment.Cells[x][y].Value + (rule.Decay/5.0)*(environment.Cells[x-1][y].Value+environment.Cells[x][y+1].Value+environment.Cells[x+1][y].Value+environment.Cells[x][y-1].Value)

	environment.ApplyRule(rule)

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
