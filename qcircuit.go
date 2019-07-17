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
	if index >= q.Length() {
		panic("qubit out of range")
	}

	operator := ExtendGate(index, q.Qubits, HadamardGate)
	return q.ApplyGate(operator)
}

func (q *QCircuit) ApplyH(index int) *QCircuit {
	return q.ApplyHadamard(index)
}

func (q *QCircuit) ApplyNot(index int) *QCircuit {
	if index >= q.Length() {
		panic("qubit out of range")
	}

	operator := ExtendGate(index, q.Qubits, NotGate)
	return q.ApplyGate(operator)
}

func (q *QCircuit) ApplyS(index int) *QCircuit {
	if index >= q.Length() {
		panic("qubit out of range")
	}

	operator := ExtendGate(index, q.Qubits, SGate)
	return q.ApplyGate(operator)
}

func (q *QCircuit) ApplyT(index int) *QCircuit {
	if index >= q.Length() {
		panic("qubit out of range")
	}

	operator := ExtendGate(index, q.Qubits, TGate)
	return q.ApplyGate(operator)
}

func (q *QCircuit) ApplyTd(index int) *QCircuit {
	if index >= q.Length() {
		panic("qubit out of range")
	}

	operator := ExtendGate(index, q.Qubits, TdGate)
	return q.ApplyGate(operator)
}

func (q *QCircuit) ApplyCNot(control, index int) *QCircuit {
	if index >= q.Length() || control >= q.Length() {
		panic("qubit out of range")
	}

	operator := ExtendControlGate(control, index, q.Qubits, NotGate)
	return q.ApplyGate(operator)
}

func (q *QCircuit) MesaureQubit(index int) int {
	panic("not implemented")
}

func (q *QCircuit) Measure() int {
	return q.State.Measure()
}
