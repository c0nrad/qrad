package main

import (
	"context"
	"fmt"
	"math"
	"math/cmplx"
	"time"

	"github.com/c0nrad/qrad"
	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/termbox"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/barchart"
	"github.com/mum4k/termdash/widgets/text"
)

type Display struct {
	Circuit *qrad.Circuit

	Terminal *termbox.Terminal
	Ctx      context.Context
	Cancel   context.CancelFunc

	ParentContainer *container.Container
	CircuitText     *text.Text
	AmplitudeChart  *barchart.BarChart
	PhaseChart      *barchart.BarChart

	LoggerText *text.Text
}

func (d *Display) InitTerminal() {
	var err error
	d.Terminal, err = termbox.New()
	if err != nil {
		panic(err)
	}

	d.Ctx, d.Cancel = context.WithCancel(context.Background())

	d.CircuitText, err = text.New()
	if err != nil {
		panic(err)
	}

	d.AmplitudeChart, err = barchart.New(barchart.Labels(d.GenerateLabels()), barchart.LabelColors(d.GenerateLabelColors()))
	if err != nil {
		panic(err)
	}

	d.PhaseChart, err = barchart.New(barchart.Labels(d.GenerateLabels()), barchart.LabelColors(d.GenerateLabelColors()))
	if err != nil {
		panic(err)
	}

	d.LoggerText, err = text.New()
	if err != nil {
		panic(err)
	}

	d.ParentContainer, err = container.New(
		d.Terminal,
		container.Border(linestyle.Light),
		container.BorderTitle("Inverse Quantum Fourier Transform"),
		container.FocusedColor(cell.ColorCyan),
		// container.BorderColor(cell.ColorNumber(8)),
		container.SplitHorizontal(
			container.Top(
				container.Border(linestyle.Light),
				container.BorderTitle("Circuit"),
				container.PlaceWidget(d.CircuitText),
				// container.BorderColor(cell.ColorYellow),
				// container.AlignHorizontal(align.HorizontalCenter),
				container.AlignVertical(align.VerticalMiddle),
			),
			container.Bottom(
				container.SplitVertical(
					container.Left(
						container.SplitHorizontal(
							container.Top(
								container.PlaceWidget(d.AmplitudeChart),
							),
							container.Bottom(
								container.PlaceWidget(d.PhaseChart),
							),
						),

						container.BorderTitle("Amplitude / Phase"),
						container.Border(linestyle.Light),
					),
					container.Right(
						container.PlaceWidget(d.LoggerText),
						container.BorderTitle("Logger"),
						container.Border(linestyle.Light),
					),
				),
			),
			container.SplitPercent(60),
		),
	)
}

func (d *Display) Run() {
	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			d.Cancel()
		}
	}

	if err := termdash.Run(d.Ctx, d.Terminal, d.ParentContainer, termdash.KeyboardSubscriber(quitter)); err != nil {
		panic(err)
	}
}

func (d *Display) DrawCircuit() {
	d.CircuitText.Reset()

	state := qrad.RenderInitialState(d.Circuit.InitialState)
	before := qrad.RenderMoments(d.Circuit.Moments[0:d.Circuit.MomentExecutionIndex])

	current := []string{}
	after := []string{}
	if len(d.Circuit.Moments) != d.Circuit.MomentExecutionIndex {
		current = qrad.RenderMoment(d.Circuit.Moments[d.Circuit.MomentExecutionIndex])
		after = qrad.RenderMoments(d.Circuit.Moments[d.Circuit.MomentExecutionIndex+1:])
	}

	for i := 0; i < len(state); i++ {
		if d.CircuitText.Write(state[i]) != nil {
			panic("dunno")
		}

		if len(before) != 0 {
			if d.CircuitText.Write(before[i], text.WriteCellOpts(cell.FgColor(cell.ColorNumber(8)))) != nil {
				panic("dunno")
			}
		}

		if len(current) != 0 {
			if d.CircuitText.Write(current[i], text.WriteCellOpts(cell.FgColor(cell.ColorBlue))) != nil {
				panic("err")
			}
		}

		if len(after) != 0 {
			if d.CircuitText.Write(after[i]) != nil {
				panic("dunno")
			}
		}

		if d.CircuitText.Write("\n") != nil {
			panic("dunno")
		}
	}
}

func (d *Display) GenerateLabels() []string {
	out := []string{}
	for i := range d.Circuit.State.Elements {
		out = append(out, fmt.Sprintf("%d", i))
	}
	return out
}

func (d *Display) GenerateLabelColors() []cell.Color {
	out := []cell.Color{}
	for _ = range d.Circuit.State.Elements {
		out = append(out, cell.ColorGreen)
	}
	return out
}

func (d *Display) DrawAmplitudeSparkline() {
	out := []int{}

	capacity := d.AmplitudeChart.ValueCapacity()
	if capacity == 0 {
		d.AmplitudeChart.Values([]int{0}, 100)
		return
	}

	for i := range d.Circuit.State.Elements {
		reversed := qrad.ReverseEndianness(i, d.Circuit.Qubits)
		e := d.Circuit.State.Elements[reversed]
		out = append(out, int(e.Modulus()*100))

	}

	// for len(out) < capacity {
	// 	out = append(out, 0)
	// }

	d.AmplitudeChart.Values(out, 100)
}

func (d *Display) DrawPhaseSparkline() {
	out := []int{}

	capacity := d.PhaseChart.ValueCapacity()
	if capacity == 0 {
		d.PhaseChart.Values([]int{0}, 100)
		return
	}

	for i := range d.Circuit.State.Elements {
		reversed := qrad.ReverseEndianness(i, d.Circuit.Qubits)
		e := d.Circuit.State.Elements[reversed]
		_, angle := cmplx.Polar(complex128(e))
		height := (int((angle/(math.Pi))*100) + 100) % 100
		out = append(out, height)
	}

	// for len(out) < capacity {
	// 	out = append(out, 0)
	// }

	d.PhaseChart.Values(out, 100)
}

func (d *Display) UpdateLogger() {
	if d.Circuit.MomentExecutionIndex == 0 {
		d.LoggerText.Write(" Putting 4 qubits into equal superpositions\n")
	}

	if d.Circuit.MomentExecutionIndex == 1 {
		d.LoggerText.Write(" Applying rotations to qubits at phase 13\n")
	}

	if d.Circuit.MomentExecutionIndex == 6 {
		d.LoggerText.Write(" Performing Quantum Fourier Transform Inverse\n")
	}

	if d.Circuit.MomentExecutionIndex == 18 {
		d.LoggerText.Write(" Measured 13, fourier transform worked!\n")
	}

}

func BuildCircuit() *qrad.Circuit {
	// // c := qrad.NewCircuit([]int{0, 1, 0, 1, 1, 1})
	// c := qrad.NewCircuit([]int{0, 0, 0, 0, 0, 0})

	// c.Append(qrad.H, []int{1, 3})
	// qrad.ApplyAdd(c, 0, 3, 4, 5)
	// c.AppendBarrier()
	// qrad.ApplyIncrement(c, 0, 3)
	// c.AppendBarrier()
	// qrad.ApplyDecrement(c, 4, 5)

	// soln := []int{0, 1, 0, 1}
	// c := qrad.NewCircuit(soln)
	// qrad.ApplyQFT(c, 0, 3)
	// qrad.ApplyInverseQFT(c, 0, 3)

	c := qrad.NewCircuit([]int{0, 0, 0, 0})
	c.Append(qrad.H, []int{0, 1, 2, 3})

	v := float64(13)

	c.Append(qrad.ROT(v*math.Pi/8, "13PI/8"), []int{0})
	c.Append(qrad.ROT(v*math.Pi/4, "13PI/4"), []int{1})
	c.Append(qrad.ROT(v*math.Pi/2, "13PI/2"), []int{2})
	c.Append(qrad.ROT(v*math.Pi, "13PI"), []int{3})

	c.AppendBarrier()

	qrad.ApplyInverseQFT(c, 0, 3)

	return c
}

func main() {
	d := Display{}
	d.Circuit = BuildCircuit()
	d.InitTerminal()
	d.DrawAmplitudeSparkline()
	d.DrawCircuit()
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				d.DrawAmplitudeSparkline()
				d.DrawPhaseSparkline()
				d.DrawCircuit()
				d.UpdateLogger()

				if d.Circuit.MomentExecutionIndex == len(d.Circuit.Moments) {
					ticker.Stop()

				} else {
					d.Circuit.Step()
				}

			case <-d.Ctx.Done():
				return
			}
		}
	}()
	d.Run()
	d.Terminal.Close()
}
