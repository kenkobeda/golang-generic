package golanggenerics

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

// # Type Sets

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Min[T Number](first T, second T) T {

	if first < second {
		return first
	} else {
		return second
	}
}

func TestTypeSets(t *testing.T) {

	var angka int = Min[int](int(100), int(200))
	fmt.Println(angka)

	assert.Equal(t, int(100), Min[int](100, 200))
	assert.Equal(t, int64(100), Min[int64](int64(100), int64(200)))
	assert.Equal(t, float64(100.0), Min[float64](100.0, 200.0))
}
