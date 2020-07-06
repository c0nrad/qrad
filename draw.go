package qrad

import (
	"fmt"
	"strings"
)

var RightJoint = "├"
var LeftJoint = "┤"
var TopJoint = "┴"
var BottomJoint = "┬"
var Dash = "─"
var Bar = "│"
var TopLeftC = "┌"
var TopRightC = "┐"
var BottomLeftC = "└"
var BottomRightC = "┘"
var Control = "■"

func RenderMoment(m Moment) []string {
	out := []string{}
	width := len(m.Gate.Symbol) + 2

	widthExtra := 0
	if width%2 == 0 {
		widthExtra = -1
	}

	for q := 0; q < m.Size; q++ {
		if m.IsBarrier {
			for i := 0; i < m.Size; i++ {
				out = append(out, Bar+Bar)
				out = append(out, Bar+Bar)
				out = append(out, Bar+Bar)
			}
			return out
		}

		if m.IsGateAt(q) {

			if m.Gate.IsBoxed {
				if m.HasConnectionAbove(q) {
					connectedTop := strings.Repeat(Dash, width/2) + TopJoint + strings.Repeat(Dash, width/2+widthExtra)
					out = append(out, TopLeftC+connectedTop+TopRightC)
				} else {
					out = append(out, TopLeftC+strings.Repeat(Dash, width)+TopRightC)
				}

				out = append(out, LeftJoint+" "+m.Gate.Symbol+" "+RightJoint)

				if m.HasConnectionBelow(q) {
					connectedBottom := strings.Repeat(Dash, width/2) + BottomJoint + strings.Repeat(Dash, width/2+widthExtra)
					out = append(out, BottomLeftC+connectedBottom+BottomRightC)
				} else {
					out = append(out, BottomLeftC+strings.Repeat(Dash, width)+BottomRightC)
				}

			} else {
				if m.HasConnectionAbove(q) {
					connectedTop := strings.Repeat(" ", width/2) + Bar + strings.Repeat(" ", width/2+widthExtra)
					out = append(out, " "+connectedTop+" ")
				} else {
					out = append(out, strings.Repeat(" ", width+2))
				}

				out = append(out, Dash+" "+m.Gate.Symbol+" "+Dash)

				if m.HasConnectionBelow(q) {
					connectedTop := strings.Repeat(" ", width/2) + Bar + strings.Repeat(" ", width/2+widthExtra)
					out = append(out, " "+connectedTop+" ")
				} else {
					out = append(out, strings.Repeat(" ", width+2))
				}
			}

		} else {
			if m.HasConnectionAbove(q) && (m.IsControlAt(q) || m.HasConnectionBelow(q)) {
				connectedTop := strings.Repeat(" ", width/2) + Bar + strings.Repeat(" ", width/2+widthExtra)
				out = append(out, " "+connectedTop+" ")
			} else {
				out = append(out, strings.Repeat(" ", width+2))
			}

			if m.IsControlAt(q) {
				out = append(out, strings.Repeat(Dash, width/2+1)+Control+strings.Repeat(Dash, width/2+1+widthExtra))
			} else {
				out = append(out, strings.Repeat(Dash, width+2))
			}
			if m.HasConnectionBelow(q) && (m.IsControlAt(q) || m.HasConnectionAbove(q)) {
				connectedBottom := strings.Repeat(" ", width/2) + Bar + strings.Repeat(" ", width/2+widthExtra)
				out = append(out, " "+connectedBottom+" ")
			} else {
				out = append(out, strings.Repeat(" ", width+2))
			}
		}
	}
	return out
}

func JoinBuffers(a, b []string) []string {
	if len(a) != len(b) {
		fmt.Println(len(a), len(b))
		panic("mis-matching lengths")
	}

	out := []string{}
	for i := range a {
		out = append(out, a[i]+b[i])
	}
	return out
}

func RenderMoments(moments []Moment) []string {
	if len(moments) == 0 {
		return []string{}
	}

	out := make([]string, moments[0].Size*3)

	for _, m := range moments {
		a := RenderMoment(m)
		out = JoinBuffers(out, a)
	}
	return out
}

func DrawMoments(moments []Moment) {
	buf := RenderMoments(moments)
	fmt.Println(strings.Join(buf, "\n"))
}

func RenderInitialState(s []int) []string {
	out := []string{}
	for _, a := range s {
		out = append(out, "    ")
		out = append(out, fmt.Sprintf("|%d>─", a))
		out = append(out, "    ")
	}
	return out
}

func DrawCircuit(c Circuit) {
	fmt.Println(strings.Join(RenderCircuit(c), "\n"))
}

func RenderCircuit(c Circuit) []string {
	initial := RenderInitialState(c.InitialState)
	moments := RenderMoments(c.Moments)
	return JoinBuffers(initial, moments)
}
