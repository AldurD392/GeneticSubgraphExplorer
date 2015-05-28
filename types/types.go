package types

import (
	"bytes"
	"container/list"
	"fmt"
)

/* Data types */
type Graph struct {
	AdjacencyMap  AdjacencyMap
	LabelsToNodes IntToIntMap
}
type AdjacencyMap map[uint32]*list.List
type IntToIntMap map[uint32]uint32

/* String() interfaces */
func (g AdjacencyMap) String() string {
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

func (m IntToIntMap) String() string {
	var s bytes.Buffer

	s.WriteString("{\n")
	for k, v := range m {
		s.WriteString(fmt.Sprintf("\t%d -> %d,\n", k, v))
	}
	s.WriteString("}")

	return s.String()
}
