package qrad

import (
	"fmt"
	"math"
)

func ApplyQFT(c *Circuit, start, stop int) {
	for g := stop; g >= start; g-- {
		c.Append(H, []int{g})

		for p := start; p < g; p++ {
			d := int(math.Pow(2, float64(g-p)))
			rot := ROT(math.Pi/float64(d), fmt.Sprintf("PI/%d", d))
			c.AppendControl(rot, []int{p}, g)
		}
	}

	half := start + (stop-start)/2
	for g := start; g <= half; g++ {
		offset := g - start
		swapStart := g
		swapEnd := stop - offset
		if swapStart == swapEnd {
			continue
		}

		s := SWAP(swapEnd - swapStart - 1)
		c.Append(s, []int{swapStart})
	}
}

func ApplyInverseQFT(c *Circuit, start, stop int) {
	half := start + (stop-start)/2
	for g := half; g >= start; g-- {
		offset := g - start
		swapStart := g
		swapEnd := stop - offset
		if swapStart == swapEnd {
			continue
		}

		s := SWAP(swapEnd - swapStart - 1)
		c.Append(s, []int{swapStart})
	}

	for g := start; g <= stop; g++ {

		for p := g - 1; p >= start; p-- {
			d := int(math.Pow(2, float64(g-p)))
			rot := ROT(-math.Pi/float64(d), fmt.Sprintf("-PI/%d", d))
			c.AppendControl(rot, []int{p}, g)
		}

		c.Append(H, []int{g})

	}

}
