/*
Package main
Genetic algorithm to find the smallest dense-enough subgraph.
*/
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"github.com/aldur/subgraphexplorer/types"
	"log"
	"os"
	"strconv"
	"strings"
)

// Profiling tools
// import (
//     "net/http"
//     _ "net/http/pprof"
// )

/* Open the graph file for reading and build the structure. */
func readInputFile(path string) *types.Graph {
	var (
		index         uint32             = 0
		adjacencyMap  types.AdjacencyMap = make(types.AdjacencyMap)
		nodesToLabels types.IntToIntMap  = make(types.IntToIntMap)
		labelsToNodes types.IntToIntMap  = make(types.IntToIntMap)
	)

	inputFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			// Skip comment lines
			continue
		}

		edge := strings.Fields(line)

		u_64, err := strconv.ParseUint(edge[0], 10, 32)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		u := uint32(u_64)
		u_index, ok := nodesToLabels[u]
		if !ok {
			nodesToLabels[u] = index
			labelsToNodes[index] = u
			u_index = index
			index += 1
		}

		v_64, err := strconv.ParseUint(edge[1], 10, 32)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		v := uint32(v_64)
		v_index, ok := nodesToLabels[v]
		if !ok {
			nodesToLabels[v] = index
			labelsToNodes[index] = v
			v_index = index
			index += 1
		}

		l, ok := adjacencyMap[u_index]
		if !ok {
			l = list.New()
			adjacencyMap[u_index] = l
		}
		l.PushBack(v_index)

		l, ok = adjacencyMap[v_index]
		if !ok {
			l = list.New()
			adjacencyMap[v_index] = l
		}
		l.PushBack(u_index)
	}

	return &types.Graph{adjacencyMap, nodesToLabels}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: subgraphexplorer INPUT_FILE")
		return
	}

	inputFile := os.Args[1]
	g := readInputFile(inputFile)
	if g == nil {
		log.Panicln("Cannot parse input file. Exiting...")
	}

	fmt.Println(g.AdjacencyMap)
	fmt.Println(g.LabelsToNodes)

	// Enable profiling
	// log.Println(http.ListenAndServe("localhost:6060", nil))
}
