package agent

import (
	"math/rand"

	"github.com/arbori/population.git/population/cellularautomata"
	"github.com/arbori/population.git/population/lattice"
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

// Retrieve the set of agents in the neighborhood of a cell.
func retriveAgentsInTheNeighborhood(e *space.Environment, x int, y int) []MobileAgent {
	result := make([]MobileAgent, 0, e.Motion.Size)

	if len(e.Motion.Directions) == 0 || len(e.Motion.Directions) != e.Motion.Size {
		return result
	}

	var i int
	var j int

	for index := 0; index < e.Motion.Size; index += 1 {
		j = e.Motion.Directions[index][0] + x
		i = e.Motion.Directions[index][1] + y

		if j < 0 {
			j = e.X + j
		} else if j >= e.X {
			j = j - e.X
		}

		if i < 0 {
			i = e.Y + i
		} else if i >= e.Y {
			i = i - e.Y
		}

		if e.Cells[j][i].Content != nil && e.Cells[j][i].Content.(MobileAgent).IsAvailable {
			result = append(result, e.Cells[j][i].Content.(MobileAgent))
		}
	}

	return result
}

// Identify gents's position in the environment to do exchange between them.
// Each 2 agents in the same neighborhood, can be choose to exchange resources.
func AgentsExchangeResources(e *space.Environment, exc *Exchange, interationRule InterationRuleType) {
	var firsti int
	var secoundi int

	for y := 0; y < e.Y; y += 1 {
		for x := 0; x < e.X; x += 1 {
			agentsInTheNeighborhood := retriveAgentsInTheNeighborhood(e, x, y)

			for len(agentsInTheNeighborhood) > 1 {
				firsti = rand.Intn(len(agentsInTheNeighborhood))

				for secoundi = firsti; secoundi == firsti; secoundi = rand.Intn(len(agentsInTheNeighborhood)) {
				}

				interationRule(&agentsInTheNeighborhood[firsti], &agentsInTheNeighborhood[secoundi], exc.ContribuitionProbability, exc.ExchangeRate)

				agentsInTheNeighborhood[firsti].IsAvailable = false
				agentsInTheNeighborhood[secoundi].IsAvailable = false

				agentsInTheNeighborhood = append(agentsInTheNeighborhood[:firsti], agentsInTheNeighborhood[firsti+1:]...)
				agentsInTheNeighborhood = append(agentsInTheNeighborhood[:secoundi], agentsInTheNeighborhood[secoundi+1:]...)
			}
		}
	}
}

type MotionRuleType func(env *lattice.Lattice, ca *cellularautomata.Cellularautomata, posiotion space.Point) space.Point

type MobileAgent struct {
	Position    space.Point
	Resources   float32
	MotionRule  MotionRuleType
	IsAvailable bool
}

func (a *MobileAgent) Walk(env *lattice.Lattice, ca *cellularautomata.Cellularautomata) {
	position := a.MotionRule(env, ca, a.Position)

	if env.At(position...) == nil {
		env.Set(nil, a.Position...)
		env.Set(a, position...)

		for i := 0; i < len(position); i += 1 {
			a.Position[i] = position[i]
		}
	}

	ca.Set(a.Resources, a.Position...)
}

// func GetNewPosition(position *space.Point, X *space.Point, motion *[][]int, directionChoosed int) space.Point {
// 	return lattice.Enclose(space.Point{
// 		(*motion)[directionChoosed][0] + (*position)[0],
// 		(*motion)[directionChoosed][1] + (*position)[1]}, *X)
// }
