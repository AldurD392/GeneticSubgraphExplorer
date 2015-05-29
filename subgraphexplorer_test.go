package main

import (
	"github.com/aldur/subgraphexplorer/types"
	"testing"
)

func TestReading(t *testing.T) {
	graph := readInputFile("./test/smallClique.txt")
	if len(graph.AdjacencyMap) != 11 {
		t.Errorf("Got wrong smallClique graph size: %d", len(graph.AdjacencyMap))
	}

	if len(graph.Labels) != 11 {
		t.Errorf("Got wrong smallClique labels size: %d", len(graph.Labels))
	}

	v, ok := graph.AdjacencyMap[2]
	if !ok || len(v) != 1 || v[0] != 0 {
		t.Errorf("Got wrong gamma(2): %s", v)
	}

	v, ok = graph.AdjacencyMap[7]
	gamma_v := []uint32{5, 6, 10, 4, 8, 9}
	if !ok || len(v) != len(gamma_v) {
		t.Errorf("Got wrong gamma(7): %s", v)
	}

	for i := range v {
		if v[i] != gamma_v[i] {
			t.Errorf("Got wrong gamma(7): %s", v)
			break
		}
	}
}

func TestDensity(t *testing.T) {
	graph := readInputFile("./test/smallClique.txt")
	cases := []struct {
		subgraph types.UIntSlice
		density  float32
	}{
		{types.UIntSlice{5}, 0},
		{types.UIntSlice{1, 2}, 0},
		{types.UIntSlice{5, 6, 7, 8, 9}, 2.0},
		{types.UIntSlice{5, 6, 7, 8, 9, 10}, 2.0},
	}

	for _, c := range cases {
		got := density(c.subgraph, graph.AdjacencyMap)
		if got != c.density {
			t.Errorf("density(%s) == %f, want %f", c.subgraph, got, c.density)
		}
	}
}
