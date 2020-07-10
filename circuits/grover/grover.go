package main

import (
	"fmt"

	"github.com/c0nrad/qrad"
)

func main() {
	// inputs := 3
	qubits := 5

	results := make(map[int]int)
	runs := 100
	for i := 0; i < runs; i++ {

		q := qrad.NewCircuit(make([]int, qubits))

		q.Append(qrad.X, []int{4})

		q.Append(qrad.H, []int{0, 1, 2, 4})
		// q.ApplyGate(ExtendGateFill([]int{0, 1, 2}, q.Qubits, HadamardGate))

		// q.Append(qrad.X, []int{0, 1, 2, 3, 4})

		q.AppendBarrier()
		for r := 0; r < 1; r++ {
			f3(q)
			f3Reverse(q, false)
			q.AppendBarrier()
			grover_mover3(q)
		}

		q.Draw()
		q.Execute()

		out := q.State.Measure() >> 2
		results[out]++
		fmt.Println(out)
	}
	// fmt.Println(results)

	maxProbability := float64(0)
	maxValue := 0
	fmt.Println("Grover Results f(x) = 1, x =...")
	for k, v := range results {
		fmt.Printf("%03b %.02f\n", k, float64(v)/float64(runs))

		if float64(v)/float64(runs) > maxProbability {
			maxValue = k
		}
	}

	if maxValue != 5 {
		fmt.Println("success, measured 101")
	} else {
		fmt.Println("not quite right...., 101 should have the highest probabilit")
	}

}

func f3(q *qrad.Circuit) {
	// 5
	// 1 a
	// 0
	// 1

	// (a & !b) & c
	// (a & !b)
	q.Append(qrad.X, []int{1})
	q.AppendControl(qrad.X, []int{0, 1}, 3)
	q.Append(qrad.X, []int{1})

	// (3) & c => 5
	q.AppendControl(qrad.PauliZGate, []int{2, 3}, 4)
}

func f3Reverse(q *qrad.Circuit, includeFinal bool) {
	if includeFinal {
		q.AppendControl(qrad.X, []int{2, 3}, 4)
	}
	q.Append(qrad.X, []int{1})
	q.AppendControl(qrad.X, []int{0, 1}, 3)
	q.Append(qrad.X, []int{1})
}

// func TestOracle3Correctness(t *testing.T) {
// 	i := 0
// 	for a := 0; a < 2; a++ {
// 		for b := 0; b < 2; b++ {
// 			for c := 0; c < 2; c++ {
// 				q := qrad.NewCircuit([]int{a, b, c, 0, 0})

// 				f3(q)
// 				out := q.Measure()

// 				// fmt.Println(a, b, c, d, fmt.Sprintf("%07b", out))
// 				if a == 1 && b == 0 && c == 1 {
// 					if out&1 != 1 {

// 						t.Error("oracle didn't work")
// 					}
// 				} else {
// 					if out&1 != 0 {
// 						t.Error("oracle gave a yes on a bad input")
// 					}
// 				}

// 				i++
// 			}
// 		}
// 	}
// }

// func TestOracleReverse(t *testing.T) {

// 	// oracles := []func(*Circuit){f}
// 	// reverseoracles := []func(*Circuit, bool){fReverse}

// 	for a := 0; a < 2; a++ {
// 		for b := 0; b < 2; b++ {
// 			for c := 0; c < 2; c++ {
// 				q := qrad.NewCircuit([]int{a, b, c, 0, 0})
// 				f3(q)
// 				f3Reverse(q, true)

// 				out := q.Measure()
// 				if out != (a<<4)+(b<<3)+(c<<2) {
// 					t.Error("failed to reverse")
// 				}

// 				q2 := NewCircuit([]int{a, b, c, 0, 0})
// 				f3(q2)
// 				f3Reverse(q2, false)
// 				out2 := q2.Measure()

// 				if a == 1 && b == 0 && c == 1 {
// 					if out2 != (a<<4)+(b<<3)+(c<<2)+1 {
// 						fmt.Println(a, b, c, fmt.Sprintf("%05b", out2))
// 						t.Error("oracle didn't work")
// 					}
// 				} else {
// 					if out2&1 != 0 {
// 						t.Error("oracle gave a yes on a bad input")
// 					}
// 				}
// 			}
// 		}

// 	}
// }

func grover_mover3(c *qrad.Circuit) {

	c.Append(qrad.H, []int{0, 1, 2})
	c.Append(qrad.X, []int{0, 1, 2})

	c.AppendControl(qrad.PauliZGate, []int{0, 1}, 2)

	c.Append(qrad.X, []int{0, 1, 2})
	c.Append(qrad.H, []int{0, 1, 2})
}
