package exp

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"

	"golang.org/x/exp/constraints"
)

// ## Constraints Package

func ExperimentalMin[T constraints.Ordered](first, second T) T {

	if first < second {
		return first
	} else {
		return second
	}
}

func TestExperimentalMin(t *testing.T) {

	assert.Equal(t, 100, ExperimentalMin[int](100, 200))
	assert.Equal(t, 100.20, ExperimentalMin(100.20, 200.20))
}

// # Maps & Slices Packages

func TestExperimentalMaps(t *testing.T) {

	first := map[string]string{
		"Name": "Budi",
	}

	second := map[string]string{
		"Name": "Budi",
	}

	assert.True(t, maps.Equal(first, second), "Harus true hasilnya")
}

// SLICE

func TestExperimentalSlice(t *testing.T) {

	first := []string{"Budi"}
	second := []string{"Budi"}
	assert.True(t, slices.Equal(first, second))
}
