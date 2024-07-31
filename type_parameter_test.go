package golanggenerics

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/go-playground/assert/v2"
)

// # Type Parameter

func Length[T interface{}](param T) T {

	fmt.Println(reflect.TypeOf(param))
	return param
}

func TestTypeParameter(t *testing.T) {

	var angka int = 22
	assert.Equal(t, 22, Length(angka))

}
