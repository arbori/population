package main

import (
	"fmt"
	"math/rand"
)

// Simulate I society surviving based in amount of energy each individuo consume (c),
// how much energy individuos exchange (e), the probability of an individuo be salidary
// with other with less energy when exchange energy (s) and tha amount of individuos
// the society need to have to be viable (n).
func SimulatedSociety(I []float32, c float32, e float32, s float32, n int) int {
	var iterations int

	// Run the simulation while the society is viable.
	for iterations = 0; len(I) > n; iterations += 1 {
		P := make([]float32, 0, len(I))

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

			if (s < solidarity && I[ra] > I[rb]) || (s > solidarity && I[ra] < I[rb]) {
				I[ra] += e
				I[rb] -= e
			} else if (s > solidarity && I[ra] > I[rb]) || (s < solidarity && I[ra] < I[rb]) {
				I[ra] -= e
				I[rb] += e
			}

			// Save the values.
			P = append(P, I[ra], I[rb])

			// Remove from population individuos that already exchanges energy.
			if ra < rb {
				I = append(append(I[:ra], I[ra+1:rb]...), I[rb+1:]...)
			} else {
				I = append(append(I[:rb], I[rb+1:ra]...), I[ra+1:]...)
			}
		}

		// Set the population for next iteration
		if len(I) > 0 {
			P = append(P, I...)
		}

		I = P

		// Remove inviable individuos
		for i := 0; i < len(I); i += 1 {
			I[i] -= c

			if I[i] < c {
				I = append(I[:i], I[i+1:]...)
			}
		}
	}

	return iterations
}

func main() {
	fmt.Println("Simulated society")

	I := make([]float32, 100)
	var c float32 = 2000
	var e float32 = c
	var s float32 = .12
	var n int = 2

	// Sortear a quantidade de recursos para os indivíduos de I.
	for i := range I {
		I[i] = 100.0 * c * rand.Float32()
	}

	fmt.Printf("\nsolidarity\texistent\n")

	for s = .0; s <= 1; s += .01 {
		existent := SimulatedSociety(I, c, e, s, n)

		fmt.Printf("%f\t%d\n", s, existent)
	}
}
