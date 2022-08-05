package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/arbori/population.git/population/individuo"
	"github.com/arbori/population.git/population/simulation"
)

func main() {
	fmt.Println("Simulated society")
	
	var individualConsumption int = 2000
	var individualExchange int = individualConsumption
	var viabilityThreshold int = 500

	var societySize = 100000
	var salidaryProbability float32
	var epochs int

	I := make([]individuo.Individuo, societySize)

	rand.Seed(int64(time.Now().Nanosecond()))

	// Sortear a quantidade de recursos para os indivíduos de I.
	for i := range I {
		I[i] = individuo.MakeIndividuo(individualConsumption * rand.Intn(100))
	}

	fmt.Printf("\nsolidarity\texistent\n")

	for salidaryProbability = .0; salidaryProbability <= 1; salidaryProbability += .01 {
		epochs = simulation.SimulatedSociety(I, individualConsumption, individualExchange, salidaryProbability, viabilityThreshold)

		fmt.Printf("%f\t%d\n", salidaryProbability, epochs)
	}
}
