package qrad

import (
	"fmt"
	"math"
)

type Complex complex128

func NewComplex(r, i float64) Complex {
	return Complex(complex(r, i))
}

func (c Complex) Modulus() float64 {
	return math.Sqrt(real(c)*real(c) + imag(c)*imag(c))
}

var EPSILON float64 = 0.00000001

func (c Complex) Equals(a Complex) bool {
	if math.Abs(real(c)-real(a)) < EPSILON &&
		math.Abs(imag(c)-imag(a)) < EPSILON {
		return true
	}

	return false
}

func NearEqual(a, b float64) bool {
	return math.Abs(a-b) < EPSILON
}

func (c Complex) String() string {
	return fmt.Sprintf("%0.2f + %0.2fj", real(c), imag(c))
}
