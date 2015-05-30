package graph

import (
	"testing"
)

const (
	test_input_file = "../test/smallClique.txt"
)

func TestOptimalSize(t *testing.T) {
	cases := []struct {
		density float32
		result  int
	}{
		{2.0, 5},
		{3.0, 7},
		{3.5, 8},
		{4.0, 9},
	}
	for _, c := range cases {
		if OptimalSize(c.density) != c.result {
			t.Errorf("Got wrong optimalSize for %f: %d", c.density, OptimalSize(c.density))
		}
	}
}

func TestDensity(t *testing.T) {
	graph := Parse(test_input_file)
	cases := []struct {
		subgraph UIntSlice
		density  float32
	}{
		{UIntSlice{5}, 0},
		{UIntSlice{1, 2}, 0},
		{UIntSlice{5, 6, 7, 8, 9}, 2.0},
		{UIntSlice{5, 6, 7, 8, 9, 10}, 2.0},
	}

	for _, c := range cases {
		got := Density(c.subgraph, graph.AdjacencyMap)
		if got != c.density {
			t.Errorf("density(%s) == %f, want %f", c.subgraph, got, c.density)
		}
	}
}

func TestParse(t *testing.T) {
	graph := Parse(test_input_file)

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
