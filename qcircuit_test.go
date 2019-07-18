package qrad

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestQCircuitBellState(t *testing.T) {
	solution := NewCVectorFromElements([]Complex{
		Complex(complex(1/math.Sqrt(2), 0)),
		Complex(complex(0, 0)),
		Complex(complex(0, 0)),
		Complex(complex(1/math.Sqrt(2), 0)),
	})

	for i := 0; i < 100; i++ {
		q := NewQCircuit([]int{0, 0})
		q.ApplyHadamard(0)
		q.ApplyCNot(0, 1)

		for i := range solution.Elements {
			if !q.State.At(i).Equals(solution.At(i)) {
				t.Error("failed to construct entangled state")
			}
		}

		out := q.Measure()
		if out != 0 && out != 3 {
			t.Error("failed to construct bell state")
		}
	}
}

func TestQCircuitGHZState(t *testing.T) {
	solution := NewCVectorFromElements([]Complex{
		Complex(complex(1/math.Sqrt(2), 0)),
		Complex(complex(0, 0)),
		Complex(complex(0, 0)),
		Complex(complex(0, 0)),
		Complex(complex(0, 0)),
		Complex(complex(0, 0)),
		Complex(complex(0, 0)),
		Complex(complex(-1/math.Sqrt(2), 0)),
	})

	for i := 0; i < 100; i++ {
		q := NewQCircuit([]int{0, 0, 0})

		q.ApplyGate(NewCMatrix().TensorProducts(*HadamardGate, *HadamardGate, *NotGate))
		q.ApplyCNot(1, 2)
		q.ApplyCNot(0, 2)
		q.ApplyGate(NewCMatrix().TensorProducts(*HadamardGate, *HadamardGate, *HadamardGate))

		for i := range solution.Elements {
			if !q.State.At(i).Equals(solution.At(i)) {
				t.Error("failed to construct entangled state")
			}
		}

		out := q.Measure()
		if out != 0 && out != 7 {
			t.Error("failed to construct bell state")
		}
	}
}

func TestGateOrder(t *testing.T) {
	q0 := NewQCircuit([]int{0, 0})
	q0.ApplyHadamard(0)
	q0.ApplyCNot(0, 1)
	q0.ApplyGate(NewCMatrix().TensorProduct(*HadamardGate, *SGate))

	q1 := NewQCircuit([]int{0, 0})
	q1.ApplyHadamard(0)
	q1.ApplyCNot(0, 1)
	q1.ApplyHadamard(0)
	q1.ApplyS(1)

	if !q0.State.Equals(q1.State) {
		t.Error("ordering matters!")
	}
}

func TestBellInequalityExperiment(t *testing.T) {
	resultsZW := make(map[int]int)
	resultsZV := make(map[int]int)
	resultsXW := make(map[int]int)
	resultsXV := make(map[int]int)

	runs := 1000

	for i := 0; i < runs; i++ {
		// H .
		//   + S H T H
		q := NewQCircuit([]int{0, 0})
		q.ApplyHadamard(0)
		q.ApplyCNot(0, 1)
		q.ApplyS(1)
		q.ApplyHadamard(1)
		q.ApplyT(1)
		q.ApplyHadamard(1)
		out := q.Measure()
		resultsZW[out]++
	}

	for i := 0; i < runs; i++ {
		// H .
		//   + S H T' H
		q := NewQCircuit([]int{0, 0})
		q.ApplyHadamard(0)
		q.ApplyCNot(0, 1)
		q.ApplyS(1)
		q.ApplyHadamard(1)
		q.ApplyTd(1)
		q.ApplyHadamard(1)
		out := q.Measure()
		resultsZV[out]++
	}

	for i := 0; i < runs; i++ {
		// H . H
		//   + S H T H
		q := NewQCircuit([]int{0, 0})
		q.ApplyHadamard(0)
		q.ApplyCNot(0, 1)
		q.ApplyGate(NewCMatrix().TensorProduct(*HadamardGate, *SGate))
		q.ApplyHadamard(1)
		q.ApplyT(1)
		q.ApplyHadamard(1)
		out := q.Measure()
		resultsXW[out]++
	}

	for i := 0; i < runs; i++ {
		// H . H
		//   + S H T' H
		q := NewQCircuit([]int{0, 0})
		q.ApplyHadamard(0)
		q.ApplyCNot(0, 1)
		q.ApplyGate(NewCMatrix().TensorProduct(*HadamardGate, *SGate))
		q.ApplyHadamard(1)
		q.ApplyTd(1)
		q.ApplyHadamard(1)
		out := q.Measure()
		resultsXV[out]++
	}

	aZW := float64(resultsZW[0]+resultsZW[3]-resultsZW[2]-resultsZW[1]) / float64(runs)
	aZV := float64(resultsZV[0]+resultsZV[3]-resultsZV[2]-resultsZV[1]) / float64(runs)
	aXW := float64(resultsXW[0]+resultsXW[3]-resultsXW[2]-resultsXW[1]) / float64(runs)
	aXV := float64(resultsXV[0]+resultsXV[3]-resultsXV[2]-resultsXV[1]) / float64(runs)

	// S = <A,B> - <A,B'> + <A',B'> + <A',B'>
	S := aXW - aXV + aZW + aZV

	if S <= 2 {
		fmt.Println(aZW, aZV, aXW, aXV)
		fmt.Println(S)

		t.Error("quantum mechanics has failed")
	}
}

func TestQCircuitGHZExperiment(t *testing.T) {

	resultsYYX := make(map[int]int)
	resultsYXY := make(map[int]int)
	resultsXYY := make(map[int]int)

	runs := 1000

	for i := 0; i < runs; i++ {
		q := NewQCircuit([]int{0, 0, 0})
		q.ApplyGate(NewCMatrix().TensorProducts(*HadamardGate, *HadamardGate, *NotGate))
		q.ApplyCNot(1, 2)
		q.ApplyCNot(0, 2)
		q.ApplyGate(NewCMatrix().TensorProducts(*HadamardGate, *HadamardGate, *HadamardGate))

		// -- Sd H
		// -- Sd H
		// --    H
		q.ApplyGate(NewCMatrix().TensorProducts(*SdGate, *SdGate, *IdentityGate))
		q.ApplyGate(NewCMatrix().TensorProducts(*HadamardGate, *HadamardGate, *HadamardGate))

		out := q.Measure()
		resultsYYX[out]++
	}

	for i := 0; i < runs; i++ {
		q := NewQCircuit([]int{0, 0, 0})
		q.ApplyGate(NewCMatrix().TensorProducts(*HadamardGate, *HadamardGate, *NotGate))
		q.ApplyCNot(1, 2)
		q.ApplyCNot(0, 2)
		q.ApplyGate(NewCMatrix().TensorProducts(*HadamardGate, *HadamardGate, *HadamardGate))

		// -- Sd H
		// --    H
		// -- Sd H
		q.ApplyGate(NewCMatrix().TensorProducts(*SdGate, *IdentityGate, *SdGate))
		q.ApplyGate(NewCMatrix().TensorProducts(*HadamardGate, *HadamardGate, *HadamardGate))

		out := q.Measure()
		resultsYXY[out]++
	}

	for i := 0; i < runs; i++ {
		q := NewQCircuit([]int{0, 0, 0})
		q.ApplyGate(NewCMatrix().TensorProducts(*HadamardGate, *HadamardGate, *NotGate))
		q.ApplyCNot(1, 2)
		q.ApplyCNot(0, 2)
		q.ApplyGate(NewCMatrix().TensorProducts(*HadamardGate, *HadamardGate, *HadamardGate))

		// --    H
		// -- Sd H
		// -- Sd H
		q.ApplyGate(NewCMatrix().TensorProducts(*IdentityGate, *SdGate, *SdGate))
		q.ApplyGate(NewCMatrix().TensorProducts(*HadamardGate, *HadamardGate, *HadamardGate))

		out := q.Measure()
		resultsXYY[out]++
	}

	// If Y's are SAME, X=1, if Y's are different, X=0
	if resultsYYX[BinToInt("011")]+resultsYYX[BinToInt("101")]+
		resultsYYX[BinToInt("110")]+resultsYYX[BinToInt("000")] != runs {
		t.Error("Failed to prove GHZ experiment")
	}

	if resultsYXY[BinToInt("011")]+resultsYXY[BinToInt("101")]+
		resultsYXY[BinToInt("110")]+resultsYXY[BinToInt("000")] != runs {
		t.Error("Failed to prove GHZ experiment")
	}

	if resultsXYY[BinToInt("011")]+resultsXYY[BinToInt("101")]+
		resultsXYY[BinToInt("110")]+resultsXYY[BinToInt("000")] != runs {
		t.Error("Failed to prove GHZ experiment")
	}
}

func BinToInt(s string) int {
	o, _ := strconv.ParseInt(s, 2, 64)
	return int(o)
}
