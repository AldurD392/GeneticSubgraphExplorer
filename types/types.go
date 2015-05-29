package types

import (
	"bytes"
	"fmt"
	"strings"
)

/* Data types */
type Graph struct {
	AdjacencyMap AdjacencyMap
	Labels       UIntSlice // Our i-th node was originally Labels[i]
}
type AdjacencyMap map[uint32]UIntSlice
type UIntSlice []uint32

/* String() interfaces */
func (g AdjacencyMap) String() string {
	var s bytes.Buffer

	s.WriteString("{\n")
	for key, value := range g {
		s.WriteString(fmt.Sprintf("\t%d: %s,\n", key, value))
	}
	s.WriteString("}")

	return s.String()
}

func (labels UIntSlice) String() string {
	var s bytes.Buffer

	s.WriteString("[")
	for _, n := range labels {
		s.WriteString(fmt.Sprintf("%d, ", n))
	}

	return strings.TrimRight(s.String(), ", ") + "]"
}
