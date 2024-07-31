package golanggenerics

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

// # # Type Approximation

type Age int

type Number2 interface {
	~int | int8 | int16 | int32 | int64 | float32 | float64
}

func Max[T Number2](first T, second T) T {

	if first > second {
		return first
	} else {
		return second
	}
}

func TestTypeApproximation(t *testing.T) {

	var angka int = Min[int](int(100), int(200))
	fmt.Println(angka)

	assert.Equal(t, int(200), Max[int](100, 200))
	assert.Equal(t, int64(200), Max[int64](int64(100), int64(200)))
	assert.Equal(t, float64(200.0), Max[float64](100.0, 200.0))
	assert.Equal(t, Age(200), Max[Age](Age(200), Age(100)))
}

// # Type Inference

func TestTypeInference(t *testing.T) {

	var angka int = Min[int](int(100), int(200))
	fmt.Println(angka)

	assert.Equal(t, int(200), Max(100, 200))
	assert.Equal(t, int64(200), Max(int64(100), int64(200)))
	assert.Equal(t, float64(200.0), Max(100.0, 200.0))
	assert.Equal(t, Age(200), Max(Age(200), Age(100)))
}
