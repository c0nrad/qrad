package main

import (
	"fmt"
	"math/rand"

	"github.com/c0nrad/qrad"
)

func main() {

	c := qrad.NewCircuit([]int{0, 0, 0, 1})
	c.Append(qrad.H, []int{0, 1, 2, 3})

	if rand.Int()%2 == 1 {
		fmt.Println("Applying a constant oracle")
		c.Moments = append(c.Moments, ConstantOracle()...)
	} else {
		fmt.Println("Applying a balanced oracle")
		c.Moments = append(c.Moments, BalancedOracle()...)
	}

	c.Append(qrad.H, []int{0, 1, 2})
	c.Execute()
	c.Draw()

	c0 := c.MeasureQubit(0)
	c1 := c.MeasureQubit(1)
	c2 := c.MeasureQubit(2)
	fmt.Println("We received ", c0, c1, c2)

	if c0+c1+c2 == 0 {
		fmt.Println("Function is constant!")
	} else {
		fmt.Println("Function is balanced!")
	}
}

func BalancedOracle() []qrad.Moment {
	out := []qrad.Moment{}

	out = append(out, qrad.NewMomentControl(4, qrad.X, 3, []int{0}))
	out = append(out, qrad.NewMomentControl(4, qrad.X, 3, []int{1}))
	out = append(out, qrad.NewMomentControl(4, qrad.X, 3, []int{2}))
	return out
}

func ConstantOracle() []qrad.Moment {
	out := []qrad.Moment{}
	out = append(out, qrad.NewMoment(4, qrad.X, 4))
	return out
}
