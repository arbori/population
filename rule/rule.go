package rule

type Rule interface {
	Transition(neighborhood []float32) float32
}

type SpreadRuleVonNeumann struct {
	Decay float32
}

func (rule SpreadRuleVonNeumann) Transition(neighborhood []float32) float32 {
	var result float32 = 0.0

	size := len(neighborhood)

	if size != 5 {
		return result
	}

	result = (1-rule.Decay)*neighborhood[0] + (rule.Decay/float32(size))*(neighborhood[1]+neighborhood[2]+neighborhood[3]+neighborhood[4])

	return result
}

type SpreadRuleMoore struct {
	Decay float32
}

func (rule SpreadRuleMoore) Transition(neighborhood []float32) float32 {
	return TransitionMax(neighborhood, rule.Decay)
}

func TransitionMax(neighborhood []float32, decay float32) float32 {
	var result float32 = 0.0

	size := len(neighborhood)

	if size != 9 {
		return result
	}

	var maxIdx int
	var maxValue float32 = 0.0

	for i := 0; i < size; i += 1 {
		if maxValue < neighborhood[i] {
			maxIdx = i
			maxValue = neighborhood[i]
		}
	}

	result = (1 - decay) * neighborhood[maxIdx]

	return result
}

func TransitionAverageNeighborhood(neighborhood []float32, decay float32) float32 {
	var result float32 = 0.0

	size := len(neighborhood)

	if size != 9 {
		return result
	}

	for i := 1; i < size; i += 1 {
		result += neighborhood[i]
	}
	result = (neighborhood[0] + (result / float32(size))) / 2.0

	return result
}

func TransitionSpreadDecay(neighborhood []float32, decay float32) float32 {
	var result float32 = 0.0

	size := len(neighborhood)

	if size != 9 {
		return result
	}

	for i := 1; i < size; i += 1 {
		result += neighborhood[i]
	}
	result = (1-decay)*neighborhood[0] + (decay/float32(size))*(result)

	return result / float32(size)
}
