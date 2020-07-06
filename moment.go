package qrad

import (
	"fmt"
	"strings"
)

type Moment struct {
	Gate     Gate
	Indexes  []int
	Controls []int

	Size      int
	IsBarrier bool
}

func NewMoment(size int, gate Gate, index int) Moment {
	m := Moment{}
	m.Gate = gate
	m.Indexes = []int{index}
	m.Size = size
	return m
}

func NewMomentMultiple(size int, gate Gate, indexes []int) Moment {
	m := Moment{}
	m.Gate = gate
	m.Size = size
	m.Indexes = indexes

	if gate.Operands() > 1 && len(indexes) > 1 {
		panic("only one large gate per moment")
	}

	return m
}

func NewMomentControl(size int, gate Gate, index int, controls []int) Moment {
	m := Moment{}
	m.Controls = controls
	m.Indexes = []int{index}
	m.Size = size
	m.Gate = gate
	return m
}

func (m Moment) Draw() {
	b := RenderMoment(m)
	fmt.Println(strings.Join(b, "\n"))
}

func (m Moment) IsGateAt(i int) bool {
	for _, t := range m.Indexes {
		if t == i || t+m.Gate.Operands()-1 == i {
			return true
		}
	}
	return false
}

func (m Moment) IsControlAt(i int) bool {
	for _, t := range m.Controls {
		if t == i {
			return true
		}
	}
	return false
}

func (m Moment) Verify() {
	if len(m.Indexes) > 1 && len(m.Controls) > 1 {
		panic("both indexes and controls can not be greater than 1")
	}
}

func (m Moment) HasConnectionAbove(i int) bool {

	existBelow := false
	// does there exist something below i?
	for g := i; g < m.Size; g++ {
		if m.IsGateAt(g) || m.IsControlAt(g) {
			existBelow = true
		}
	}

	existAbove := false
	for g := 0; g < i; g++ {
		if m.IsGateAt(g) || m.IsControlAt(g) {
			existAbove = true
		}
	}

	return existBelow && existAbove
}

func (m Moment) HasConnectionBelow(i int) bool {

	existAbove := false
	for g := 0; g <= i; g++ {
		if m.IsGateAt(g) || m.IsControlAt(g) {
			existAbove = true
		}
	}

	existBelow := false
	// does there exist something below i?
	for g := i + 1; g < m.Size; g++ {
		if m.IsGateAt(g) || m.IsControlAt(g) {
			existBelow = true
		}
	}

	return existBelow && existAbove
}

func (m Moment) Matrix() Matrix {
	out := ConstructMomentMatrix(m)
	if !out.IsUnitary() {
		out.PPrint()
		panic("matrix should be unitary")
	}

	// out := NewMatrix()

	// for i := 0; i < m.Size; i++ {
	// 	if m.IsGateAt(i) {
	// 		out.TensorProduct(*out, m.Gate.Matrix)
	// 	} else {
	// 		out.TensorProduct(*out, I.Matrix)
	// 	}
	// }

	return out
}

func ConstructMomentMatrix(moment Moment) Matrix {

	gates := []Gate{}

	// if moment.Gate.Symbol == "SWAP" {
	// 	return ConstructSwapMomentMatrix(moment)
	// }

	for i := 0; i < moment.Size; i++ {
		if moment.IsGateAt(i) {
			gates = append(gates, moment.Gate)
			i += (moment.Gate.Operands() - 1)
		} else if moment.IsControlAt(i) {
			gates = append(gates, Gate{Matrix: *NewMatrix(), Symbol: "C"})
		} else {
			gates = append(gates, I)
		}
	}

	// fmt.Println("gates", gates)

	for len(gates) != 1 {
		hasControl := false
		for _, g := range gates {
			if g.IsControl() {
				hasControl = true
				break
			}
		}

		if !hasControl {
			out := *NewMatrix()
			for _, g := range gates {
				out.TensorProduct(out, g.Matrix)
			}
			return out
		}

		for i, g := range gates {
			if g.IsControl() || g.Matrix.IsIdentity() {
				continue
			}

			gateIndex := i
			otherIndex := i + 1
			newSymbol := ""

			if i+1 >= len(gates) {
				otherIndex = i - 1
				newSymbol = gates[otherIndex].Symbol + gates[gateIndex].Symbol
			} else {
				newSymbol = gates[gateIndex].Symbol + gates[otherIndex].Symbol

			}

			// fmt.Println(gateIndex, otherIndex, gates[gateIndex].Symbol, gates[otherIndex].Symbol)

			var merged Gate
			if gates[otherIndex].IsControl() {
				if otherIndex > gateIndex {
					merged = ExtendControlGate(1, 0, 2, gates[gateIndex])
				} else {
					merged = ExtendControlGate(0, 1, 2, gates[gateIndex])
				}

			} else {
				if gateIndex < otherIndex {
					merged = Gate{Matrix: *NewMatrix().TensorProduct(gates[gateIndex].Matrix, gates[otherIndex].Matrix), Symbol: newSymbol}
				} else {
					merged = Gate{Matrix: *NewMatrix().TensorProduct(gates[otherIndex].Matrix, gates[gateIndex].Matrix), Symbol: newSymbol}
				}
			}

			gates[gateIndex] = merged
			gates = append(gates[:otherIndex], gates[otherIndex+1:]...)
			break
		}

	}

	return gates[0].Matrix
}

func (m Moment) String() string {
	return fmt.Sprintf("Moment{G: %s, Index: %d, Control: %d", m.Gate.Symbol, m.Indexes, m.Controls)
}

func ExtendControlGate(cIndex, gIndex, total int, gate Gate) Gate {
	// fmt.Println("ExtendControlGate", cIndex, gIndex, total, gate)
	zero := NewMatrixFromElements([][]Complex{
		{Complex(complex(1, 0)), Complex(complex(0, 0))}})
	zerot := NewMatrix().Transpose(*zero)

	one := NewMatrixFromElements([][]Complex{
		{Complex(complex(0, 0)), Complex(complex(1, 0))}})
	onet := NewMatrix().Transpose(*one)

	zeroMatrix := NewMatrix().TensorProduct(*zero, *zerot)
	oneMatrix := NewMatrix().TensorProduct(*one, *onet)

	identityMatrix := ConstructNIdentity(gate.Operands())

	outControl := NewMatrix()
	outGate := NewMatrix()
	for i := 0; i < total; i++ {
		if i == cIndex {
			outControl.TensorProduct(*outControl, *zeroMatrix)
		} else {
			outControl.TensorProduct(*outControl, identityMatrix.Matrix)
		}

		if i == cIndex {
			outGate.TensorProduct(*outGate, *oneMatrix)
		} else if i == gIndex {
			outGate.TensorProduct(*outGate, gate.Matrix)
		} else {
			outGate.TensorProduct(*outGate, identityMatrix.Matrix)
		}
	}

	// fmt.Println("After ExtenControlGate")
	// outControl.PPrint()
	// outGate.PPrint()

	outControl.Add(*outControl, *outGate)
	return Gate{Matrix: *outControl, Symbol: "C" + gate.Symbol}
}

func ConstructSwapMomentMatrix(moment Moment) Matrix {
	return SWAP(2).Matrix
}
