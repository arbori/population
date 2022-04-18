package rule

import "math"

type Rule interface {
	Transition(neighborhood []float32) float32
}

type StateTransitionFunction struct {
	numberOfStates            int
	numberOfNeighborhoodCells int
	ruleNumber                int
	transitionTable           [][]float32
}

func (st StateTransitionFunction) Transition(neighborhood []float32) float32 {
	if len(neighborhood) == 0 {
		return 0
	}

	for i := 0; i < len(st.transitionTable); i += 1 {
		match := true

		for j := 0; match && j < len(st.transitionTable[i])-1; j += 1 {
			match = match && neighborhood[j] == st.transitionTable[i][j]
		}

		if match {
			return st.transitionTable[i][len(st.transitionTable[i])-1]
		}
	}

	return 0
}

func MakeStateTransitionFunction(statesSize int, m int, ruleNumber int) StateTransitionFunction {
	base := float64(statesSize)
	cells := float64(m)

	rule := StateTransitionFunction{
		numberOfStates:            statesSize,
		numberOfNeighborhoodCells: m,
		ruleNumber:                ruleNumber,
	}

	var transitionsSize int = int(math.Pow(base, cells))

	rule.transitionTable = make([][]float32, transitionsSize)

	ruleconverted := numBaseConv(ruleNumber, statesSize)
	convPos := 0
	convSize := len(ruleconverted)

	value := float32(ruleconverted[convPos])
	convPos += 1

	rule.transitionTable[0] = nextTransition(nil, cells, value, float32(statesSize))

	for i := 1; i < transitionsSize; i += 1 {
		if convPos < convSize {
			value = float32(ruleconverted[convPos])
			convPos += 1
		} else {
			value = 0
		}

		rule.transitionTable[i] = nextTransition(rule.transitionTable[i-1], cells, value, float32(statesSize))
	}

	return rule
}

func nextTransition(current []float32, cells float64, value float32, max float32) []float32 {
	m := int(cells)

	next := make([]float32, m+1)

	for j := 0; j < m; j += 1 {
		if current == nil {
			next[j] = 0
		} else {
			next[j] = current[j]
		}
	}

	for j := m - 1; current != nil && j >= 0; j -= 1 {
		next[j] = current[j] + 1

		if next[j] < max {
			break
		}

		next[j] = 0
	}

	next[m] = value

	return next
}

func numBaseConv(n int, base int) []int {
	result := make([]int, 0)

	for n >= base {
		result = append(result, n%base)
		n = n / base
	}

	if n > 0 {
		result = append(result, n)
	}

	return result
}

type AverageRuleVonNeumann struct {
}

func (rule AverageRuleVonNeumann) Transition(neighborhood []float32) float32 {
	return average(neighborhood, 5)
}

type AverageRuleMoore struct {
}

func (rule AverageRuleMoore) Transition(neighborhood []float32) float32 {
	return average(neighborhood, 9)
}

func average(neighborhood []float32, expectedSize int) float32 {
	var result float32 = 0.0

	size := len(neighborhood)

	if size != expectedSize {
		return result
	}

	for i := 0; i < size; i += 1 {
		result += neighborhood[i]
	}

	return (result / float32(size))
}
