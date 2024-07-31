package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// # Comparable

func IsSame[T comparable](value1, value2 T) bool {

	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestComparable(t *testing.T) {

	assert.True(t, IsSame[string]("dua", "dua"))
	assert.True(t, IsSame[int](100, 100))
	assert.True(t, IsSame[bool](true, true))
}
