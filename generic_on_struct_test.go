package golanggenerics

import (
	"fmt"
	"testing"
)

// # Generic Struct

type Data[T any] struct {
	First  T
	Second T
}

func TestGenericStruct(t *testing.T) {

	data := Data[string]{
		First:  "budi",
		Second: "joko",
	}

	fmt.Printf("%s", data)

}
