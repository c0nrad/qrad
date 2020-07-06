package qrad

import "testing"

func TestDraw(t *testing.T) {
	c := NewCircuit([]int{0, 0, 0})
	c.Append(H, []int{0, 2})

	c.Append(SWAP(1), []int{0})
	c.Append(SWAP(0), []int{0})
	c.Append(I, []int{1})

	// c.Draw()
}
