package golanggenerics

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

// # Type Parameter Inheritance

type Employee interface {
	GetName() string
}

func GetName[T Employee](parameter T) string {
	return parameter.GetName()
}

// MANAGER

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

// VICE PRESIDENT

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (m *MyVicePresident) GetName() string {
	return m.Name
}

func (m *MyVicePresident) GetVicePresidentName() string {
	return m.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Budi", GetName[Manager](&MyManager{Name: "Budi"}))
	assert.Equal(t, "Budi", GetName[VicePresident](&MyVicePresident{Name: "Budi"}))
}
