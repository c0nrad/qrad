package qrad

import (
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
