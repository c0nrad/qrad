package qrad

import (
	"testing"
)

func TestIncrement(t *testing.T) {
	c := NewCircuit([]int{0, 0, 0, 0})
	c.Append(H, []int{1, 2, 3})

	ApplyIncrement(c, 0, 3)

	c.Execute()

	if c.MeasureQubit(0) == 0 {
		t.Error("can not be even")
	}
}

func TestDecrement(t *testing.T) {
	c := NewCircuit([]int{0, 1, 0, 1})

	ApplyIncrement(c, 0, 3)
	ApplyDecrement(c, 0, 3)

	c.Execute()
	// technically measure is wrong, it's reaching the wrong endian
	if c.Measure() != 10 {
		t.Error("failed to increment and decrement")
	}
}

func TestAdd(t *testing.T) {
	c := NewCircuit([]int{0, 1, 0, 1, 1, 1})

	ApplyAdd(c, 0, 3, 4, 5)
	c.Execute()

	if c.MeasureRange(0, 3) != 10+3 {
		t.Error("failed to add", c.MeasureRange(0, 3))
	}
}
