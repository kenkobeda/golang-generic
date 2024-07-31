package golanggenerics

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

// # Generic Method

type Data2[T any] struct {
	First  T
	Second T
}

func (d *Data2[_]) SayHello(name string) string {
	// karena tidak perlu bermain dengan field pada struct
	// maka bisa tanpa type parameter, atau di ganti dengan underscore (_)
	return "Hello " + name
}

func (d *Data2[E]) ChangeFirst(first E) E {
	// jika ingin bermain dengan data field di struct
	// wajib menggunakan type parameter nya.
	d.First = first
	return d.First
}
func TestGenericStructMethod(t *testing.T) {

	data := Data2[string]{
		First:  "andy",
		Second: "indro",
	}

	assert.Equal(t, "Budi", data.ChangeFirst("Budi"))
	assert.Equal(t, "Hello Budi", data.SayHello("Budi"))

}
