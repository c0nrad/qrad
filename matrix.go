package qrad

import (
	"fmt"
	"math/cmplx"
)

type Matrix struct {
	Elements []Complex

	Width, Height int
}

func NewMatrix() *Matrix {
	return &Matrix{Elements: make([]Complex, 0), Width: 0, Height: 0}
}

func NewMatrixFromElements(elements [][]Complex) *Matrix {
	c := &Matrix{}
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

func (c *Matrix) Resize(x, y int) *Matrix {
	c.Elements = make([]Complex, x*y)
	c.Width = x
	c.Height = y

	return c
}

func (c Matrix) At(x, y int) Complex {
	// 0, 1, 2, 3
	// 4, 5, 6, 7

	if x >= c.Width || y >= c.Height {
		fmt.Println(x, c.Width, y, c.Height)
		panic("Invalid  dimensions")

	}
	return c.Elements[x+c.Width*y]
}

func (c *Matrix) Set(x, y int, e Complex) {
	if x >= c.Width || y >= c.Height {
		panic("Invalid  dimensions")
	}
	c.Elements[x+c.Width*y] = e
}

func (c *Matrix) Add(a, b Matrix) *Matrix {
	if a.Width != b.Width || a.Height != b.Height {
		panic("Invalid dimensions")
	}

	c.Resize(a.Width, a.Height)

	for x := 0; x < a.Width; x++ {
		for y := 0; y < a.Height; y++ {
			c.Set(x, y, a.At(x, y)+b.At(x, y))
		}
	}

	return c
}

func (c *Matrix) Sub(a, b Matrix) *Matrix {
	if a.Width != b.Width || a.Height != b.Height {
		panic("Invalid dimensions")
	}

	c.Resize(a.Width, a.Height)

	for x := 0; x < a.Width; x++ {
		for y := 0; y < a.Height; y++ {
			c.Set(x, y, a.At(x, y)-b.At(x, y))
		}
	}

	return c
}

func (c *Matrix) MulScalar(a Matrix, e Complex) *Matrix {
	c.Resize(a.Width, a.Height)

	for x := 0; x < a.Width; x++ {
		for y := 0; y < a.Height; y++ {
			c.Set(x, y, e*a.At(x, y))
		}
	}

	return c
}

func (c *Matrix) Transpose(a Matrix) *Matrix {
	c.Resize(a.Height, a.Width)

	for x := 0; x < a.Width; x++ {
		for y := 0; y < a.Height; y++ {
			c.Set(y, x, a.At(x, y))
		}
	}

	return c
}

func (c *Matrix) Conjugate(a Matrix) *Matrix {
	c.Resize(a.Width, a.Height)

	for x := 0; x < a.Width; x++ {
		for y := 0; y < a.Height; y++ {
			c.Set(x, y, Complex(cmplx.Conj(complex128(a.At(x, y)))))
		}
	}

	return c
}

func (c *Matrix) Adjoint(a Matrix) *Matrix {
	c.Transpose(a)
	c.Conjugate(*c)
	return c
}

func (c *Matrix) MulMatrix(a, b Matrix) *Matrix {
	if a.Width != b.Height {
		fmt.Println(a.Width, b.Height)
		panic("Invalid dimensions")
	}
	c.Resize(b.Width, a.Height)

	for w := 0; w < b.Width; w++ {
		for h := 0; h < a.Height; h++ {
			sum := NewComplex(0, 0)
			for i := 0; i < a.Width; i++ {
				sum += a.At(i, h) * b.At(w, i)
			}
			c.Set(w, h, sum)
		}
	}
	return c
}

func (c *Matrix) Clone(a Matrix) *Matrix {
	c.Elements = a.Elements[:]
	c.Height = a.Height
	c.Width = a.Width
	return c
}

func (c *Matrix) TensorProduct(a, b Matrix) *Matrix {
	if a.Width == 0 || a.Height == 0 {
		c.Clone(b)
		return c
	} else if b.Width == 0 || b.Height == 0 {
		c.Clone(a)
		return c
	}

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

	return c
}

func (c *Matrix) TensorProducts(a ...Matrix) *Matrix {
	for _, m := range a {
		c.TensorProduct(*c, m)
	}

	return c
}

func (c Matrix) Equals(a Matrix) bool {
	for h := 0; h < a.Height; h++ {
		for w := 0; w < a.Width; w++ {
			if !c.At(w, h).Equals(a.At(w, h)) {
				return false
			}
		}
	}
	return true
}

func (a Matrix) PPrint() {
	fmt.Print("[")
	for h := 0; h < a.Height; h++ {
		fmt.Print("[")

		if h != 0 {
			fmt.Print(" ")
		}
		for w := 0; w < a.Width; w++ {
			fmt.Printf("%.01f + %.01fj, ", real(a.At(w, h)), imag(a.At(w, h)))

		}
		fmt.Print("]")
		if h != a.Height-1 {
			fmt.Println("")
		}
	}
	fmt.Print("]\n")
}

func (m Matrix) IsSquare() bool {
	return m.Height == m.Width
}

func (m Matrix) EnsureDimensions() {
	if len(m.Elements) != m.Width*m.Height {
		panic("invalid dimensions")
	}
}

func (m Matrix) IsIdentity() bool {
	m.EnsureDimensions()

	if !m.IsSquare() {
		return false
	}

	for h := 0; h < m.Height; h++ {
		for w := 0; w < m.Width; w++ {
			if w == h {
				if !m.At(w, h).Equals(NewComplex(1, 0)) {
					return false
				}
			} else {
				if !m.At(w, h).Equals(NewComplex(0, 0)) {
					return false
				}
			}
		}
	}
	return true
}

func (m Matrix) IsUnitary() bool {
	a := NewMatrix().Adjoint(m)
	a.MulMatrix(*a, m)
	return a.IsIdentity()
}
