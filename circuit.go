package qrad

import "fmt"

type Circuit struct {
	Moments []Moment

	InitialState []int
	State        Vector
	Qubits       int

	MomentExecutionIndex int
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

	return &Circuit{State: *state, Qubits: len(initialState), InitialState: initialState}
}

func (c *Circuit) Reset() {
	state := NewVector()
	for _, e := range c.InitialState {
		if e != 0 && e != 1 {
			panic("initial state must be bits")
		}

		state.TensorProduct(*state, *NewQubit(e))
	}

	c.State = *state
	c.MomentExecutionIndex = 0
}

func (q Circuit) Length() int {
	return q.State.Size()
}

func (c *Circuit) Execute() int {
	steps := 0
	for _, m := range c.Moments[c.MomentExecutionIndex:] {
		if m.IsBarrier {
			steps++
			continue
		}

		c.State.MulMatrix(c.State, m.Matrix())
		if !c.State.IsNormalized() {
			panic("no longer normalized")
		}
		steps++
	}

	c.MomentExecutionIndex = len(c.Moments)
	return steps
}

func (c *Circuit) Step() {
	m := c.Moments[c.MomentExecutionIndex]

	c.State.MulMatrix(c.State, m.Matrix())
	if !c.State.IsNormalized() {
		panic("no longer normalized")
	}
	c.MomentExecutionIndex++
}

func (c *Circuit) Append(g Gate, i []int) {
	c.Moments = append(c.Moments, NewMomentMultiple(c.Qubits, g, i))
}

func (c *Circuit) AppendControl(g Gate, controls []int, i int) {
	c.Moments = append(c.Moments, NewMomentControl(c.Qubits, g, i, controls))
}

func (c *Circuit) AppendBarrier() {
	c.Moments = append(c.Moments, Moment{IsBarrier: true, Size: c.Qubits})
}

func (c Circuit) Draw() {
	DrawCircuit(c)
}

func (c Circuit) PrintStates() {
	for i, s := range c.State.Elements {
		fmt.Printf("|%02b> %s\n", i, s)
	}
}

func (q *Circuit) MeasureQubit(index int) int {

	return q.State.MeasureQubit(index)
}

func (q *Circuit) Measure() int {
	if q.MomentExecutionIndex == 0 {
		panic("circuit not yet executed")
	}
	return q.State.Measure()
}

func (c *Circuit) MeasureEach() []int {
	out := []int{}
	for i := 0; i < c.Qubits; i++ {
		out = append(out, c.MeasureQubit(i))
	}
	return out
}

func (c *Circuit) MeasureRange(start, stop int) int {
	out := 0
	for i := stop; i >= start; i-- {
		out <<= 1
		out += c.MeasureQubit(i)
	}
	return out
}
