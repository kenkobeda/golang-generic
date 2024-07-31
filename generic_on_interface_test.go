package golanggenerics

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

// # Generic Interface

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type MyData[T any] struct {
	Value T
}

func (m *MyData[T]) GetValue() T {
	return m.Value
}

func (m *MyData[T]) SetValue(value T) {
	m.Value = value
}

func TestGenericInterface(t *testing.T) {

	myData := &MyData[string]{
		Value: "perdana",
	}

	result := ChangeValue[string](myData, "Budi")

	assert.Equal(t, "Budi", result)
	assert.Equal(t, "Budi", myData.Value)
}
