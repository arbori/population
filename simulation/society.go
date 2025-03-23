package simulation

import "math/rand"

// Rage is a configuration for the algorithm iteration.
type Range struct {
	Minimum interface{}
	Maximum interface{}
	Delta   interface{}
}

// This is the parameters that will use to simulate the iteration between individuals.
// IndividualConsumption is the amount of resources an individual consume in each iteration.
// IndividualExchange is the amount of resources two individual exchanges in each iteration.
// SocietySize is the initial size of society.
// ViabilityAmount is the minimmum resource for individual be alive.
// SolidarityProbability is the probability that instead an strong individual  get resource from an weak one, it will give an IndividualExchange number of resource instead.
type SocietyParameters struct {
	IndividualConsumption int
	IndividualExchange    int
	SocietySize           int
	ViabilityAmount       int
	SolidarityProbability float32
}

// SocietyRealization save the information of algorithm running.
// Society is the parameters of execution.
// Ephocs is the number of iterations.
// Deaths is the individuals dead in the end of the simulation.
type SocietyRealization struct {
	Society SocietyParameters
	Ephocs  int
	Deaths  []Individual
}

// Running the simulation.
func (s SocietyRealization) Run(viabilityRange Range, solidarityRange Range) []SocietyRealization {
	surface := make([]SocietyRealization, 0)

	population := MakeSociety(s.Society.SocietySize, s.Society.IndividualConsumption)

	for viability := viabilityRange.Minimum.(int); viability <= viabilityRange.Maximum.(int); viability += viabilityRange.Delta.(int) {
		for solidarity := solidarityRange.Minimum.(float32); solidarity <= solidarityRange.Maximum.(float32); solidarity += solidarityRange.Delta.(float32) {
			realization := SocietyRealization{
				Society: SocietyParameters{
					IndividualConsumption: s.Society.IndividualConsumption,
					IndividualExchange:    s.Society.IndividualExchange,
					SocietySize:           s.Society.SocietySize,
					ViabilityAmount:       viability,
					SolidarityProbability: solidarity,
				},
				Ephocs: 0,
				Deaths: make([]Individual, 0),
			}

			simulatedSociety(population, &realization)

			surface = append(surface, realization)

		}
	}

	return surface
}

// Simulate a population surviving based in amount of energy each individuo consumption,
// how much energy individuos exchange, the probability of an individuo be salidary
// with other with less energy when exchange energy and tha amount of individuos
// the society need to have to be viable, viability threshold.
func simulatedSociety(population []Individual, realization *SocietyRealization) {
	var iterations int
	var removed []Individual

	deaths := make([]Individual, 0)

	// Run the simulation while the society is viable.
	for iterations = 0; len(population) > realization.Society.ViabilityAmount; iterations += 1 {
		population, removed = exchangeResource(population, realization.Society.IndividualConsumption, realization.Society.IndividualExchange, realization.Society.SolidarityProbability)

		deaths = append(deaths, removed...)
	}

	// TODO: Send dead individuals to save information of simulation dynamics
}

// Have each individual exchange resources with someone else. The rule is: In the
// exchange between two individuals, those who have the most receive and those who
// have the least give.
func exchangeResource(population []Individual, individualConsumption int, individualExchange int, solidarityProbability float32) ([]Individual, []Individual) {
	newPopulation := make([]Individual, 0, len(population))
	var notViable []Individual

	removed := make([]Individual, 0)

	// Make exchange while more than two individuals did not exchange yet.
	for len(population) >= 2 {
		firstIndex, secondIndex := chooseIndividualsIndexes(len(population))

		ExchangeResourceRule(&population[firstIndex], &population[secondIndex], individualExchange, solidarityProbability)

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

// Moves individual to the new population if it is still viable after resource consumption.
func moveViablesIndividuals(newPopulation []Individual, IndividualConsumption int, individuals ...*Individual) ([]Individual, []Individual) {
	notViable := make([]Individual, 0)

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
func individualConsumption(individual *Individual, IndividualConsumption int) {
	individual.Resources -= IndividualConsumption
	individual.History = append(individual.History, individual.Resources)
}
