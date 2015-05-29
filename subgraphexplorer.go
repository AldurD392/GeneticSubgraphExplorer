/*
Package main
Genetic algorithm to find the smallest dense-enough subgraph.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aldur/subgraphexplorer/types"
)

// Profiling tools
// import (
// "net/http"
// _ "net/http/pprof"
// )

func density(subgraph types.UIntSlice, adjacencyMap types.AdjacencyMap) float32 {
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

func parseInput(scanner *bufio.Scanner) *types.Graph {
	var (
		adjacencyMap types.AdjacencyMap = make(types.AdjacencyMap)
		labeledNodes map[uint32]uint32  = make(map[uint32]uint32)
		labels       types.UIntSlice
	)

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

	return &types.Graph{adjacencyMap, labels}
}

/* Open the graph file for reading and build the structure. */
func readInputFile(path string) *types.Graph {
	inputFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	return parseInput(scanner)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: subgraphexplorer INPUT_FILE")
		return
	}

	start := time.Now()
	inputFile := os.Args[1]
	g := readInputFile(inputFile)
	if g == nil {
		log.Panicln("Cannot parse input file. Exiting...")
	}
	elapsed := time.Since(start)
	log.Printf("Graph input reading took %s", elapsed)

	fmt.Println(g.AdjacencyMap)
	fmt.Println(g.Labels)

	// Enable profiling
	// log.Println(http.ListenAndServe("localhost:6060", nil))
}
