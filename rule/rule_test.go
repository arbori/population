package rule

import (
	"fmt"
	"testing"
)

var spreadRuleVonNeumann SpreadRuleVonNeumann = SpreadRuleVonNeumann{
	Decay: .15,
}

func transitionFunctionSpreadRuleVonNeumann(neighborhood []float32, decay float32) float32 {
	size := len(neighborhood)
	return (1-decay)*neighborhood[0] + (decay/float32(size))*(neighborhood[1]+neighborhood[2]+neighborhood[3]+neighborhood[4])
}

func TestSpreadRuleVonNeumannWithouNeighborhood(t *testing.T) {
	transitionValue := spreadRuleVonNeumann.Transition(nil)

	fmt.Println("The expected value for none neighborhood is 0.0")
	if transitionValue != 0 {
		t.Fatalf("The expected value is 0.0, but the transition function return %f\n", transitionValue)
	}
}

func TestSpreadRuleVonNeumannEmptyNeighborhood(t *testing.T) {
	expectedValue := transitionFunctionSpreadRuleVonNeumann([]float32{0, 0, 0, 0, 0}, spreadRuleVonNeumann.Decay)
	transitionValue := spreadRuleVonNeumann.Transition([]float32{0, 0, 0, 0, 0})

	fmt.Printf("The expected value for neighborhood with zeros is %f\n", expectedValue)
	if transitionValue != expectedValue {
		t.Fatalf("The expected value is %f, but the transition function return %f\n", expectedValue, transitionValue)
	}
}

func TestSpreadRuleVonNeumann(t *testing.T) {
	neighborhood := []float32{1, 2, 3, 4, 5}
	expectedValue := transitionFunctionSpreadRuleVonNeumann(neighborhood, spreadRuleVonNeumann.Decay)
	transitionValue := spreadRuleVonNeumann.Transition(neighborhood)

	fmt.Printf("The expected value for neighborhood {1, 2, 3, 4, 5} is %f\n", expectedValue)
	if transitionValue != expectedValue {
		t.Fatalf("The expected value is %f, but the transition function return %f\n", expectedValue, transitionValue)
	}
}
