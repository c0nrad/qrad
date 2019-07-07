package qrad

import (
	"testing"
)

func TestCVectorAdd(t *testing.T) {
	a1 := []Complex{
		complex(6, -4), complex(7, 3), complex(4.2, -8.1), complex(0, -3),
	}

	b1 := []Complex{
		complex(16, 2.3), complex(0, -7), complex(6, 0), complex(0, -4),
	}

	s1 := []Complex{
		complex(22, -1.7), complex(7, -4), complex(10.2, -8.1), complex(0, -7),
	}

	av1 := CVector{Elements: a1}
	bv1 := CVector{Elements: b1}
	sv1 := NewCVector()
	sv1.Add(av1, bv1)

	for i := range s1 {
		if !s1[i].Equals(sv1.At(i)) {
			t.Log(s1[i], sv1.At(i))
			t.Error("Failed to add matrix")
		}
	}
}

func TestCVectorSub(t *testing.T) {
	a1 := []Complex{
		complex(6, -4), complex(7, 3), complex(4.2, -8.1), complex(0, -3),
	}

	s1 := []Complex{
		complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0),
	}

	av1 := CVector{Elements: a1}
	sv1 := NewCVector()
	sv1.Sub(av1, av1)

	for i := range s1 {
		if !s1[i].Equals(sv1.At(i)) {
			t.Log(s1[i], sv1.At(i))
			t.Error("Failed to add matrix")
		}
	}
}

func TestCVectorMulScalar(t *testing.T) {
	a1 := []Complex{
		complex(6, 3), complex(0, 0), complex(5, 1), complex(4, 0),
	}

	b1 := Complex(complex(3, 2))

	s1 := []Complex{
		complex(12, 21), complex(0, 0), complex(13, 13), complex(12, 8),
	}

	av1 := CVector{Elements: a1}
	sv1 := NewCVector()
	sv1.MulScalar(b1, av1)

	for i := range s1 {
		if !s1[i].Equals(sv1.At(i)) {
			t.Log(s1[i], sv1.At(i))
			t.Error("Failed to add matrix")
		}
	}
}

func TestCVectorMulMatrix(t *testing.T) {
	a1 := [][]Complex{
		[]Complex{complex(4, 0), complex(-1, 0)},
		[]Complex{complex(2, 0), complex(1, 0)},
	}

	b1 := []Complex{complex(1, 0), complex(2, 0)}

	s1 := []Complex{
		complex(2, 0), complex(4, 0),
	}

	am1 := NewCMatrixFromElements(a1)
	bv1 := NewCVectorFromElements(b1)
	sv1 := NewCVector()
	sv1.MulMatrix(*bv1, *am1)

	for i := range s1 {
		if !s1[i].Equals(sv1.At(i)) {
			t.Log(s1[i], sv1.At(i))
			t.Error("Failed to mul matrix")
		}
	}
}

func TestCVectorTensorProduct(t *testing.T) {
	a := NewCVectorFromElements([]Complex{complex(2, 0), complex(3, 0)})

	b := NewCVectorFromElements([]Complex{complex(4, 0), complex(6, 0), complex(3, 0)})

	s := NewCVectorFromElements([]Complex{complex(8, 0), complex(12, 0),
		complex(6, 0), complex(12, 0), complex(18, 0), complex(9, 0)})

	g := NewCVector()
	g.TensorProduct(*a, *b)

	for h := 0; h < g.Length(); h++ {
		if !s.At(h).Equals(g.At(h)) {
			t.Error("Failed to take tensor product")
		}
	}

}
