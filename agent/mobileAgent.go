package agent

import (
	"github.com/arbori/population.git/population/space"
)

type InterationRuleType func(a1 *MobileAgent, a2 *MobileAgent, contribuitionProbability float32, exchangeRate float32)

type Exchange struct {
	ContribuitionProbability float32
	ExchangeRate             float32
	interationRule           InterationRuleType
}

func MakeExchange(contribuitionProbability float32, exchangeRate float32, interationRule InterationRuleType) Exchange {
	return Exchange{
		ContribuitionProbability: contribuitionProbability,
		ExchangeRate:             exchangeRate,
		interationRule:           interationRule,
	}
}

func (e *Exchange) Interation(a1 *MobileAgent, a2 *MobileAgent) {
	e.interationRule(a1, a2, e.ContribuitionProbability, e.ExchangeRate)
}

type MotionRuleType func(environment *space.Environment, position *space.Point) space.Point

type MobileAgent struct {
	Position   space.Point
	Foodstuffs float32
	MotionRule MotionRuleType
}

func (a *MobileAgent) Walk(env *space.Environment) {
	velocity := a.MotionRule(env, &a.Position)

	a.Position.Assign(&velocity)
}
