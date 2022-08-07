package simulation

// Individou of population.
type Individuo struct {
	Resources int
	History   []int
}

// Create a individuo with amount of resource and history of transactions.
func MakeIndividuo(resources int) Individuo {
	return Individuo{
		Resources: resources,
		History:   make([]int, 0),
	}
}
