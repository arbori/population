package rule

type Rule interface {
	Transition(neighborhood []float32) float32
}

type AverageRuleVonNeumann struct {
}

func (rule AverageRuleVonNeumann) Transition(neighborhood []float32) float32 {
	return average(neighborhood, 9)
}

type AverageRuleMoore struct {
}

func (rule AverageRuleMoore) Transition(neighborhood []float32) float32 {
	return average(neighborhood, 5)
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
