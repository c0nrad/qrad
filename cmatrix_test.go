package qrad

import (
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

	for h := 0; h < g.Height; h++ {
		for w := 0; w < g.Width; w++ {
			if !s.At(w, h).Equals(g.At(w, h)) {
				t.Error("Failed to take tensor product")
			}
		}
	}
}
