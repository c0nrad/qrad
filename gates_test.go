package qrad

import (
	"math"
	"testing"
)

func TestConstructIdentity(t *testing.T) {
	if !ConstructNIdentity(1).Matrix.Equals(I.Matrix) {
		t.Error("failed to construct 1-Identiy")
	}

	identity2 := NewMatrix().TensorProduct(I.Matrix, I.Matrix)
	if !identity2.Equals(ConstructNIdentity(2).Matrix) {
		t.Error("failed to construct 2-identity")
	}

}

func TestBellStateConstruction(t *testing.T) {

	// Prepare two qubits |00>
	state0 := NewVector()
	state0.TensorProduct(*NewQubit(0), *NewQubit(0))

	// Hadamard
	// ---[ H ]
	// --------
	operator1 := NewMatrix()
	operator1.TensorProduct(HadamardGate.Matrix, IdentityGate.Matrix)

	state1 := NewVector()
	state1.MulMatrix(*state0, *operator1)

	// CNOT
	// -----.--
	// ----(X)-
	operator2 := CNOT
	state2 := NewVector()
	state2.MulMatrix(*state1, operator2.Matrix)

	solution := NewVectorFromElements([]Complex{
		Complex(complex(1/math.Sqrt(2), 0)),
		Complex(complex(0, 0)),
		Complex(complex(0, 0)),
		Complex(complex(1/math.Sqrt(2), 0)),
	})

	for i := range solution.Elements {
		if !state2.At(i).Equals(solution.At(i)) {
			t.Error("failed to construct entangled state")
		}
	}
}

// 1, 0, 2 1
func encodeCharacter(r rune) []int {
	out := []int{}

	for i := 0; i < 4; i++ {
		out = append(out, int(r%4))
		r >>= 2
	}
	return out
}

func decodeCharacter(in []int) string {
	sum := byte(0)

	for i := 3; i >= 0; i-- {
		sum <<= 2
		sum += byte(in[i])
	}
	return string(sum)
}

func TestQuantumSuperDenseCoding(t *testing.T) {

	messageStr := "hello world!!"
	message := []int{}
	for _, c := range messageStr {
		message = append(message, encodeCharacter(rune(c))...)
	}

	out := ""
	buff := []int{}

	for _, m := range message {
		// Construct Bellstate
		state0 := NewVector()
		state0.TensorProduct(*NewQubit(0), *NewQubit(0))

		operator1 := NewMatrix()
		operator1.TensorProduct(HadamardGate.Matrix, IdentityGate.Matrix)
		state1 := NewVector()
		state1.MulMatrix(*state0, *operator1)

		operator2 := CNOT
		state2 := NewVector()
		state2.MulMatrix(*state1, operator2.Matrix)

		// Alice encoders her info
		aliceState := NewVector()
		switch m {
		case 0:
			operator3 := NewMatrix()
			operator3.TensorProduct(IdentityGate.Matrix, IdentityGate.Matrix)
			aliceState.MulMatrix(*state2, *operator3)
		case 1:
			operator3 := NewMatrix()
			operator3.TensorProduct(PauliXGate.Matrix, IdentityGate.Matrix)
			aliceState.MulMatrix(*state2, *operator3)
		case 2:
			operator3 := NewMatrix()
			operator3.TensorProduct(PauliZGate.Matrix, IdentityGate.Matrix)
			aliceState.MulMatrix(*state2, *operator3)
		case 3:
			operator3 := NewMatrix()
			operator3.TensorProduct(PauliZGate.Matrix, IdentityGate.Matrix)

			operator4 := NewMatrix()
			operator4.TensorProduct(PauliXGate.Matrix, IdentityGate.Matrix)

			aliceState.MulMatrix(*state2, *operator3)
			aliceState.MulMatrix(*aliceState, *operator4)
		}

		// Bob decoders the info
		state3 := NewVector()
		state3.MulMatrix(*aliceState, CNOT.Matrix)

		state4 := NewVector()
		operator5 := NewMatrix()
		operator5.TensorProduct(HadamardGate.Matrix, IdentityGate.Matrix)

		state4.MulMatrix(*state3, *operator5)

		if state4.Measure() != m {
			t.Error("Failed to encode information")
		}

		buff = append(buff, m)

		if len(buff) == 4 {
			out += decodeCharacter(buff)
			buff = []int{}
		}
	}

	if out != messageStr {
		t.Error("Failed to decode message")
	}
}

func TestExtendControlGate(t *testing.T) {
	// ----.----
	// ----X----
	cnot2 := ExtendControlGate(0, 1, 2, X)

	if !cnot2.Matrix.Equals(CNOT.Matrix) {
		t.Error("Failed to build standard CNOT gate")
	}

	// ----.----
	// ----|----
	// ----X----
	cnot3 := ExtendControlGate(0, 2, 3, X)
	cnot3sol := NewMatrixFromElements([][]Complex{
		[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0))},
	})

	if !cnot3.Matrix.Equals(*cnot3sol) {
		t.Error("Failed to build extended 3 bit CNOT gate")
	}
}

// func TestExtendControlGateSimple(t *testing.T) {
// 	cnot2 := ExtendControlGate(0, 1, 2, NotGate)
// }

func TestExtendGate(t *testing.T) {
	operator1 := NewMatrix()
	operator1.TensorProduct(HadamardGate.Matrix, IdentityGate.Matrix)

	hadamard2 := ExtendGate(0, 2, HadamardGate)

	if !hadamard2.Matrix.Equals(*operator1) {
		t.Error("Failed to construct gate")
	}
}

func TestExtendGates2(t *testing.T) {
	operator1 := NewMatrix()

	operator1.TensorProducts(CNOT.Matrix, IdentityGate.Matrix)
	// operator1.PPrint()
}

// func TestExtendGateFill(t *testing.T) {
// 	qubits := 5
// 	q1 := NewCircuit(make([]int, qubits))
// 	for i := 0; i < qubits; i++ {
// 		q.ApplyHadamard(i)
// 	}
// 	q.ApplyGate(ExtendGateFill([]int{0, 1, 2, 3, 4, 5}, q.Qubits, HadamardGate))

// ?}

func TestToffoliGate(t *testing.T) {
	for x := 0; x < 2; x++ {
		for y := 0; y < 2; y++ {
			for z := 0; z < 2; z++ {
				q := NewCircuit([]int{x, y, z})
				q.ApplyGate(ToffoliGate)

				out := q.Measure()

				zSol := (x & y) ^ z
				sol := (zSol << 0) + (y << 1) + (x << 2)
				if sol != out {
					t.Error("Failed to solve toffoli gate")
				}

			}
		}
	}
}

func TestExtendControlControlGateSimple(t *testing.T) {
	toffoliGuess := ExtendControlControlGate(0, 1, 2, 3, NotGate)
	if !toffoliGuess.Matrix.Equals(ToffoliGate.Matrix) {
		toffoliGuess.Matrix.PPrint()
		t.Error("Failed to construct toffoli gate")
	}
}

func TestExtendControlControlGateReversed(t *testing.T) {
	toffoliGuess := ExtendControlControlGate(1, 2, 0, 3, NotGate)

	for x := 0; x < 2; x++ {
		for y := 0; y < 2; y++ {
			for z := 0; z < 2; z++ {
				q := NewCircuit([]int{z, x, y})
				q.ApplyGate(toffoliGuess)

				out := q.Measure()

				zSol := (x & y) ^ z
				sol := (zSol << 2) + (y << 0) + (x << 1)
				if sol != out {
					t.Error("Failed to solve toffoli gate")
				}

			}
		}
	}
}

func TestExtendControlControlGate(t *testing.T) {
	for x := 0; x < 2; x++ {
		for y := 0; y < 2; y++ {
			for z := 0; z < 2; z++ {
				q := NewCircuit([]int{x, y, 0, z})
				q.ApplyGate(ExtendControlControlGate(0, 1, 3, 4, NotGate))

				out := q.Measure()

				zSol := (x & y) ^ z
				sol := (zSol << 0) + (y << 2) + (x << 3)
				if sol != out {
					t.Error("Failed to solve toffoli gate")
				}

			}
		}
	}
}
