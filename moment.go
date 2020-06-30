package qrad

import (
	"fmt"
	"strings"
)

type Moment struct {
	Gate     Gate
	Indexes  []int
	Controls []int

	Size int
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
		if t == i {
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
	if i == 0 {
		return false
	}

	if len(m.Controls) == 0 {
		return false
	}

	for _, g := range m.Controls {
		if g < i {
			return true
		}
	}

	for _, g := range m.Indexes {
		if g < i {
			return true
		}
	}

	return false
}

func (m Moment) HasConnectionBelow(i int) bool {
	if len(m.Controls) == 0 {
		return false
	}

	for _, g := range m.Controls {
		if g > i {
			return true
		}
	}

	for _, g := range m.Indexes {
		if g > i {
			return true
		}
	}

	return false
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

	for i := 0; i < moment.Size; i++ {
		if moment.IsGateAt(i) {
			gates = append(gates, moment.Gate)
		} else if moment.IsControlAt(i) {
			gates = append(gates, Gate{Matrix: *NewMatrix(), Symbol: "C"})
		} else {
			gates = append(gates, I)
		}
	}

	// fmt.Println("Gates", gates)

	for len(gates) != 1 {
		for i, g := range gates {
			if g.IsControl() {
				continue
			}

			gateIndex := i
			otherIndex := i + 1

			if i+1 >= len(gates) {
				otherIndex = i - 1
			}

			var merged Gate
			if gates[otherIndex].IsControl() {
				if otherIndex > gateIndex {
					merged = ExtendControlGate(1, 0, 2, gates[gateIndex])
				} else {
					merged = ExtendControlGate(0, 1, 2, gates[gateIndex])
				}

			} else {
				merged = Gate{Matrix: *NewMatrix().TensorProduct(gates[gateIndex].Matrix, gates[otherIndex].Matrix), Symbol: gates[gateIndex].Symbol + "×" + gates[otherIndex].Symbol}
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
