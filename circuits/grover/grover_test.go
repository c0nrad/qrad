package qrad

import (
	"fmt"
	"testing"
)

func f3(q *Circuit) {
	// 5
	// 1 a
	// 0
	// 1

	// (a & !b) & c
	// (a & !b)
	q.ApplyNot(1)
	q.ApplyToffoliGate(0, 1, 3)
	q.ApplyNot(1)

	// (3) & c => 5
	q.ApplyToffoliGate(3, 2, 4)
}

func f3Reverse(q *Circuit, includeFinal bool) {
	if includeFinal {
		q.ApplyToffoliGate(3, 2, 4)
	}
	q.ApplyNot(1)
	q.ApplyToffoliGate(0, 1, 3)
	q.ApplyNot(1)
}

func TestOracle3Correctness(t *testing.T) {
	i := 0
	for a := 0; a < 2; a++ {
		for b := 0; b < 2; b++ {
			for c := 0; c < 2; c++ {
				q := NewCircuit([]int{a, b, c, 0, 0})

				f3(q)
				out := q.Measure()

				// fmt.Println(a, b, c, d, fmt.Sprintf("%07b", out))
				if a == 1 && b == 0 && c == 1 {
					if out&1 != 1 {

						t.Error("oracle didn't work")
					}
				} else {
					if out&1 != 0 {
						t.Error("oracle gave a yes on a bad input")
					}
				}

				i++
			}
		}
	}
}

func TestOracleReverse(t *testing.T) {

	// oracles := []func(*Circuit){f}
	// reverseoracles := []func(*Circuit, bool){fReverse}

	for a := 0; a < 2; a++ {
		for b := 0; b < 2; b++ {
			for c := 0; c < 2; c++ {
				q := NewCircuit([]int{a, b, c, 0, 0})
				f3(q)
				f3Reverse(q, true)

				out := q.Measure()
				if out != (a<<4)+(b<<3)+(c<<2) {
					t.Error("failed to reverse")
				}

				q2 := NewCircuit([]int{a, b, c, 0, 0})
				f3(q2)
				f3Reverse(q2, false)
				out2 := q2.Measure()

				if a == 1 && b == 0 && c == 1 {
					if out2 != (a<<4)+(b<<3)+(c<<2)+1 {
						fmt.Println(a, b, c, fmt.Sprintf("%05b", out2))
						t.Error("oracle didn't work")
					}
				} else {
					if out2&1 != 0 {
						t.Error("oracle gave a yes on a bad input")
					}
				}
			}
		}

	}
}

func grover_mover3(q *Circuit) {
	q.ApplyGate(ExtendGateFill([]int{0, 1, 2}, q.Qubits, HadamardGate))
	q.ApplyGate(ExtendGateFill([]int{0, 1, 2}, q.Qubits, NotGate))

	// q.ApplyHadamard(2)
	q.ApplyGate(ExtendControlControlGate(0, 1, 2, q.Qubits, PauliZGate))
	// q.ApplyHadamard(2)

	q.ApplyGate(ExtendGateFill([]int{0, 1, 2}, q.Qubits, NotGate))
	q.ApplyGate(ExtendGateFill([]int{0, 1, 2}, q.Qubits, HadamardGate))
}

func TestGroversAlgorithm(t *testing.T) {
	// inputs := 3
	qubits := 5

	results := make(map[int]int)
	runs := 1000
	for i := 0; i < runs; i++ {

		q := NewCircuit(make([]int, qubits))
		q.ApplyGate(ExtendGateFill([]int{0, 1, 2}, q.Qubits, HadamardGate))

		q.ApplyNot(qubits - 1)
		q.ApplyH(qubits - 1)

		for i := 0; i < 1; i++ {
			f3(q)
			f3Reverse(q, false)
			grover_mover3(q)

		}
		out := q.State.Measure() >> 2
		results[out]++
	}
	// fmt.Println(results)
	fmt.Println("Grover Results f(x) = 1, x =...")
	for k, v := range results {
		fmt.Printf("%03b %.02f\n", k, float64(v)/float64(runs))
	}
}
