package golanggenerics

import (
	"fmt"
	"testing"
)

// # Generic Type

type Bag[T interface{}] []T

func PrintBag[T any](bag Bag[T]) {

	for _, v := range bag {

		fmt.Println(v)

	}
}

func TestBag(t *testing.T) {

	numbers := Bag[int]{1, 2, 3, 4, 5}

	PrintBag(numbers)

	names := Bag[string]{"Budi", "Indro", "Joko"}
	PrintBag[string](names)
}
