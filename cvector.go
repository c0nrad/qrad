package qrad

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

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

func (c *CVector) Resize(i int) *CVector {
	c.Elements = make([]Complex, i)
	return c
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

func (c *CVector) TensorProduct(a, b CVector) *CVector {
	if a.Length() == 0 {
		c.Elements = b.Elements[:]
		return c
	}

	if b.Length() == 0 {
		c.Elements = a.Elements[:]
		return c
	}

	c.Resize(a.Length() * b.Length())

	for ah := 0; ah < a.Length(); ah++ {
		for bh := 0; bh < b.Length(); bh++ {
			ch := ah*b.Length() + bh

			c.Set(ch, a.At(ah)*b.At(bh))
		}
	}

	return c
}

func (c CVector) CMatrix() *CMatrix {
	m := NewCMatrix()
	m.Resize(1, c.Length())

	for i := 0; i < c.Length(); i++ {
		m.Set(0, i, c.At(i))
	}
	return m
}

func (c CVector) Norm() float64 {
	// | < p | p > |
	bra := c.CMatrix().Dagger()
	key := c.CMatrix()

	innerProduct := bra.MulMatrix(*bra, *key)

	if innerProduct.Width != 1 && innerProduct.Height != 1 {
		panic("invalid inner product")
	}

	return innerProduct.At(0, 0).Modulus()
}

func (c CVector) Probabilities() map[int]float64 {
	out := make(map[int]float64)
	norm := c.Norm()
	for i, e := range c.Elements {
		out[i] = e.Modulus() * e.Modulus() / norm
	}
	return out
}

func (c CVector) Measure() int {
	norm := c.Norm()
	guess := rand.Float64()

	for i, e := range c.Elements {
		guess -= (e.Modulus() * e.Modulus() / norm)
		if guess < 0 {
			return i
		}
	}
	// There's like a super super super super small chance of this happening...
	panic("the numbers mason, what do they mean?")
}

func (c CVector) Equals(b CVector) bool {
	if c.Length() != b.Length() {
		return false
	}

	for i := range c.Elements {
		if !c.At(i).Equals(b.At(i)) {
			return false
		}
	}
	return true
}

func (c *CVector) MeasureQubit(index uint) int {
	norm := c.Norm()
	guess := rand.Float64()

	isOne := false

	for i, e := range c.Elements {
		if i&(1<<index) == 0 {
			// fmt.Println("CURIOYUS?")
			continue
		}

		guess -= (e.Modulus() * e.Modulus() / norm)
		if guess < 0 {
			isOne = true
		}
	}

	for i, _ := range c.Elements {
		if i&(1<<index) == 0 && isOne {
			c.Elements[i] = NewComplex(0, 0)
		} else if i&(1<<index) != 0 && !isOne {
			c.Elements[i] = NewComplex(0, 0)

		}
	}

	norm = c.Norm()
	for i, e := range c.Elements {
		if e != NewComplex(0, 0) {
			c.Elements[i] /= NewComplex(norm, 0)
		}
	}

	if isOne {
		return 1
	} else {
		return 0
	}
}

func (c CVector) PrintProbabilities() {
	probs := c.Probabilities()
	for i := 0; i < c.Length(); i++ {
		fmt.Printf("%2d %08b %.2f\n", i, i, probs[i])
	}
}

func (c CVector) PrintChance(bits, total int) {
	norm := c.Norm()

	chances := make(map[int]float64)

	for i, e := range c.Elements {
		bucket := i >> uint(total-bits)
		// fmt.Println("bucket", bucket, i, total-bits)
		chances[bucket] += (e.Modulus() * e.Modulus() / norm)
	}

	for i := 0; i < 1<<uint(bits); i++ {
		// fmt.Printf("%04b %.02f\n", i, chances[i])
	}
}
