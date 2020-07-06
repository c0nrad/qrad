package qrad

import (
	"math"
	"testing"
)

func TestHasConnectionAbove(t *testing.T) {
	m := NewMomentControl(2, X, 1, []int{0})
	if !m.HasConnectionAbove(1) {
		t.Error("failed to determine connection above")
	}

	if !m.HasConnectionBelow(0) {
		t.Error("failed to determine connection below")
	}
}

func TestMomentMatrixIdentity(t *testing.T) {
	m := NewMomentMultiple(2, I, []int{0, 1})
	if !m.Matrix().IsIdentity() {
		t.Error("Should be an identity matrix")
	}
}

func TestMomentMatrixHadamard(t *testing.T) {
	m := NewMomentMultiple(2, H, []int{0, 1})
	if !m.Matrix().IsUnitary() {
		t.Error("matrix should be unitary")
	}

}

func TestMomentCCNOT(t *testing.T) {
	m := NewMomentControl(3, X, 2, []int{0, 1})

	var CCNotGate = Gate{
		Matrix: *NewMatrixFromElements([][]Complex{
			{Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
			{Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
			{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
			{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
			{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
			{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
			{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0))},
			{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0))},
		}),
		IsBoxed: false,
		Symbol:  "CCNOT",
	}

	mMatrix := m.Matrix()
	if !mMatrix.Equals(CCNotGate.Matrix) {
		mMatrix.PPrint()
		CCNotGate.Matrix.PPrint()
		t.Error("did not construct CCNOT in moment")
	}
}

func TestMomentSWAP(t *testing.T) {
	c := NewCircuit([]int{1, 0, 0, 0})
	c.Append(SWAP(1), []int{0})

	c.Execute()

	if c.MeasureQubit(2) != 1 {
		t.Error("failed to swap qubit")
	}

	if c.MeasureQubit(0) != 0 {
		t.Error("failed to swap qubit")
	}

	if c.MeasureQubit(3) != 0 {
		t.Error("failed to ignore qubit")
	}
}

func TestIsConnectionBelow(t *testing.T) {
	m := NewMoment(4, SWAP(1), 0)

	if !m.HasConnectionAbove(2) {
		t.Error("there should be a connection above 2")
	}

	if !m.HasConnectionBelow(0) {
		t.Error("faikled to determine connection below")
	}
	if m.HasConnectionAbove(0) {
		t.Error("not possible")
	}

	if m.HasConnectionBelow(2) {
		t.Error("no connection below 2")
	}

}

func TestCNOTGaps(t *testing.T) {
	c1 := NewCircuit([]int{0, 1, 0})
	c1.AppendControl(X, []int{1}, 2)
	c1.Execute()

	c2 := NewCircuit([]int{0, 0, 1})
	c2.AppendControl(X, []int{2}, 1)
	c2.Execute()

	if !c1.State.Equals(c2.State) {
		t.Error("CNOT should be reversable")
	}
}

func TestCROTGaps(t *testing.T) {
	c1 := NewCircuit([]int{0, 0, 0})
	c1.Append(H, []int{0, 1, 2})
	c1.AppendControl(ROT(math.Pi/2, "PI/2"), []int{1}, 2)
	c1.Execute()

	c2 := NewCircuit([]int{0, 0, 0})
	c2.Append(H, []int{0, 1, 2})
	c2.AppendControl(ROT(math.Pi/2, "PI/2"), []int{2}, 1)
	c2.Execute()

	if !c1.State.Equals(c2.State) {
		t.Error("CROT should be reversable")
	}
}
