package agent

import (
	"github.com/arbori/population.git/population/space"
)

type InterationRuleType func(a1 *MobileAgent, a2 *MobileAgent)

type Exchange struct {
	ContribuitionRate float32
	ExchangeRate      float32
	InterationRule    InterationRuleType
}

type MotionRuleType func(environment *space.Environment, position *space.Point) space.Point

type MobileAgent struct {
	Position   space.Point
	Foodstuffs float32
	MotionRule MotionRuleType
}

func (a *MobileAgent) Walk(env *space.Environment) {
	velocity := a.MotionRule(env, &a.Position)

	a.Position.Add(&velocity)
}
