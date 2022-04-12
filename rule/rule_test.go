package rule

import (
	"fmt"
	"testing"
)

var spreadRuleVonNeumann AverageRuleVonNeumann = AverageRuleVonNeumann{}

func TestSpreadRuleVonNeumannWithouNeighborhood(t *testing.T) {
	transitionValue := spreadRuleVonNeumann.Transition(nil)

	fmt.Println("The expected value for none neighborhood is 0.0")
	if transitionValue != 0 {
		t.Fatalf("The expected value is 0.0, but the transition function return %f\n", transitionValue)
	}
}
