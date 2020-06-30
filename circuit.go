package qrad

import "fmt"

type Circuit struct {
	Moments []Moment

	State  Vector
	Qubits int
}

// func (c Circuit) Draw() {
// 	for , range c.Moments {

// 	}
// }

func NewCircuit(initialState []int) *Circuit {
	state := NewVector()
	for _, e := range initialState {
		if e != 0 && e != 1 {
			panic("initial state must be bits")
		}

		state.TensorProduct(*state, *NewQubit(e))
	}

	return &Circuit{State: *state, Qubits: len(initialState)}
}

func (q Circuit) Length() int {
	return q.State.Length()
}

func (c *Circuit) Execute() {
	for _, m := range c.Moments {
		c.State.MulMatrix(c.State, m.Matrix())
	}
}

func (c *Circuit) Append(g Gate, i int) {
	c.Moments = append(c.Moments, NewMoment(c.Qubits, g, i))
}

func (c *Circuit) AppendControl(g Gate, control, i int) {
	c.Moments = append(c.Moments, NewMomentControl(c.Qubits, g, i, []int{control}))
}

func (c Circuit) Draw() {
	DrawMoments(c.Moments)
}

func (c Circuit) PrintStates() {
	for i, s := range c.State.Elements {
		fmt.Printf("|%02b> %s\n", i, s)
	}
}

func (q *Circuit) ApplyGate(operator Gate) *Circuit {
	if operator.Matrix.Width != operator.Matrix.Height {
		panic("invalid operator dimensions")
	}

	if operator.Matrix.Width != q.Length() {
		panic("invald operator size for state vector")
	}

	q.State.MulMatrix(q.State, operator.Matrix)
	return q
}

func (q *Circuit) ApplyHadamard(index int) *Circuit {
	if index >= q.Qubits {
		panic("qubit out of range")
	}

	operator := ExtendGate(index, q.Qubits, HadamardGate)
	return q.ApplyGate(operator)
}

func (q *Circuit) ApplyH(index int) *Circuit {
	return q.ApplyHadamard(index)
}

func (q *Circuit) ApplyNot(index int) *Circuit {
	if index >= q.Qubits {
		panic("qubit out of range")
	}

	operator := ExtendGate(index, q.Qubits, NotGate)
	return q.ApplyGate(operator)
}

func (q *Circuit) ApplyS(index int) *Circuit {
	if index >= q.Qubits {
		panic("qubit out of range")
	}

	operator := ExtendGate(index, q.Qubits, SGate)
	return q.ApplyGate(operator)
}

func (q *Circuit) ApplyT(index int) *Circuit {
	if index >= q.Qubits {
		panic("qubit out of range")
	}

	operator := ExtendGate(index, q.Qubits, TGate)
	return q.ApplyGate(operator)
}

func (q *Circuit) ApplyTd(index int) *Circuit {
	if index >= q.Qubits {
		panic("qubit out of range")
	}

	operator := ExtendGate(index, q.Qubits, TdGate)
	return q.ApplyGate(operator)
}

func (q *Circuit) ApplyCNot(control, target int) *Circuit {
	if target >= q.Qubits || control >= q.Qubits {
		panic("qubit out of range")
	}

	operator := ExtendControlGate(control, target, q.Qubits, NotGate)
	return q.ApplyGate(operator)
}

func (q *Circuit) ApplyToffoliGate(control0, control1, target int) *Circuit {
	if target >= q.Qubits || control0 >= q.Qubits || control1 >= q.Qubits {
		panic("qubit out of range")
	}

	operator := ExtendControlControlGate(control0, control1, target, q.Qubits, NotGate)
	return q.ApplyGate(operator)
}

func (q *Circuit) ApplyOrGate(control0, control1, target int) *Circuit {
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

func (q *Circuit) Apply3OrGate(control0, control1, control2, interTarget, target int) *Circuit {
	q.ApplyGate(ExtendGate(interTarget, q.Qubits, NotGate))
	q.ApplyOrGate(control0, control1, interTarget)

	q.ApplyGate(ExtendGate(target, q.Qubits, NotGate))
	q.ApplyOrGate(control2, interTarget, target)

	return q
}

func (q *Circuit) MesaureQubit(index int) int {
	panic("not implemented")
}

func (q *Circuit) Measure() int {
	return q.State.Measure()
}
