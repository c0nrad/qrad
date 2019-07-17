package qrad

import (
	"fmt"
	"testing"
)

func TestCMatrixTensorProduct(t *testing.T) {
	a := NewCMatrixFromElements([][]Complex{
		[]Complex{complex(2, 0)},
		[]Complex{complex(3, 0)}})

	b := NewCMatrixFromElements([][]Complex{
		[]Complex{complex(4, 0)},
		[]Complex{complex(6, 0)},
		[]Complex{complex(3, 0)}})

	s := NewCMatrixFromElements([][]Complex{
		[]Complex{complex(8, 0)},
		[]Complex{complex(12, 0)},
		[]Complex{complex(6, 0)},
		[]Complex{complex(12, 0)},
		[]Complex{complex(18, 0)},
		[]Complex{complex(9, 0)}})

	g := NewCMatrix()
	g.TensorProduct(*a, *b)

	if !s.Equals(*g) {
		t.Error("Failed to take tensor product")
	}

}

func TestCMatrxAdd(t *testing.T) {
	sum := NewCMatrix()
	sum.Resize(1, 2)

	a := NewCMatrixFromElements([][]Complex{
		[]Complex{complex(5, 0)},
		[]Complex{complex(0, 10)}})

	sol := NewCMatrixFromElements([][]Complex{
		[]Complex{complex(15, 0)},
		[]Complex{complex(0, 30)}})

	if !sum.Add(*sum, *a).Add(*sum, *a).Add(*sum, *a).Equals(*sol) {
		fmt.Println(sum, sol)
		t.Error("Failed to self add 3 times")
	}
}

func TestCMatrixTranspose(t *testing.T) {
	a := NewCMatrixFromElements([][]Complex{
		[]Complex{complex(8, 0)},
		[]Complex{complex(12, 0)}})

	b := NewCMatrixFromElements([][]Complex{
		[]Complex{complex(8, 0), complex(12, 0)}})

	if !a.Transpose(*a).Equals(*b) {
		t.Error("Failed to transpose")
	}
}
