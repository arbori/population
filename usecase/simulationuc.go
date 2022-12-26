package simulation

import (
	"math/rand"

	domain "github.com/arbori/population.git/population/domain"
)

// Choose two indexes for two individual based in the size population.
func chooseIndividualsIndexes(size int) (int, int) {
	firstIndex := rand.Intn(size)
	secondIndex := rand.Intn(size)

	// The individuos must to be diferent
	for firstIndex == secondIndex {
		secondIndex = rand.Intn(size)
	}

	return firstIndex, secondIndex
}

// Have each individual exchange resources with someone else. The rule is: In the
// exchange between two individuals, those who have the most receive and those who
// have the least give.
func exchangeResource(population []domain.Individual, individualConsumption int, individualExchange int, solidarityProbability float32) ([]domain.Individual, []domain.Individual) {
	newPopulation := make([]domain.Individual, 0, len(population))
	var notViable []domain.Individual

	removed := make([]domain.Individual, 0)

	// Make exchange while more than two individuals did not exchange yet.
	for len(population) >= 2 {
		firstIndex, secondIndex := chooseIndividualsIndexes(len(population))

		domain.ExchangeResourceRule(&population[firstIndex], &population[secondIndex], individualExchange, solidarityProbability)

		// Move to new population the viables individuals and to removed the inviables.
		newPopulation, notViable = moveViablesIndividuals(newPopulation, individualConsumption, &population[firstIndex], &population[secondIndex])
		removed = append(removed, notViable...)

		// Remove from current population which ones that already exchanges resources.
		if firstIndex < secondIndex {
			population = append(append(population[:firstIndex], population[firstIndex+1:secondIndex]...), population[secondIndex+1:]...)
		} else {
			population = append(append(population[:secondIndex], population[secondIndex+1:firstIndex]...), population[firstIndex+1:]...)
		}
	}

	// Set the new population
	if len(population) > 0 {
		newPopulation = append(newPopulation, population...)
	}

	return newPopulation, removed
}

// Moves individual to the new population if it is still viable after resource consumption.
func moveViablesIndividuals(newPopulation []domain.Individual, IndividualConsumption int, individuals ...*domain.Individual) ([]domain.Individual, []domain.Individual) {
	notViable := make([]domain.Individual, 0)

	for _, individual := range individuals {
		individualConsumption(individual, IndividualConsumption)

		if individual.Resources >= IndividualConsumption {
			newPopulation = append(newPopulation, *individual)
		} else {
			notViable = append(notViable, *individual)
		}
	}

	return newPopulation, notViable
}

// Individual consumption of amount
func individualConsumption(individual *domain.Individual, IndividualConsumption int) {
	individual.Resources -= IndividualConsumption
	individual.History = append(individual.History, individual.Resources)
}

/*
// Remove from population the inviable individuos, returning the new population set
// and the individuos removed.
func removeInviableIndividuals(population []domain.Individual, c int) ([]domain.Individual, []domain.Individual) {
	removed := make([]domain.Individual, 0)

	// Remove inviable individuals
	for i := 0; i < len(population); i += 1 {
		population[i].Resources -= c
		population[i].History = append(population[i].History, population[i].Resources)

		if population[i].Resources < c {
			removed = append(removed, population[i])

			population = append(population[:i], population[i+1:]...)
		}
	}

	return population, removed
}
*/

// Simulate a population surviving based in amount of energy each individuo consumption,
// how much energy individuos exchange, the probability of an individuo be salidary
// with other with less energy when exchange energy and tha amount of individuos
// the society need to have to be viable, viability threshold.
func SimulatedSociety(population []domain.Individual, individualConsumption int, individualExchange int, salidaryProbability float32, viabilityAmount int) int {
	var iterations int
	var removed []domain.Individual

	deaths := make([]domain.Individual, 0)

	// Run the simulation while the society is viable.
	for iterations = 0; len(population) > viabilityAmount; iterations += 1 {
		population, removed = exchangeResource(population, individualConsumption, individualExchange, salidaryProbability)

		deaths = append(deaths, removed...)
	}

	// TODO: Send dead individuals to save information of simulation dynamics

	return iterations
}

func Run(simulationData domain.SimulationData, viabilityRange domain.Range, solidarityRange domain.Range) []domain.SimulationData {
	surface := make([]domain.SimulationData, 0)

	population := domain.MakePopulation(simulationData.SocietySize, simulationData.IndividualConsumption)

	dataPoit := domain.SimulationData{
		IndividualConsumption: simulationData.IndividualConsumption,
		IndividualExchange:    simulationData.IndividualExchange,
		SocietySize:           simulationData.SocietySize,
		ViabilityAmount:       0,
		SolidarityProbability: 0,
		Ephocs:                0,
	}

	for viability := viabilityRange.Minimum.(int); viability <= viabilityRange.Maximum.(int); viability += viabilityRange.Delta.(int) {
		for solidarity := solidarityRange.Minimum.(float32); solidarity <= solidarityRange.Maximum.(float32); solidarity += solidarityRange.Delta.(float32) {
			dataPoit.ViabilityAmount = viability
			dataPoit.SolidarityProbability = solidarity
			dataPoit.Ephocs = SimulatedSociety(population, dataPoit.IndividualConsumption, dataPoit.IndividualExchange, dataPoit.SolidarityProbability, dataPoit.ViabilityAmount)

			surface = append(surface, dataPoit)
		}
	}

	return surface
}
