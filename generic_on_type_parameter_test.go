package golanggenerics

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

// # Generic di Type Parameter

func GetFirst[T []E, E any](data T) E {
	first := data[0]
	return first
}

func TestGenericTypeParamter(t *testing.T) {

	names := []string{
		"Budi", "Joko", "Indro",
	}

	first := GetFirst[[]string, string](names)

	assert.Equal(t, "Budi", first)
}
