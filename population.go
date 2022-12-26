package main

import (
	"fmt"
	"math/rand"
	"time"

	domain "github.com/arbori/population.git/population/domain"
	simulation "github.com/arbori/population.git/population/usecase"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	fmt.Println("Simulated society")

	simulationData := domain.SimulationData{
		IndividualConsumption: 2000,
		IndividualExchange:    2000,
		SocietySize:           50000,
		ViabilityAmount:       0,
		SolidarityProbability: 0.0,
		Ephocs:                0,
	}

	viabilityRange := domain.Range{
		Minimum: 500,
		Maximum: 10500,
		Delta:   500,
	}

	solidarityRange := domain.Range{
		Minimum: float32(0.0),
		Maximum: float32(1.0),
		Delta:   float32(0.1),
	}

	simulationSurface := simulation.Run(simulationData, viabilityRange, solidarityRange)

	fmt.Printf("IndividualConsumption: %d\nIndividualExchange: %d\nSocietySize: %d\n", simulationData.IndividualConsumption, simulationData.IndividualExchange, simulationData.SocietySize)
	fmt.Println("ViabilityAmount\tSolidarityProbability\tEphocs")

	for i, data := range simulationSurface {
		fmt.Printf("%d\t%d\t%.2f\t%d\n", i, data.ViabilityAmount, data.SolidarityProbability, data.Ephocs)
	}
}
