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
