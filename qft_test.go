package qrad

import (
	"math"
	"testing"
)

func TestQFTDraw(t *testing.T) {
	c := NewCircuit([]int{0, 1, 0, 1})
	ApplyQFT(c, 0, 3)

	c2 := NewCircuit([]int{0, 1, 0})
	ApplyQFT(c2, 0, 2)
}

func TestInverseQFTDraw(t *testing.T) {
	c := NewCircuit([]int{0, 1, 0, 1})
	ApplyInverseQFT(c, 0, 3)

	c2 := NewCircuit([]int{0, 1, 0})
	ApplyInverseQFT(c2, 0, 2)
}

func TestQFTReversable(t *testing.T) {
	soln := []int{0, 1, 0, 1}
	c := NewCircuit(soln)
	ApplyQFT(c, 0, 3)
	ApplyInverseQFT(c, 0, 3)

	c.Execute()

	for i, b := range soln {
		if c.MeasureQubit(i) != b {
			t.Error("failed to run reversable QFT")
		}
	}
}

func TestQFTInverse(t *testing.T) {
	c := NewCircuit([]int{0, 0, 0})
	c.Append(H, []int{0, 1, 2})

	c.Append(ROT(7*math.Pi/4, "7PI/4"), []int{0})
	c.Append(ROT(7*math.Pi/2, "7PI/2"), []int{1})
	c.Append(ROT(7*math.Pi, "7PI"), []int{2})

	ApplyInverseQFT(c, 0, 2)

	c.Execute()

	for i := 0; i < 3; i++ {
		if c.MeasureQubit(i) != 1 {
			t.Error("failed to inverse QFT")
		}
	}
}

func TestQFTInverseOffset(t *testing.T) {
	c := NewCircuit([]int{0, 0, 0, 0})
	c.Append(H, []int{1, 2, 3})

	c.Append(ROT(7*math.Pi/4, "7PI/4"), []int{1})
	c.Append(ROT(7*math.Pi/2, "7PI/2"), []int{2})
	c.Append(ROT(7*math.Pi, "7PI"), []int{3})

	ApplyInverseQFT(c, 1, 3)

	c.Execute()

	for i := 1; i < 4; i++ {
		if c.MeasureQubit(i) != 1 {
			t.Error("failed to inverse QFT")
		}
	}
}

// func TestQFTInverse2(t *testing.T) {
// 	c := NewCircuit([]int{0, 0, 0, 0})
// 	c.Append(H, []int{2, 3})

// 	ApplyInverseQFT(c, 0, 3)

// 	c.Execute()
// 	c.Draw()

// 	for i := 0; i < 4; i++ {
// 		fmt.Println(c.MeasureQubit(i))
// 		// if c.MeasureQubit(i) != 1 {
// 		// 	t.Error("failed to inverse QFT")
// 		// }
// 	}
// }
