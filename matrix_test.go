package qrad

import (
	"fmt"
	"testing"
)

func TestMatrixTensorProduct(t *testing.T) {
	a := NewMatrixFromElements([][]Complex{
		{complex(2, 0)},
		{complex(3, 0)}})

	b := NewMatrixFromElements([][]Complex{
		{complex(4, 0)},
		{complex(6, 0)},
		{complex(3, 0)}})

	s := NewMatrixFromElements([][]Complex{
		{complex(8, 0)},
		{complex(12, 0)},
		{complex(6, 0)},
		{complex(12, 0)},
		{complex(18, 0)},
		{complex(9, 0)}})

	g := NewMatrix()
	g.TensorProduct(*a, *b)

	if !s.Equals(*g) {
		t.Error("Failed to take tensor product")
	}

}

func TestCMatrxAdd(t *testing.T) {
	sum := NewMatrix()
	sum.Resize(1, 2)

	a := NewMatrixFromElements([][]Complex{
		{complex(5, 0)},
		{complex(0, 10)}})

	sol := NewMatrixFromElements([][]Complex{
		{complex(15, 0)},
		{complex(0, 30)}})

	if !sum.Add(*sum, *a).Add(*sum, *a).Add(*sum, *a).Equals(*sol) {
		fmt.Println(sum, sol)
		t.Error("Failed to self add 3 times")
	}
}

func TestMatrixTranspose(t *testing.T) {
	a := NewMatrixFromElements([][]Complex{
		{complex(8, 0)},
		{complex(12, 0)}})

	b := NewMatrixFromElements([][]Complex{
		{complex(8, 0), complex(12, 0)}})

	if !a.Transpose(*a).Equals(*b) {
		t.Error("Failed to transpose")
	}
}

func TestMatrixAdjoint(t *testing.T) {
	m := CCNotGate.Matrix

	out := NewMatrix().Adjoint(m)
	out.Adjoint(*out)

	if !out.Equals(m) {
		t.Error("failed to take double Adjoint")
	}
}
