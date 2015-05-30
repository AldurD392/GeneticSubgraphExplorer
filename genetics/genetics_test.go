package genetics

import (
	"testing"
)

func TestInitialPopulation(t *testing.T) {
	p := InitialPopulation(25, 3, 10)

	if len(p) != 3 {
		t.Errorf("Got wrong population size: %d", len(p))
	}
}

func TestSample(t *testing.T) {
	v := Sample(10, 3)

	if len(v.Allels) != 10 {
		t.Errorf("Got a random sample of the wrong size: %d", len(v.Allels))
	}

	if v.Size != 3 {
		t.Errorf("Got a sampled population of the wrong size: %d", v.Size)
	}
}
