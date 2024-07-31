package golanggenerics

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

// # In Line Type Constraint

type Gelombang float64

// menggunakan inline jika case nya simple atau sederhana
func FindMin[T interface {
	int | int32 | float32 | ~float64
}](first, second T) T {

	if first < second {
		return first
	} else {
		return second
	}
}

func TestInlineConstraint(t *testing.T) {

	assert.Equal(t, int(100), FindMin[int](100, 200))
	assert.Equal(t, float32(12.5), FindMin(float32(12.5), float32(12.6)))
	assert.Equal(t, Gelombang(200.32), FindMin[Gelombang](Gelombang(200.32), Gelombang(233.12)))
}
