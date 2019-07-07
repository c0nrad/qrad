package qrad

import "math/rand"

type CVector struct {
	Elements []Complex
}

func (c CVector) At(i int) Complex {
	if i > len(c.Elements) {
		panic("Invalid offset")
	}

	return c.Elements[i]
}

func (c *CVector) Set(i int, e Complex) {
	if i >= c.Length() {
		panic("Invalid offset")
	}

	c.Elements[i] = e
}

func (c *CVector) Resize(i int) {
	c.Elements = make([]Complex, i)
}

func (c CVector) Length() int {
	return len(c.Elements)
}

func NewCVector() *CVector {
	return &CVector{Elements: make([]Complex, 0)}
}

func NewQubit(i int) *CVector {
	if i == 0 {
		return &CVector{Elements: []Complex{complex(1, 0), complex(0, 0)}}
	} else {
		return &CVector{Elements: []Complex{complex(0, 0), complex(1, 0)}}
	}
}

func NewCVectorFromElements(elements []Complex) *CVector {
	return &CVector{Elements: elements}
}

func (c *CVector) Add(a, b CVector) {
	if a.Length() != b.Length() {
		panic("Invalid vector lengths")
	}

	c.Resize(a.Length())

	for i := 0; i < a.Length(); i++ {
		c.Set(i, a.At(i)+b.At(i))
	}
}

func (c *CVector) Sub(a, b CVector) {
	if a.Length() != b.Length() {
		panic("Invalid vector lengths")
	}

	c.Resize(a.Length())

	for i := 0; i < a.Length(); i++ {
		c.Set(i, a.At(i)-b.At(i))
	}
}

func (c *CVector) MulScalar(scalar Complex, v CVector) {
	c.Resize(v.Length())

	for i := 0; i < v.Length(); i++ {
		c.Set(i, scalar*v.At(i))
	}
}

func (c *CVector) MulMatrix(v CVector, m CMatrix) {
	if v.Length() != m.Width {
		panic("Invalid dimensions")
	}

	c.Resize(v.Length())

	for h := 0; h < c.Length(); h++ {
		sum := Complex(complex(0, 0))
		for w := 0; w < m.Width; w++ {
			sum += v.At(w) * m.At(w, h)
		}

		c.Set(h, sum)
	}
}

func (c *CVector) TensorProduct(a, b CVector) {
	c.Resize(a.Length() * b.Length())

	for ah := 0; ah < a.Length(); ah++ {
		for bh := 0; bh < b.Length(); bh++ {
			ch := ah*b.Length() + bh

			c.Set(ch, a.At(ah)*b.At(bh))
		}
	}
}

func (c *CVector) Measure() int {
	for i, e := range c.Elements {
		if e.Modulus() > rand.Float64() {
			return i
		}
	}
	panic("it should of worked")
}
