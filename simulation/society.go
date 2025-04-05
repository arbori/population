package simulation

import "math/rand"

// Rage is a configuration for the algorithm iteration.
type Range[T any] struct {
	Minimum T
	Maximum T
	Delta   T
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
func (s SocietyRealization) Run(viabilityRange Range[int], solidarityRange Range[float32]) []SocietyRealization {
	surface := make([]SocietyRealization, 0)

	population := MakeSociety(s.Society.SocietySize, s.Society.IndividualConsumption)

	for viability := viabilityRange.Minimum; viability <= viabilityRange.Maximum; viability += viabilityRange.Delta {
		for solidarity := solidarityRange.Minimum; solidarity <= solidarityRange.Maximum; solidarity += solidarityRange.Delta {
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

			realization.simulatedSociety(population)

			surface = append(surface, realization)
		}
	}

	return surface
}

// Simulate a population surviving based in amount of energy each individuo consumption,
// how much energy individuos exchange, the probability of an individuo be salidary
// with other with less energy when exchange energy and tha amount of individuos
// the society need to have to be viable, viability threshold.
func (s SocietyRealization) simulatedSociety(population []Individual) {
	var iterations int
	var removed []Individual

	deaths := make([]Individual, 0)

	// Run the simulation while the society is viable.
	for iterations = 0; len(population) > s.Society.ViabilityAmount; iterations += 1 {
		population, removed = s.exchangeResource(population)

		deaths = append(deaths, removed...)
	}

	// TODO: Send dead individuals to save information of simulation dynamics
}
// Have each individual exchange resources with someone else. The rule is: In the
// exchange between two individuals, those who have the most receive and those who
// have the least give.
func (s SocietyRealization) exchangeResource(population []Individual) ([]Individual, []Individual) {
	newPopulation := make([]Individual, 0, len(population))
	var notViable []Individual

	removed := make([]Individual, 0)

	// Make exchange while more than two individuals did not exchange yet.
	for len(population) >= 2 {
		firstIndex, secondIndex := s.chooseIndividualsIndexes(len(population))

		ExchangeResourceRule(&population[firstIndex], &population[secondIndex], s.Society.IndividualExchange, s.Society.SolidarityProbability)

		// Move to new population the viables individuals and to removed the inviables.
		newPopulation, notViable = s.moveViablesIndividuals(newPopulation, &population[firstIndex], &population[secondIndex])
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
func (s SocietyRealization) chooseIndividualsIndexes(size int) (int, int) {
	firstIndex := rand.Intn(size)
	secondIndex := rand.Intn(size)

	// The individuos must to be diferent
	for firstIndex == secondIndex {
		secondIndex = rand.Intn(size)
	}

	return firstIndex, secondIndex
}

// Moves individual to the new population if it is still viable after resource consumption.
func (s SocietyRealization) moveViablesIndividuals(newPopulation []Individual, individuals ...*Individual) ([]Individual, []Individual) {
	notViable := make([]Individual, 0)

	for _, individual := range individuals {
		s.individualConsumption(individual)

		if individual.Resources >= s.Society.IndividualConsumption {
			newPopulation = append(newPopulation, *individual)
		} else {
			notViable = append(notViable, *individual)
		}
	}

	return newPopulation, notViable
}

// Individual consumption of amount
func (s SocietyRealization) individualConsumption(individual *Individual) {
	individual.Resources -= s.Society.IndividualConsumption
	individual.History = append(individual.History, individual.Resources)
}
