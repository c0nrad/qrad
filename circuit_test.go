package qrad

import (
	"math"
	"testing"
)

func TestCircuitHadamardSimple(t *testing.T) {
	solution := NewVectorFromElements([]Complex{
		Complex(complex(1/math.Sqrt(2), 0)),
		Complex(complex(0, 0)),
		Complex(complex(1/math.Sqrt(2), 0)),
		Complex(complex(0, 0)),
	})

	for i := 0; i < 100; i++ {
		c := NewCircuit([]int{0, 0})
		c.Append(H, 0)
		c.Execute()

		for i := range solution.Elements {
			if !c.State.At(i).Equals(solution.At(i)) {
				t.Error("failed to construct entangled state")
			}
		}

		out := c.Measure()
		if out != 0 && out != 2 {
			t.Error("failed to construct bell state")
		}
	}
}

func TestBellStateReversed(t *testing.T) {
	c1 := NewCircuit([]int{1, 0})
	c1.Append(H, 0)
	c1.AppendControl(X, 0, 1)
	c1.Execute()

	c2 := NewCircuit([]int{0, 1})
	c2.Append(H, 1)
	c2.AppendControl(X, 1, 0)
	c2.Execute()

	if !c1.State.Equals(c2.State) {
		c1.Draw()
		c2.Draw()
		c1.PrintStates()
		c2.PrintStates()
		t.Error("reversed bell state incorrect")
	}
}

func TestCCNOTMoment(t *testing.T) {
	c1 := NewCircuit([]int{1, 0})
	c1.Append(H, 0)
	c1.AppendControl(X, 0, 1)
	c1.Execute()

	c2 := NewCircuit([]int{0, 1})
	c2.Append(H, 1)
	c2.AppendControl(X, 1, 0)
	c2.Execute()

	if !c1.State.Equals(c2.State) {
		c1.Draw()
		c2.Draw()
		c1.PrintStates()
		c2.PrintStates()
		t.Error("reversed bell state incorrect")
	}
}

// func TestCircuitGHZState(t *testing.T) {
// 	solution := NewVectorFromElements([]Complex{
// 		Complex(complex(1/math.Sqrt(2), 0)),
// 		Complex(complex(0, 0)),
// 		Complex(complex(0, 0)),
// 		Complex(complex(0, 0)),
// 		Complex(complex(0, 0)),
// 		Complex(complex(0, 0)),
// 		Complex(complex(0, 0)),
// 		Complex(complex(-1/math.Sqrt(2), 0)),
// 	})

// 	for i := 0; i < 100; i++ {
// 		q := NewCircuit([]int{0, 0, 0})

// 		q.ApplyGate(NewMatrix().TensorProducts(*HadamardGate, *HadamardGate, *NotGate))
// 		q.ApplyCNot(1, 2)
// 		q.ApplyCNot(0, 2)
// 		q.ApplyGate(NewMatrix().TensorProducts(*HadamardGate, *HadamardGate, *HadamardGate))

// 		for i := range solution.Elements {
// 			if !q.State.At(i).Equals(solution.At(i)) {
// 				t.Error("failed to construct entangled state")
// 			}
// 		}

// 		out := q.Measure()
// 		if out != 0 && out != 7 {
// 			t.Error("failed to construct bell state")
// 		}
// 	}
// }

// func TestGateOrder(t *testing.T) {
// 	q0 := NewCircuit([]int{0, 0})
// 	q0.ApplyHadamard(0)
// 	q0.ApplyCNot(0, 1)
// 	q0.ApplyGate(NewMatrix().TensorProduct(*HadamardGate, *SGate))

// 	q1 := NewCircuit([]int{0, 0})
// 	q1.ApplyHadamard(0)
// 	q1.ApplyCNot(0, 1)
// 	q1.ApplyHadamard(0)
// 	q1.ApplyS(1)

// 	if !q0.State.Equals(q1.State) {
// 		t.Error("ordering matters!")
// 	}
// }

// func TestBellInequalityExperiment(t *testing.T) {
// 	resultsZW := make(map[int]int)
// 	resultsZV := make(map[int]int)
// 	resultsXW := make(map[int]int)
// 	resultsXV := make(map[int]int)

// 	runs := 1000

// 	for i := 0; i < runs; i++ {
// 		// H .
// 		//   + S H T H
// 		q := NewCircuit([]int{0, 0})
// 		q.ApplyHadamard(0)
// 		q.ApplyCNot(0, 1)
// 		q.ApplyS(1)
// 		q.ApplyHadamard(1)
// 		q.ApplyT(1)
// 		q.ApplyHadamard(1)
// 		out := q.Measure()
// 		resultsZW[out]++
// 	}

// 	for i := 0; i < runs; i++ {
// 		// H .
// 		//   + S H T' H
// 		q := NewCircuit([]int{0, 0})
// 		q.ApplyHadamard(0)
// 		q.ApplyCNot(0, 1)
// 		q.ApplyS(1)
// 		q.ApplyHadamard(1)
// 		q.ApplyTd(1)
// 		q.ApplyHadamard(1)
// 		out := q.Measure()
// 		resultsZV[out]++
// 	}

// 	for i := 0; i < runs; i++ {
// 		// H . H
// 		//   + S H T H
// 		q := NewCircuit([]int{0, 0})
// 		q.ApplyHadamard(0)
// 		q.ApplyCNot(0, 1)
// 		q.ApplyGate(NewMatrix().TensorProduct(*HadamardGate, *SGate))
// 		q.ApplyHadamard(1)
// 		q.ApplyT(1)
// 		q.ApplyHadamard(1)
// 		out := q.Measure()
// 		resultsXW[out]++
// 	}

// 	for i := 0; i < runs; i++ {
// 		// H . H
// 		//   + S H T' H
// 		q := NewCircuit([]int{0, 0})
// 		q.ApplyHadamard(0)
// 		q.ApplyCNot(0, 1)
// 		q.ApplyGate(NewMatrix().TensorProduct(*HadamardGate, *SGate))
// 		q.ApplyHadamard(1)
// 		q.ApplyTd(1)
// 		q.ApplyHadamard(1)
// 		out := q.Measure()
// 		resultsXV[out]++
// 	}

// 	aZW := float64(resultsZW[0]+resultsZW[3]-resultsZW[2]-resultsZW[1]) / float64(runs)
// 	aZV := float64(resultsZV[0]+resultsZV[3]-resultsZV[2]-resultsZV[1]) / float64(runs)
// 	aXW := float64(resultsXW[0]+resultsXW[3]-resultsXW[2]-resultsXW[1]) / float64(runs)
// 	aXV := float64(resultsXV[0]+resultsXV[3]-resultsXV[2]-resultsXV[1]) / float64(runs)

// 	// S = <A,B> - <A,B'> + <A',B'> + <A',B'>
// 	S := aXW - aXV + aZW + aZV

// 	if S <= 2 {
// 		fmt.Println(aZW, aZV, aXW, aXV)
// 		fmt.Println(S)

// 		t.Error("quantum mechanics has failed")
// 	}
// }

// func TestCircuitGHZExperiment(t *testing.T) {

// 	resultsYYX := make(map[int]int)
// 	resultsYXY := make(map[int]int)
// 	resultsXYY := make(map[int]int)

// 	runs := 1000

// 	for i := 0; i < runs; i++ {
// 		q := NewCircuit([]int{0, 0, 0})
// 		q.ApplyGate(NewMatrix().TensorProducts(*HadamardGate, *HadamardGate, *NotGate))
// 		q.ApplyCNot(1, 2)
// 		q.ApplyCNot(0, 2)
// 		q.ApplyGate(NewMatrix().TensorProducts(*HadamardGate, *HadamardGate, *HadamardGate))

// 		// -- Sd H
// 		// -- Sd H
// 		// --    H
// 		q.ApplyGate(NewMatrix().TensorProducts(*SdGate, *SdGate, *IdentityGate))
// 		q.ApplyGate(NewMatrix().TensorProducts(*HadamardGate, *HadamardGate, *HadamardGate))

// 		out := q.Measure()
// 		resultsYYX[out]++
// 	}

// 	for i := 0; i < runs; i++ {
// 		q := NewCircuit([]int{0, 0, 0})
// 		q.ApplyGate(NewMatrix().TensorProducts(*HadamardGate, *HadamardGate, *NotGate))
// 		q.ApplyCNot(1, 2)
// 		q.ApplyCNot(0, 2)
// 		q.ApplyGate(NewMatrix().TensorProducts(*HadamardGate, *HadamardGate, *HadamardGate))

// 		// -- Sd H
// 		// --    H
// 		// -- Sd H
// 		q.ApplyGate(NewMatrix().TensorProducts(*SdGate, *IdentityGate, *SdGate))
// 		q.ApplyGate(NewMatrix().TensorProducts(*HadamardGate, *HadamardGate, *HadamardGate))

// 		out := q.Measure()
// 		resultsYXY[out]++
// 	}

// 	for i := 0; i < runs; i++ {
// 		q := NewCircuit([]int{0, 0, 0})
// 		q.ApplyGate(NewMatrix().TensorProducts(*HadamardGate, *HadamardGate, *NotGate))
// 		q.ApplyCNot(1, 2)
// 		q.ApplyCNot(0, 2)
// 		q.ApplyGate(NewMatrix().TensorProducts(*HadamardGate, *HadamardGate, *HadamardGate))

// 		// --    H
// 		// -- Sd H
// 		// -- Sd H
// 		q.ApplyGate(NewMatrix().TensorProducts(*IdentityGate, *SdGate, *SdGate))
// 		q.ApplyGate(NewMatrix().TensorProducts(*HadamardGate, *HadamardGate, *HadamardGate))

// 		out := q.Measure()
// 		resultsXYY[out]++
// 	}

// 	// If Y's are SAME, X=1, if Y's are different, X=0
// 	if resultsYYX[BinToInt("011")]+resultsYYX[BinToInt("101")]+
// 		resultsYYX[BinToInt("110")]+resultsYYX[BinToInt("000")] != runs {
// 		t.Error("Failed to prove GHZ experiment")
// 	}

// 	if resultsYXY[BinToInt("011")]+resultsYXY[BinToInt("101")]+
// 		resultsYXY[BinToInt("110")]+resultsYXY[BinToInt("000")] != runs {
// 		t.Error("Failed to prove GHZ experiment")
// 	}

// 	if resultsXYY[BinToInt("011")]+resultsXYY[BinToInt("101")]+
// 		resultsXYY[BinToInt("110")]+resultsXYY[BinToInt("000")] != runs {
// 		t.Error("Failed to prove GHZ experiment")
// 	}
// }

// func BinToInt(s string) int {
// 	o, _ := strconv.ParseInt(s, 2, 64)
// 	return int(o)
// }

// func TestQCiruitOrReversable(t *testing.T) {
// 	for x := 0; x < 2; x++ {
// 		for y := 0; y < 2; y++ {
// 			for z := 0; z < 2; z++ {

// 				q := NewCircuit([]int{x, y, z})
// 				q.ApplyGate(ExtendGate(2, 3, NotGate))
// 				q.ApplyOrGate(0, 1, 2)
// 				q.ApplyOrGate(0, 1, 2)
// 				q.ApplyGate(ExtendGate(2, 3, NotGate))

// 				out := q.Measure()
// 				if out != (x<<2)+(y<<1)+(z<<0) {
// 					fmt.Println(x, y, z, fmt.Sprintf("%03b", out))
// 					t.Error("failed to show OR is reversable")
// 				}
// 			}
// 		}
// 	}
// }

// func TestCircuit3Or(t *testing.T) {
// 	for x := 0; x < 2; x++ {
// 		for y := 0; y < 2; y++ {
// 			for z := 0; z < 2; z++ {

// 				//0    o
// 				//1    o
// 				//2        o
// 				//3  X +   o
// 				//4      X +
// 				q := NewCircuit([]int{x, y, z, 0, 0})
// 				q.ApplyGate(ExtendGate(3, 5, NotGate))
// 				q.ApplyOrGate(0, 1, 3)

// 				q.ApplyGate(ExtendGate(4, 5, NotGate))
// 				q.ApplyOrGate(2, 3, 4)

// 				out := q.Measure()

// 				if out&1 != Or(Or(x, y), z) {
// 					fmt.Println(x, y, z, (x+y+z)&1, fmt.Sprintf("%05b", out))
// 					t.Error("failed to perform 3 gate OR")
// 				}

// 				q2 := NewCircuit([]int{x, y, z, 0, 0})
// 				q2.Apply3OrGate(0, 1, 2, 3, 4)
// 				out2 := q2.Measure()
// 				if out2 != out {
// 					t.Error("3OrGate did not work")
// 				}
// 			}
// 		}
// 	}
// }

// func Or(x, y int) int {
// 	if x == 1 {
// 		return 1
// 	}

// 	if y == 1 {
// 		return 1
// 	}

// 	return 0
// }
