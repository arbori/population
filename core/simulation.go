package simulation

import (
	"math/rand"
)

// Have each individual exchange resources with someone else. The rule is: In the
// exchange between two individuals, those who have the most receive and those who
// have the least give.
func ExchangeResource(I []Individuo, e int, s float32) []Individuo {
	P := make([]Individuo, 0, len(I))

	// Make exchange while more than two individous did not exchange yet.
	for len(I) >= 2 {
		ra := rand.Intn(len(I))
		rb := rand.Intn(len(I))

		// The individuos must to be diferent
		for ra == rb {
			rb = rand.Intn(len(I))
		}

		// Make exchange between ra and rb.
		solidarity := rand.Float32()

		if (s < solidarity && I[ra].Resources > I[rb].Resources) || (s > solidarity && I[ra].Resources < I[rb].Resources) {
			I[ra].Resources += e
			I[rb].Resources -= e
		} else if (s > solidarity && I[ra].Resources > I[rb].Resources) || (s < solidarity && I[ra].Resources < I[rb].Resources) {
			I[ra].Resources -= e
			I[rb].Resources += e
		}

		// Save the values.
		P = append(P, I[ra], I[rb])

		// Move to new population P the individuos that already exchanges energy.
		if ra < rb {
			I = append(append(I[:ra], I[ra+1:rb]...), I[rb+1:]...)
		} else {
			I = append(append(I[:rb], I[rb+1:ra]...), I[ra+1:]...)
		}
	}

	// Set the new population
	if len(I) > 0 {
		P = append(P, I...)
	}

	return P
}

// Remove from population the inviable individuos, returning the new population set
// and the individuos removed.
func RemoveInviableIndividuos(I []Individuo, c int) ([]Individuo, []Individuo) {
	removed := make([]Individuo, 0)

	// Remove inviable individuos
	for i := 0; i < len(I); i += 1 {
		I[i].Resources -= c
		I[i].History = append(I[i].History, I[i].Resources)

		if I[i].Resources < c {
			removed = append(removed, I[i])

			I = append(I[:i], I[i+1:]...)
		}
	}

	return I, removed
}

// Simulate I society surviving based in amount of energy each individuo consumption,
// how much energy individuos exchange, the probability of an individuo be salidary
// with other with less energy when exchange energy and tha amount of individuos
// the society need to have to be viable, viability threshold.
func SimulatedSociety(I []Individuo, individualConsumption int, individualExchange int, salidaryProbability float32, viabilityThreshold int) int {
	var iterations int

	removed := make([]Individuo, 0)
	deaths := make([]Individuo, 0)

	// Run the simulation while the society is viable.
	for iterations = 0; len(I) > viabilityThreshold; iterations += 1 {
		I = ExchangeResource(I, individualExchange, salidaryProbability)

		// Remove inviable individuos
		I, removed = RemoveInviableIndividuos(I, individualConsumption)

		deaths = append(deaths, removed...)
	}

	return iterations
}
