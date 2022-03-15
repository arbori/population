package rule

type Rule interface {
	Transition(neighborhood []float32) float32
}

type SpreadRuleVonNeumann struct {
	Decay float32
}

func (rule SpreadRuleVonNeumann) Transition(neighborhood []float32) float32 {
	return TransitionAverageNeighborhoodVonNeumann(neighborhood, rule.Decay)
}

func TransitionSpreadDecayVonNeumann(neighborhood []float32, decay float32) float32 {
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

func TransitionSpreadIntegerVonNeumann(neighborhood []float32, decay float32) float32 {
	var result float32 = 0.0
	var greater float32 = 0.0
	var less float32 = 0.0

	size := len(neighborhood)

	if size != 5 {
		return neighborhood[0]
	}

	for i := 1; i < size; i += 1 {
		if neighborhood[i] > neighborhood[0] {
			greater += 1.0
		} else if neighborhood[i] < neighborhood[0] {
			less += 1.0
		}
	}

	if neighborhood[0]-decay >= 0.0 {
		result = (neighborhood[0] - decay) + greater - less
	} else {
		result = neighborhood[0] + greater - less
	}

	if result <= 0.0 {
		result = 0.0
	}

	return result
}

func TransitionAverageNeighborhoodVonNeumann(neighborhood []float32, decay float32) float32 {
	var result float32 = 0.0

	size := len(neighborhood)

	if size != 5 {
		return result
	}

	for i := 0; i < size; i += 1 {
		result += neighborhood[i]
	}

	return (result / float32(size))
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
