package qrad

import "fmt"

type CMatrix struct {
	Elements []Complex

	Width, Height int
}

func NewCMatrix() *CMatrix {
	return &CMatrix{Elements: make([]Complex, 0), Width: 0, Height: 0}
}

func NewCMatrixFromElements(elements [][]Complex) *CMatrix {
	c := &CMatrix{}
	c.Height = len(elements)
	c.Width = len(elements[0])

	c.Elements = make([]Complex, c.Width*c.Height)

	for w := 0; w < c.Width; w++ {
		for h := 0; h < c.Height; h++ {
			c.Set(w, h, elements[h][w])
		}
	}

	return c
}

func (c *CMatrix) Resize(x, y int) {
	c.Elements = make([]Complex, x*y)
	c.Width = x
	c.Height = y
}

func (c CMatrix) At(x, y int) Complex {
	// 0, 1, 2, 3
	// 4, 5, 6, 7

	if x >= c.Width || y >= c.Height {
		fmt.Println(c, x, y)
		panic("Invalid  dimensions")
	}
	return c.Elements[x+c.Width*y]
}

func (c *CMatrix) Set(x, y int, e Complex) {
	if x >= c.Width || y >= c.Height {
		panic("Invalid  dimensions")
	}
	c.Elements[x+c.Width*y] = e
}

func (c *CMatrix) Add(a, b CMatrix) {
	if a.Width != b.Width || a.Height != b.Height {
		panic("Invalid dimensions")
	}

	c.Resize(a.Width, a.Height)

	for x := 0; x < a.Width; x++ {
		for y := 0; y < a.Height; y++ {
			c.Set(x, y, a.At(x, y)+b.At(x, y))
		}
	}
}

func (c *CMatrix) Sub(a, b CMatrix) {
	if a.Width != b.Width || a.Height != b.Height {
		panic("Invalid dimensions")
	}

	c.Resize(a.Width, a.Height)

	for x := 0; x < a.Width; x++ {
		for y := 0; y < a.Height; y++ {
			c.Set(x, y, a.At(x, y)-b.At(x, y))
		}
	}
}

func (c *CMatrix) MulScalar(a CMatrix, e Complex) {
	c.Resize(a.Width, a.Height)

	for x := 0; x < a.Width; x++ {
		for y := 0; y < a.Height; y++ {
			c.Set(x, y, e*a.At(x, y))
		}
	}
}

func (c *CMatrix) MulMatrix(a, b CMatrix) {
	panic("implement me!")
}

func (c *CMatrix) TensorProduct(a, b CMatrix) {
	c.Resize(a.Width*b.Width, a.Height*b.Height)

	for aw := 0; aw < a.Width; aw++ {
		for ah := 0; ah < a.Height; ah++ {
			for bw := 0; bw < b.Width; bw++ {
				for bh := 0; bh < b.Height; bh++ {
					cw := aw*b.Width + bw
					ch := ah*b.Height + bh

					c.Set(cw, ch, a.At(aw, ah)*b.At(bw, bh))
				}
			}
		}
	}
}
