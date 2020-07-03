package main

import (
	"math"

	"github.com/c0nrad/qrad"
)

func main() {
	c := qrad.NewCircuit([]int{0, 0})

	c.Append(qrad.H, []int{0})
	c.AppendControl(qrad.X, []int{0}, 1)

	c.Draw()
	c.Execute()
	c.PrintStates()

	VerifyBellState(c.State)
}

func VerifyBellState(v qrad.Vector) {
	solution := qrad.NewVectorFromElements([]qrad.Complex{
		qrad.Complex(complex(1/math.Sqrt(2), 0)),
		qrad.Complex(complex(0, 0)),
		qrad.Complex(complex(0, 0)),
		qrad.Complex(complex(1/math.Sqrt(2), 0)),
	})

	for i := range solution.Elements {
		if !v.At(i).Equals(solution.At(i)) {
			panic("wrong")
		}
	}

}
