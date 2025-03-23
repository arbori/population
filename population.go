package main

import (
	"fmt"
	"math/rand"
	"time"

	simulation "github.com/arbori/population.git/population/simulation"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	fmt.Println("Simulated society")

	society := simulation.SocietyParameters{
		IndividualConsumption: 2000,
		IndividualExchange:    2000,
		SocietySize:           50000,
		ViabilityAmount:       0,
		SolidarityProbability: 0.0,
	}

	simulationData := simulation.SocietyRealization{
		Society: society,
		Ephocs:  0,
	}

	viabilityRange := simulation.Range{
		Minimum: 500,
		Maximum: 10500,
		Delta:   500,
	}

	solidarityRange := simulation.Range{
		Minimum: float32(0.0),
		Maximum: float32(1.0),
		Delta:   float32(0.1),
	}

	simulationSurface := simulationData.Run(viabilityRange, solidarityRange)

	fmt.Printf("IndividualConsumption: %d\nIndividualExchange: %d\nSocietySize: %d\n", simulationData.Society.IndividualConsumption, simulationData.Society.IndividualExchange, simulationData.Society.SocietySize)
	fmt.Println("ViabilityAmount\tSolidarityProbability\tEphocs")

	for i, data := range simulationSurface {
		fmt.Printf("%d\t%d\t%.2f\t%d\n", i, data.Society.ViabilityAmount, data.Society.SolidarityProbability, data.Ephocs)
	}
}
