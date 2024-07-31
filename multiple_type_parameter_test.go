package golanggenerics

import (
	"fmt"
	"testing"
)

// # Multiple Type Parameter

func MultipleParameter[T any, K any](param1 T, param2 K) {

	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleParameter(t *testing.T) {

	MultipleParameter[string, int]("budi", 22)
	MultipleParameter[int, string](109, "wallet")
}
