package qrad

import (
	"fmt"
	"math"
	"testing"
)

func TestVectorAdd(t *testing.T) {
	a1 := []Complex{
		complex(6, -4), complex(7, 3), complex(4.2, -8.1), complex(0, -3),
	}

	b1 := []Complex{
		complex(16, 2.3), complex(0, -7), complex(6, 0), complex(0, -4),
	}

	s1 := []Complex{
		complex(22, -1.7), complex(7, -4), complex(10.2, -8.1), complex(0, -7),
	}

	av1 := Vector{Elements: a1}
	bv1 := Vector{Elements: b1}
	sv1 := NewVector()
	sv1.Add(av1, bv1)

	for i := range s1 {
		if !s1[i].Equals(sv1.At(i)) {
			t.Log(s1[i], sv1.At(i))
			t.Error("Failed to add matrix")
		}
	}
}

func TestVectorSub(t *testing.T) {
	a1 := []Complex{
		complex(6, -4), complex(7, 3), complex(4.2, -8.1), complex(0, -3),
	}

	s1 := []Complex{
		complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0),
	}

	av1 := Vector{Elements: a1}
	sv1 := NewVector()
	sv1.Sub(av1, av1)

	for i := range s1 {
		if !s1[i].Equals(sv1.At(i)) {
			t.Log(s1[i], sv1.At(i))
			t.Error("Failed to add matrix")
		}
	}
}

func TestVectorMulScalar(t *testing.T) {
	a1 := []Complex{
		complex(6, 3), complex(0, 0), complex(5, 1), complex(4, 0),
	}

	b1 := Complex(complex(3, 2))

	s1 := []Complex{
		complex(12, 21), complex(0, 0), complex(13, 13), complex(12, 8),
	}

	av1 := Vector{Elements: a1}
	sv1 := NewVector()
	sv1.MulScalar(b1, av1)

	for i := range s1 {
		if !s1[i].Equals(sv1.At(i)) {
			t.Log(s1[i], sv1.At(i))
			t.Error("Failed to add matrix")
		}
	}
}

func TestVectorMulMatrix(t *testing.T) {
	a1 := [][]Complex{
		[]Complex{complex(4, 0), complex(-1, 0)},
		[]Complex{complex(2, 0), complex(1, 0)},
	}

	b1 := []Complex{complex(1, 0), complex(2, 0)}

	s1 := []Complex{
		complex(2, 0), complex(4, 0),
	}

	am1 := NewMatrixFromElements(a1)
	bv1 := NewVectorFromElements(b1)
	sv1 := NewVector()
	sv1.MulMatrix(*bv1, *am1)

	for i := range s1 {
		if !s1[i].Equals(sv1.At(i)) {
			t.Log(s1[i], sv1.At(i))
			t.Error("Failed to mul matrix")
		}
	}
}

func TestVectorTensorProduct(t *testing.T) {
	a := NewVectorFromElements([]Complex{complex(2, 0), complex(3, 0)})

	b := NewVectorFromElements([]Complex{complex(4, 0), complex(6, 0), complex(3, 0)})

	s := NewVectorFromElements([]Complex{complex(8, 0), complex(12, 0),
		complex(6, 0), complex(12, 0), complex(18, 0), complex(9, 0)})

	g := NewVector()
	g.TensorProduct(*a, *b)

	for h := 0; h < g.Length(); h++ {
		if !s.At(h).Equals(g.At(h)) {
			t.Error("Failed to take tensor product")
		}
	}

}

func TestVectorMeasure(t *testing.T) {
	// 1. Make sure bell state normalized correctly
	bellstate := NewVectorFromElements([]Complex{
		Complex(complex(1/math.Sqrt(2), 0)),
		Complex(complex(0, 0)),
		Complex(complex(0, 0)),
		Complex(complex(1/math.Sqrt(2), 0)),
	})

	if !NearEqual(bellstate.Norm(), 1) {
		fmt.Println(bellstate.Norm())
		t.Error("failed to get norm")
	}

	// Ensure probablities sum to 1
	bellstate.Probabilities()
	probSums := float64(0)
	for _, v := range bellstate.Probabilities() {
		probSums += v
	}
	if !NearEqual(probSums, 1) {
		t.Error("Failed to sum probabilities to 1")
	}

	// Ensure measurement is fair-ish
	results := make(map[int]int)
	for i := 0; i < 1000; i++ {
		results[bellstate.Measure()]++
	}

	if results[1] != 0 || results[2] != 0 {
		t.Error("bellstate shouldn't collapse to these states")
	}

	// really these should be equal...
	if results[0] == 0 || results[3] == 0 {
		t.Error("we should be getting at least a few results each")
	}
}

func TestVectorBitMeasure(t *testing.T) {

	results := make(map[int]int)

	for i := 0; i < 1; i++ {
		q := NewCircuit([]int{0, 0, 0})
		q.ApplyGate(ExtendGateFill([]int{0, 1, 2}, 3, HadamardGate))

		// q.State.PrintProbabilities()
		results[q.State.MeasureQubit(0)]++
		// q.State.PrintProbabilities()
	}
	// fmt.Println(results)

}
