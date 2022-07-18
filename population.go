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

	I := make([]individuo.Individuo, 1000)
	var c int = 2000
	var e int = c
	var s float32
	var n int = 2

	var existent int

	rand.Seed(int64(time.Now().Nanosecond()))

	// Sortear a quantidade de recursos para os indivíduos de I.
	for i := range I {
		I[i] = individuo.MakeIndividuo(c * rand.Intn(100))
	}

	fmt.Printf("\nsolidarity\texistent\n")

	for s = .0; s <= 1; s += .01 {
		existent = simulation.SimulatedSociety(I, c, e, s, n)

		fmt.Printf("%f\t%d\n", s, existent)
	}
}
