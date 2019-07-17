package qrad

import (
	"math"
	"testing"
)

func TestModulus(t *testing.T) {
	c := Complex(complex(1, -1))
	if !NearEqual(c.Modulus(), math.Sqrt(2)) {
		t.Error("Failed to calculate modulus")
	}
}
