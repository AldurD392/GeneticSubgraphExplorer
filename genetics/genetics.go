// Package genetics provides genetics-related tools
package genetics

import (
	"math/rand"
	"time"
)

/*
Generate a random solution (i.e. a subgraph) as a boolean vector.
v[i] is True is i-th node is part of the subgraph.
*/
func Sample(subgraphLen int, samplingSize int) (individual Individual) {
	rand.Seed(time.Now().UnixNano())

	individual = Individual{
		make([]bool, subgraphLen),
		0,
	}

	p := rand.Perm(subgraphLen)
	for i := 0; i < samplingSize; i++ {
		individual.Allels[p[i]] = true
	}
	individual.Size = samplingSize

	return individual
}

func InitialPopulation(subgraphLen int,
	populationSize int, samplingSize int) (population []Individual) {
	population = make([]Individual, populationSize)

	for i := 0; i < populationSize; i++ {
		population[i] = Sample(subgraphLen, samplingSize)
	}

	return population
}

/*
Fitness returns a 0 <= f <= 1
Its 0 if the result is not a solution at all (doesn't meet required density)
Its 1 if the result is optimal (a clique)
*/
func Fitness(individual Individual) float32 {
	// TODO
	return 0.0
}
