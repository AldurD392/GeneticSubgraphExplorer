/*
Package subgraphexplorer
Genetic algorithm to find the smallest dense-enough subgraph.
*/
package main

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
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

/* Data types */
type graph struct {
	adjacencyMap  adjacencyMap
	labelsToNodes intToIntMap
}
type adjacencyMap map[uint32]*list.List
type intToIntMap map[uint32]uint32

/* String() interfaces */
func (g adjacencyMap) String() string {
	var s bytes.Buffer

	s.WriteString("{\n")
	for key, value := range g {
		s.WriteString(fmt.Sprintf("\t%d: [", key))

		for e := value.Front(); e != nil; e = e.Next() {
			if e.Next() != nil {
				s.WriteString(fmt.Sprintf("%d, ", e.Value))
			} else {
				s.WriteString(fmt.Sprintf("%d", e.Value))
			}
		}

		// s = strings.TrimRight(s, ", ")
		s.WriteString("],\n")
	}
	// s = strings.TrimRight(s, ",\n\t")
	s.WriteString("}")

	return s.String()
}

func (m intToIntMap) String() string {
	var s bytes.Buffer

	s.WriteString("{\n")
	for k, v := range m {
		s.WriteString(fmt.Sprintf("\t%d -> %d,\n", k, v))
	}
	s.WriteString("}")

	return s.String()
}

/* Open the graph file for reading and build the structure. */
func readInputFile(path string) *graph {
	var (
		index         uint32       = 0
		adjacencyMap  adjacencyMap = make(adjacencyMap)
		nodesToLabels intToIntMap  = make(intToIntMap)
		labelsToNodes intToIntMap  = make(intToIntMap)
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

	return &graph{adjacencyMap, nodesToLabels}
}

func main() {
	inputFile := os.Args[1]
	g := readInputFile(inputFile)
	if g == nil {
		log.Panicln("Cannot parse input file. Exiting...")
	}

	fmt.Println(g.adjacencyMap)
	fmt.Println(g.labelsToNodes)

	// Enable profiling
	// log.Println(http.ListenAndServe("localhost:6060", nil))

}
