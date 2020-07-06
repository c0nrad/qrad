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
	"github.com/mum4k/termdash/widgets/sparkline"
	"github.com/mum4k/termdash/widgets/text"
)

type Display struct {
	Circuit *qrad.Circuit

	Terminal *termbox.Terminal
	Ctx      context.Context
	Cancel   context.CancelFunc

	ParentContainer *container.Container
	CircuitText     *text.Text
	AmplitudeSpark  *sparkline.SparkLine
	PhaseSpark      *sparkline.SparkLine

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

	d.AmplitudeSpark, err = sparkline.New()
	if err != nil {
		panic(err)
	}

	d.PhaseSpark, err = sparkline.New()
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
		container.SplitHorizontal(
			container.Top(
				container.Border(linestyle.Light),
				container.BorderTitle("Circuit"),
				container.PlaceWidget(d.CircuitText),
				// container.AlignHorizontal(align.HorizontalCenter),
				container.AlignVertical(align.VerticalMiddle),
			),
			container.Bottom(
				container.SplitVertical(
					container.Left(
						container.SplitHorizontal(
							container.Top(
								container.PlaceWidget(d.AmplitudeSpark),
							),
							container.Bottom(
								container.PlaceWidget(d.PhaseSpark),
							),
						),

						container.BorderTitle("Amplitude / Phase"),
						container.Border(linestyle.Light),
					),
					container.Right(
						container.PlaceWidget(d.LoggerText),
					),
				),
			),
			container.SplitPercent(75),
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
	current := qrad.RenderMoment(d.Circuit.Moments[d.Circuit.MomentExecutionIndex])
	after := qrad.RenderMoments(d.Circuit.Moments[d.Circuit.MomentExecutionIndex+1:])

	for i := 0; i < len(state); i++ {
		if d.CircuitText.Write(state[i]) != nil {
			panic("dunno")
		}

		if len(before) != 0 {
			if d.CircuitText.Write(before[i]) != nil {
				panic("dunno")
			}
		}

		if d.CircuitText.Write(current[i], text.WriteCellOpts(cell.FgColor(cell.ColorBlue))) != nil {
			panic("err")
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

func (d *Display) DrawAmplitudeSparkline() {
	out := []int{}

	capacity := d.AmplitudeSpark.ValueCapacity()
	if capacity == 0 {
		d.AmplitudeSpark.Add([]int{0})
		return
	}

	for _, e := range d.Circuit.State.Elements {
		out = append(out, int(e.Modulus()*100))
		// out = append(out, i)
	}

	for len(out) < capacity {
		out = append(out, 0)
	}

	d.AmplitudeSpark.Add(out)
}

func (d *Display) DrawPhaseSparkline() {
	out := []int{}

	capacity := d.PhaseSpark.ValueCapacity()
	if capacity == 0 {
		d.PhaseSpark.Add([]int{0})
		return
	}

	for _, e := range d.Circuit.State.Elements {
		_, angle := cmplx.Polar(complex128(e))
		height := (int((angle/(math.Pi))*100) + 100) % 100
		out = append(out, height)
		d.LoggerText.Write(fmt.Sprintf("%d\n", height))
	}

	for len(out) < capacity {
		out = append(out, 0)
	}

	d.PhaseSpark.Add(out)
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

	c.Append(qrad.ROT(13*math.Pi/8, "13PI/8"), []int{0})
	c.Append(qrad.ROT(13*math.Pi/4, "13PI/4"), []int{1})
	c.Append(qrad.ROT(13*math.Pi/2, "13PI/2"), []int{2})
	c.Append(qrad.ROT(13*math.Pi, "13PI"), []int{3})

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

				if d.Circuit.MomentExecutionIndex == len(d.Circuit.Moments)-1 {
					d.Circuit.Reset()
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
