package types

import (
	"bytes"
	"fmt"
)

/* Data types */
type Graph struct {
	AdjacencyMap  AdjacencyMap
	LabelsToNodes IntToIntMap
}
type AdjacencyMap map[uint32][]uint32
type IntToIntMap map[uint32]uint32

/* String() interfaces */
func (g AdjacencyMap) String() string {
	var s bytes.Buffer

	s.WriteString("{\n")
	for key, value := range g {
		s.WriteString(fmt.Sprintf("\t%d: %s", key, value))
	}
	s.WriteString("}")

	return s.String()
}

func (m IntToIntMap) String() string {
	var s bytes.Buffer

	s.WriteString("{\n")
	for k, v := range m {
		s.WriteString(fmt.Sprintf("\t%d -> %d,\n", k, v))
	}
	s.WriteString("}")

	return s.String()
}
