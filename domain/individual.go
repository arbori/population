package domain

import "math/rand"

// Individou of population.
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

func MakePopulation(societySize int, individualConsumption int) []Individual {
	population := make([]Individual, societySize)

	// Sortear a quantidade de recursos para os indivíduos de I.
	for i := range population {
		population[i] = MakeIndividual(individualConsumption * rand.Intn(100))
	}

	return population
}
