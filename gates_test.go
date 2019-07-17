package qrad

import (
	"math"
	"testing"
)

func TestBellStateConstruction(t *testing.T) {

	// Prepare two qubits |00>
	state0 := NewCVector()
	state0.TensorProduct(*NewQubit(0), *NewQubit(0))

	// Hadamard
	// ---[ H ]
	// --------
	operator1 := NewCMatrix()
	operator1.TensorProduct(*HadamardGate, *IdentityGate)

	state1 := NewCVector()
	state1.MulMatrix(*state0, *operator1)

	// CNOT
	// -----.--
	// ----(X)-
	operator2 := CNotGate
	state2 := NewCVector()
	state2.MulMatrix(*state1, *operator2)

	solution := NewCVectorFromElements([]Complex{
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
		state0 := NewCVector()
		state0.TensorProduct(*NewQubit(0), *NewQubit(0))

		operator1 := NewCMatrix()
		operator1.TensorProduct(*HadamardGate, *IdentityGate)
		state1 := NewCVector()
		state1.MulMatrix(*state0, *operator1)

		operator2 := CNotGate
		state2 := NewCVector()
		state2.MulMatrix(*state1, *operator2)

		// Alice encoders her info
		aliceState := NewCVector()
		switch m {
		case 0:
			operator3 := NewCMatrix()
			operator3.TensorProduct(*IdentityGate, *IdentityGate)
			aliceState.MulMatrix(*state2, *operator3)
		case 1:
			operator3 := NewCMatrix()
			operator3.TensorProduct(*PauliXGate, *IdentityGate)
			aliceState.MulMatrix(*state2, *operator3)
		case 2:
			operator3 := NewCMatrix()
			operator3.TensorProduct(*PauliZGate, *IdentityGate)
			aliceState.MulMatrix(*state2, *operator3)
		case 3:
			operator3 := NewCMatrix()
			operator3.TensorProduct(*PauliZGate, *IdentityGate)

			operator4 := NewCMatrix()
			operator4.TensorProduct(*PauliXGate, *IdentityGate)

			aliceState.MulMatrix(*state2, *operator3)
			aliceState.MulMatrix(*aliceState, *operator4)
		}

		// Bob decoders the info
		state3 := NewCVector()
		state3.MulMatrix(*aliceState, *CNotGate)

		state4 := NewCVector()
		operator5 := NewCMatrix()
		operator5.TensorProduct(*HadamardGate, *IdentityGate)

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
	cnot2 := ExtendControlGate(0, 1, 2, NotGate)

	if !cnot2.Equals(*CNotGate) {
		t.Error("Failed to build standard CNOT gate")
	}

	// ----.----
	// ----|----
	// ----X----
	cnot3 := ExtendControlGate(0, 2, 3, NotGate)
	cnot3sol := NewCMatrixFromElements([][]Complex{
		[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0))},
	})

	if !cnot3.Equals(*cnot3sol) {
		t.Error("Failed to build extended 3 bit CNOT gate")
	}
}

// func TestExtendControlGateSimple(t *testing.T) {
// 	cnot2 := ExtendControlGate(0, 1, 2, NotGate)
// }

func TestExtendGate(t *testing.T) {
	operator1 := NewCMatrix()
	operator1.TensorProduct(*HadamardGate, *IdentityGate)

	hadamard2 := ExtendGate(0, 2, HadamardGate)

	if !hadamard2.Equals(*operator1) {
		t.Error("Failed to construct gate")
	}
}
