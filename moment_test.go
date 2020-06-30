package qrad

import "testing"

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

	m.Matrix().PPrint()
}

func TestMomentCCNOT(t *testing.T) {
	m := NewMomentControl(3, X, 2, []int{0, 1})

	mMatrix := m.Matrix()
	if !mMatrix.Equals(CCNotGate.Matrix) {
		mMatrix.PPrint()
		CCNotGate.Matrix.PPrint()
		t.Error("did not construct CCNOT in moment")
	}
}

func TestMomentReversedBell(t *testing.T) {

}
