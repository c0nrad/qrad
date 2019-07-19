package qrad

import "fmt"

type QCircuit struct {
	State  CVector
	Qubits int
}

func NewQCircuit(initialState []int) *QCircuit {
	state := NewCVector()
	for _, e := range initialState {
		if e != 0 && e != 1 {
			panic("initial state must be bits")
		}

		state.TensorProduct(*state, *NewQubit(e))
	}

	return &QCircuit{State: *state, Qubits: len(initialState)}
}

func (q QCircuit) Length() int {
	return q.State.Length()
}

func (q *QCircuit) ApplyGate(operator *CMatrix) *QCircuit {
	if operator.Width != operator.Height {
		panic("invalid operator dimensions")
	}

	if operator.Width != q.Length() {
		fmt.Println(operator.Width, q.Length())
		panic("invald operator size for state vector")
	}

	q.State.MulMatrix(q.State, *operator)
	return q
}

func (q *QCircuit) ApplyHadamard(index int) *QCircuit {
	if index >= q.Qubits {
		panic("qubit out of range")
	}

	operator := ExtendGate(index, q.Qubits, HadamardGate)
	return q.ApplyGate(operator)
}

func (q *QCircuit) ApplyH(index int) *QCircuit {
	return q.ApplyHadamard(index)
}

func (q *QCircuit) ApplyNot(index int) *QCircuit {
	if index >= q.Qubits {
		panic("qubit out of range")
	}

	operator := ExtendGate(index, q.Qubits, NotGate)
	return q.ApplyGate(operator)
}

func (q *QCircuit) ApplyS(index int) *QCircuit {
	if index >= q.Qubits {
		panic("qubit out of range")
	}

	operator := ExtendGate(index, q.Qubits, SGate)
	return q.ApplyGate(operator)
}

func (q *QCircuit) ApplyT(index int) *QCircuit {
	if index >= q.Qubits {
		panic("qubit out of range")
	}

	operator := ExtendGate(index, q.Qubits, TGate)
	return q.ApplyGate(operator)
}

func (q *QCircuit) ApplyTd(index int) *QCircuit {
	if index >= q.Qubits {
		panic("qubit out of range")
	}

	operator := ExtendGate(index, q.Qubits, TdGate)
	return q.ApplyGate(operator)
}

func (q *QCircuit) ApplyCNot(control, target int) *QCircuit {
	if target >= q.Qubits || control >= q.Qubits {
		panic("qubit out of range")
	}

	operator := ExtendControlGate(control, target, q.Qubits, NotGate)
	return q.ApplyGate(operator)
}

func (q *QCircuit) ApplyToffoliGate(control0, control1, target int) *QCircuit {
	if target >= q.Qubits || control0 >= q.Qubits || control1 >= q.Qubits {
		panic("qubit out of range")
	}

	operator := ExtendControlControlGate(control0, control1, target, q.Qubits, NotGate)
	return q.ApplyGate(operator)
}

func (q *QCircuit) ApplyOrGate(control0, control1, target int) *QCircuit {
	if target >= q.Qubits || control0 >= q.Qubits || control1 >= q.Qubits {
		panic("qubit out of range")
	}

	// X . X   ;control0
	// X . X   ;control1
	//   +     ;target
	q.ApplyGate(ExtendGateFill([]int{control0, control1}, q.Qubits, NotGate))
	q.ApplyGate(ExtendControlControlGate(control0, control1, target, q.Qubits, NotGate))
	q.ApplyGate(ExtendGateFill([]int{control0, control1}, q.Qubits, NotGate))

	return q
}

func (q *QCircuit) Apply3OrGate(control0, control1, control2, interTarget, target int) *QCircuit {
	q.ApplyGate(ExtendGate(interTarget, q.Qubits, NotGate))
	q.ApplyOrGate(control0, control1, interTarget)

	q.ApplyGate(ExtendGate(target, q.Qubits, NotGate))
	q.ApplyOrGate(control2, interTarget, target)

	return q
}

func (q *QCircuit) MesaureQubit(index int) int {
	panic("not implemented")
}

func (q *QCircuit) Measure() int {
	return q.State.Measure()
}
