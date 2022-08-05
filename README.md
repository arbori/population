# Population
This project try to shed light on what could be the kind of relation between individuals to the society strong possibility to survive and continue indefinitely.

# Version 1
The Population project seeks to shed light on what the relationship between individuals in a society must have so that it has a high probability of thriving even in moments of impact on the livelihood elements of the analyzed society. In the first version of this simulation, social relations were reduced to the exchange of resources between their individuals. 

The environment considered for the simulation is one where it is not possible to generate new resources and inevitably all the individuals will perish, but before, the society itself becomes unfeasible for having the number below the minimum necessary for its viability. At this point, the simulation ends, and it is possible to evaluate the parameters that allowed it to survive for the indicated time.

For the exchange between individuals, the following metaphor was considered, an individual with more resources obtains resources from the one who has less. However, this relationship can be inverted according to the probability that the individual with more resources be solidary and in fact, gives resources to the one who has less.

In the first version these parameter was considered fix, because based in some biological facts:

| Parameter  | Description  |
|---|---|
|  individual consumption | How much resources a individuo need to consume to be alive per day. It is 2000 Kcal/day. |
|  individual exchange | It is the amount of resource individuals exchange when they relate to each other. Was assumed that it is the same the individual consumption | 
| viability threshold | It is the minimal size of society to its viability. It was assumed be 500. |

The variable of simulation is:

| Variable | Description |
|----------| ------------|
| society size | The number of individuals in the begining of simulation. |
| salidary probability | The probability that a individual with more resources give to one with less. |
| epochs | The number of interation that society have been viable |
