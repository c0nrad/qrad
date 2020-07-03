package main

import (
	"fmt"

	"github.com/c0nrad/qrad"
)

func main() {

	initState := 1
	c := qrad.NewCircuit([]int{initState, 0, 0})

	// Alice wants to transmit Qubit 0 to Qubit 2.
	// |0> Alice's transmitting qubit
	// |0> Alice's entangled qubit
	// |0> Bob's entangled qubit

	// First let's put Qubit 1 and 2 into bell state
	fmt.Println("Entangle Alice and Bob's qubit")
	c.Append(qrad.H, []int{1})
	c.AppendControl(qrad.X, []int{1}, 2)
	steps1 := c.Execute()
	if steps1 != 2 {
		panic("too many steps1")
	}

	c.Draw()
	c.PrintStates()

	fmt.Println("Alice prepares the qubit she wants to teleport")
	c.AppendControl(qrad.X, []int{0}, 1)
	c.Append(qrad.H, []int{0})
	steps2 := c.Execute()
	if steps2 != 2 {
		panic("too many steps2")
	}

	c.Draw()
	c.PrintStates()

	fmt.Println("Alice measure's her two qubits")
	q0 := c.MeasureQubit(0)
	q1 := c.MeasureQubit(1)

	fmt.Println("She gets q0=", q0, " and q1=", q1)

	c.PrintStates()

	fmt.Println("Bob conditionally applies some gates depending on the bits from alice")
	if q1 == 1 {
		// Apply X
		c.Append(qrad.X, []int{2})
	}

	if q0 == 1 {
		// Apply Z
		c.Append(qrad.PauliZGate, []int{2})
	}

	steps := c.Execute()
	if steps != q1+q0 {
		panic("not enough steps")
	}

	c.Draw()
	c.PrintStates()
	// fmt.Println(c.State.Elements)

	out := c.MeasureQubit(2)
	if out != initState {
		fmt.Println(out, initState)
		panic("failed to teleport qubit")
	}

	fmt.Println("Success! The initial state was ", initState, " and the teleported state was ", out)
}
