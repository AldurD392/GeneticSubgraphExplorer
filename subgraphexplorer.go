/*
Package main
Genetic algorithm to find the smallest dense-enough subgraph.
*/
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aldur/subgraphexplorer/graph"
)

// Profiling tools
// import (
// "net/http"
// _ "net/http/pprof"
// )

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: subgraphexplorer INPUT_FILE")
		return
	}

	start := time.Now()
	inputFile := os.Args[1]

	g := graph.Parse(inputFile)
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
