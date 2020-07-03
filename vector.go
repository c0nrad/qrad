package qrad

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Vector struct {
	Elements []Complex
}

func (c Vector) At(i int) Complex {
	if i > len(c.Elements) {
		panic("Invalid offset")
	}

	return c.Elements[i]
}

func (c *Vector) Set(i int, e Complex) {
	if i >= c.Size() {
		panic("Invalid offset")
	}

	c.Elements[i] = e
}

func (c *Vector) Resize(i int) *Vector {
	c.Elements = make([]Complex, i)
	return c
}

func (c Vector) Size() int {
	return len(c.Elements)
}

func NewVector() *Vector {
	return &Vector{Elements: make([]Complex, 0)}
}

func NewQubit(i int) *Vector {
	if i == 0 {
		return &Vector{Elements: []Complex{complex(1, 0), complex(0, 0)}}
	} else {
		return &Vector{Elements: []Complex{complex(0, 0), complex(1, 0)}}
	}
}

func NewVectorFromElements(elements []Complex) *Vector {
	return &Vector{Elements: elements}
}

func (c *Vector) Add(a, b Vector) {
	if a.Size() != b.Size() {
		panic("Invalid vector lengths")
	}

	c.Resize(a.Size())

	for i := 0; i < a.Size(); i++ {
		c.Set(i, a.At(i)+b.At(i))
	}
}

func (c *Vector) Sub(a, b Vector) {
	if a.Size() != b.Size() {
		panic("Invalid vector lengths")
	}

	c.Resize(a.Size())

	for i := 0; i < a.Size(); i++ {
		c.Set(i, a.At(i)-b.At(i))
	}
}

func (c *Vector) MulScalar(scalar Complex, v Vector) {
	c.Resize(v.Size())

	for i := 0; i < v.Size(); i++ {
		c.Set(i, scalar*v.At(i))
	}
}

func (c *Vector) MulMatrix(v Vector, m Matrix) {
	if v.Size() != m.Width {
		panic("Invalid dimensions")
	}

	c.Resize(v.Size())

	for h := 0; h < c.Size(); h++ {
		sum := Complex(complex(0, 0))
		for w := 0; w < m.Width; w++ {
			sum += v.At(w) * m.At(w, h)
		}

		c.Set(h, sum)
	}
}

func (c *Vector) TensorProduct(a, b Vector) *Vector {
	if a.Size() == 0 {
		c.Elements = b.Elements[:]
		return c
	}

	if b.Size() == 0 {
		c.Elements = a.Elements[:]
		return c
	}

	c.Resize(a.Size() * b.Size())

	for ah := 0; ah < a.Size(); ah++ {
		for bh := 0; bh < b.Size(); bh++ {
			ch := ah*b.Size() + bh

			c.Set(ch, a.At(ah)*b.At(bh))
		}
	}

	return c
}

func (c Vector) Matrix() *Matrix {
	m := NewMatrix()
	m.Resize(1, c.Size())

	for i := 0; i < c.Size(); i++ {
		m.Set(0, i, c.At(i))
	}
	return m
}

func (c Vector) IsNormalized() bool {
	return NearEqual(c.Norm(), 1.0)
}

func (c *Vector) Normalize() {
	l := c.Length()
	c.MulScalar(NewComplex(1/l, 0), *c)
}

func (c *Vector) Length() float64 {
	sum := float64(0)
	for _, e := range c.Elements {
		sum += e.Modulus() * e.Modulus()
	}

	return math.Sqrt(sum)
}

func (c Vector) Norm() float64 {
	// | < p | p > |
	bra := NewMatrix().Adjoint(*c.Matrix())
	key := c.Matrix()

	bra.Elements[0].Modulus()
	innerProduct := bra.MulMatrix(*bra, *key)

	if innerProduct.Width != 1 && innerProduct.Height != 1 {
		panic("invalid inner product")
	}

	return innerProduct.At(0, 0).Modulus()
}

func (c Vector) Probabilities() map[int]float64 {
	out := make(map[int]float64)
	norm := c.Norm()
	for i, e := range c.Elements {
		out[i] = e.Modulus() * e.Modulus() / norm
	}
	return out
}

func (c Vector) Measure() int {
	c.Normalize()
	guess := rand.Float64()

	for i, e := range c.Elements {
		guess -= (e.Modulus() * e.Modulus())
		if guess < 0 {
			return i
		}
	}
	// There's like a super super super super small chance of this happening...
	panic("the numbers mason, what do they mean?")
}

func (c Vector) Equals(b Vector) bool {
	if c.Size() != b.Size() {
		return false
	}

	for i := range c.Elements {
		if !c.At(i).Equals(b.At(i)) {
			return false
		}
	}
	return true
}

func (c *Vector) MeasureQubit(index int) int {
	norm := c.Norm()
	guess := rand.Float64()
	isOne := false

	qubits := int(math.Log2(float64(len(c.Elements))))
	// so, we have to reverse the qubit order, i'm not 100% sure why

	// first determine if this qubit is collapsing to zero or one
	for i, e := range c.Elements {
		// fmt.Println(i, qubits-1-index, i&(qubits-1-index), e.Modulus()*e.Modulus()/norm)
		// fmt.Printf("|%02b> %s\n", i, e)

		if i&(1<<(qubits-1-index)) == 0 {
			continue
		} else {
			// fmt.Println("YAS")
		}

		guess -= (e.Modulus() * e.Modulus() / norm)

		if guess < 0 {
			isOne = true
			break
		}
	}

	for i := range c.Elements {

		// find all the zero elements and set them to zero
		if isOne {
			if i&(1<<(qubits-1-index)) == 0 {
				c.Elements[i] = NewComplex(0, 0)
			}
		}

		if !isOne {
			if i&(1<<(qubits-1-index)) != 0 {
				c.Elements[i] = NewComplex(0, 0)
			}
		}
	}

	c.Normalize()

	if isOne {
		return 1
	}
	return 0
}

func (c Vector) PrintProbabilities() {
	probs := c.Probabilities()
	for i := 0; i < c.Size(); i++ {
		fmt.Printf("%2d %08b %.2f\n", i, i, probs[i])
	}
}

func (c Vector) PrintChance(bits, total int) {
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
