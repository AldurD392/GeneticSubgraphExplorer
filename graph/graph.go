package graph

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

/* Given the density calculate the size of the optimal solution (clique) */
func OptimalSize(density float32) int {
	return int(density*2) + 1
}

/* Calculate the density of a subgraph V as
(|E(V)| / |V|) / 2
*/
func Density(subgraph UIntSlice, adjacencyMap AdjacencyMap) float32 {
	var (
		i int

		subgraph_set map[uint32]struct{} = make(map[uint32]struct{})
		taken        struct{}
	)

	for _, u := range subgraph {
		subgraph_set[u] = taken
	}

	for _, u := range subgraph {
		var gamma_u = adjacencyMap[u]

		for _, v := range gamma_u {
			if _, ok := subgraph_set[v]; ok {
				i++
			}
		}
	}

	return (float32(i) / float32(len(subgraph))) / 2.0
}

func Parse(path string) *Graph {
	inputFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer inputFile.Close()

	var (
		adjacencyMap AdjacencyMap      = make(AdjacencyMap)
		labeledNodes map[uint32]uint32 = make(map[uint32]uint32)
		labels       UIntSlice
	)

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			// Skip comment lines
			continue
		}

		edge := strings.Fields(line)

		/* Parse u and v from the string */
		u_64, err := strconv.ParseUint(edge[0], 10, 32)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		u := uint32(u_64)

		v_64, err := strconv.ParseUint(edge[1], 10, 32)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		v := uint32(v_64)

		// Check if we need to label them
		u_index, ok := labeledNodes[u]
		if !ok {
			u_index = uint32(len(labels))
			labeledNodes[u] = u_index
			labels = append(labels, u)
		}

		v_index, ok := labeledNodes[v]
		if !ok {
			v_index = uint32(len(labels))
			labeledNodes[v] = v_index
			labels = append(labels, v)
		}

		adjacencyMap[u_index] = append(adjacencyMap[u_index], v_index)
		adjacencyMap[v_index] = append(adjacencyMap[v_index], u_index)
	}

	return &Graph{adjacencyMap, labels}
}
