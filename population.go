package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Have each individual exchange resources with someone else. The rule is: In the
// exchange between two individuals, those who have the most receive and those who
// have the least give.
func ExchangeResource(I []int, e int, s float32) []int {
	P := make([]int, 0, len(I))

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

func RemoveInviableIndividuos(I []int, c int) []int {
	// Remove inviable individuos
	for i := 0; i < len(I); i += 1 {
		I[i] -= c

		if I[i] < c {
			I = append(I[:i], I[i+1:]...)
		}
	}

	return I
}

// Simulate I society surviving based in amount of energy each individuo consume (c),
// how much energy individuos exchange (e), the probability of an individuo be salidary
// with other with less energy when exchange energy (s) and tha amount of individuos
// the society need to have to be viable (n).
func SimulatedSociety(I []int, c int, e int, s float32, n int) int {
	var iterations int
	var deathsAmount int

	deaths := make([]int, 0)

	// Run the simulation while the society is viable.
	for iterations = 0; len(I) > n; iterations += 1 {
		I = ExchangeResource(I, e, s)

		// Remove inviable individuos
		I = RemoveInviableIndividuos(I, c)

		deaths = append(deaths, deathsAmount)
	}

	return iterations
}

func main() {
	fmt.Println("Simulated society")

	I := make([]int, 1000)
	var c int = 2000
	var e int = c
	var s float32
	var n int = 2

	var existent int

	rand.Seed(int64(time.Now().Nanosecond()))

	// Sortear a quantidade de recursos para os indivíduos de I.
	for i := range I {
		I[i] = c * rand.Intn(100)
	}

	fmt.Printf("\nsolidarity\texistent\n")

	for s = .0; s <= 1; s += .01 {
		existent = SimulatedSociety(I, c, e, s, n)

		fmt.Printf("%f\t%d\n", s, existent)
	}
}
