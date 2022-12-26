package domain

import "math/rand"

type Range struct {
	Minimum interface{}
	Maximum interface{}
	Delta   interface{}
}

type SimulationData struct {
	IndividualConsumption int
	IndividualExchange    int
	SocietySize           int
	ViabilityAmount       int
	SolidarityProbability float32
	Ephocs                int
}

func ExchangeResourceRule(first *Individual, second *Individual, individualExchange int, solidarityProbability float32) {
	// Make exchange between ra and rb.
	solidarity := rand.Float32()

	if (solidarityProbability < solidarity && first.Resources > second.Resources) ||
		(solidarityProbability > solidarity && first.Resources < second.Resources) {
		first.Resources += individualExchange
		second.Resources -= individualExchange
	} else if (solidarityProbability > solidarity && first.Resources > second.Resources) ||
		(solidarityProbability < solidarity && first.Resources < second.Resources) {
		first.Resources -= individualExchange
		second.Resources += individualExchange
	}
}
