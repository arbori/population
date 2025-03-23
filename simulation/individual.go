package simulation

import "math/rand"

// Individual simulate a person of population.
// Resources is the amount of resources the individual possess.
// History is the set of resources the individual possess during simulation.
type Individual struct {
	Resources int
	History   []int
}

// Create a individuo with amount of resource and history of transactions.
func MakeIndividual(resources int) Individual {
	return Individual{
		Resources: resources,
		History:   make([]int, 0),
	}
}

// Initializes the society gives the number of individuals and level of consumption.
func MakeSociety(societySize int, individualConsumption int) []Individual {
	population := make([]Individual, societySize)

	// Sortear a quantidade de recursos para os indivíduos de I.
	for i := range population {
		population[i] = MakeIndividual(individualConsumption * rand.Intn(100))
	}

	return population
}

// Apply the rule of exchange between two individuals.
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
