package qrad

func ApplyIncrement(c *Circuit, start, stop int) {
	controls := []int{}
	for i := start; i < stop; i++ {
		controls = append(controls, i)
	}

	for i := stop; i > start; i-- {
		c.AppendControl(X, controls, i)
		controls = controls[0 : len(controls)-1]
	}

	c.Append(X, []int{start})
}

func ApplyDecrement(c *Circuit, start, stop int) {
	controls := []int{}
	for i := start; i <= stop; i++ {
		c.AppendControl(X, controls, i)
		controls = append(controls, i)
	}
}

func ApplyAdd(c *Circuit, startA, stopA, startB, stopB int) {
	for b := startB; b <= stopB; b++ {
		bOffset := b - startB
		r := Range(startA+bOffset, stopA)
		r = append(r, b)

		for g := stopA; g >= startA+bOffset; g-- {
			c.AppendControl(X, r[:], g)
			if len(r) != 1 {
				r = append([]int{}, r...)
				r = append(r[0:len(r)-2], r[len(r)-1:]...)
			}
		}
	}
}

func Range(start, stop int) []int {
	out := []int{}
	for i := start; i < stop; i++ {
		out = append(out, i)
	}
	return out
}
