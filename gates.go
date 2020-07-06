package qrad

import (
	"math"
	"math/cmplx"
)

type Gate struct {
	Matrix Matrix

	Name string

	// Draw
	IsBoxed bool
	Symbol  string
}

func (g Gate) Operands() int {
	return int(math.Log2(float64(g.Matrix.Height)))
}

func (g Gate) IsControl() bool {
	return g.Symbol == "C" && g.Matrix.Height == 0
}

var NotGate = Gate{
	Matrix: *NewMatrixFromElements([][]Complex{
		{Complex(complex(0, 0)), Complex(complex(1, 0))},
		{Complex(complex(1, 0)), Complex(complex(0, 0))},
	}),
	Name:    "Pauli X",
	Symbol:  "X",
	IsBoxed: true,
}
var X = NotGate
var PauliXGate = NotGate

var HadamardGate = Gate{
	Matrix: *NewMatrixFromElements([][]Complex{
		{Complex(complex(1/math.Sqrt(2), 0)), Complex(complex(1/math.Sqrt(2), 0))},
		{Complex(complex(1/math.Sqrt(2), 0)), Complex(complex(-1/math.Sqrt(2), 0))},
	}),
	Name:    "Hadamard",
	Symbol:  "H",
	IsBoxed: true,
}

var H = HadamardGate

var IdentityGate = Gate{
	Matrix: *NewMatrixFromElements([][]Complex{
		{Complex(complex(1, 0)), Complex(complex(0, 0))},
		{Complex(complex(0, 0)), Complex(complex(1, 0))},
	}),
	Name:    "Identity",
	Symbol:  "I",
	IsBoxed: true,
}
var I = IdentityGate

func ConstructNIdentity(s int) Gate {
	out := NewMatrix()
	for i := 0; i < s; i++ {
		out.TensorProduct(*out, I.Matrix)
	}
	return Gate{Matrix: *out, Symbol: "I", Name: "Identity", IsBoxed: true}
}

var PauliZGate = Gate{
	Matrix: *NewMatrixFromElements([][]Complex{
		{Complex(complex(1, 0)), Complex(complex(0, 0))},
		{Complex(complex(0, 0)), Complex(complex(-1, 0))},
	}),
	Name:    "Pauli Z",
	Symbol:  "Z",
	IsBoxed: true,
}

var SGate = Gate{
	Matrix: *NewMatrixFromElements([][]Complex{
		{Complex(complex(1, 0)), Complex(complex(0, 0))},
		{Complex(complex(0, 0)), Complex(cmplx.Exp(1i * math.Pi / 2))},
	}),
	IsBoxed: true,
	Name:    "S",
	Symbol:  "S",
}

var SdGate = Gate{
	Matrix: *NewMatrixFromElements([][]Complex{
		{Complex(complex(1, 0)), Complex(complex(0, 0))},
		{Complex(complex(0, 0)), Complex(cmplx.Exp(-1i * math.Pi / 2))},
	}),
	Name:    "Sd",
	Symbol:  "Sd",
	IsBoxed: true,
}

var TGate = Gate{
	Matrix: *NewMatrixFromElements([][]Complex{
		{Complex(complex(1, 0)), Complex(complex(0, 0))},
		{Complex(complex(0, 0)), Complex(cmplx.Exp(1i * math.Pi / 4))},
	}),
	Name:    "T",
	Symbol:  "T",
	IsBoxed: true,
}

var TdGate = Gate{
	Matrix: *NewMatrixFromElements([][]Complex{
		{Complex(complex(1, 0)), Complex(complex(0, 0))},
		{Complex(complex(0, 0)), Complex(cmplx.Exp(-1i * math.Pi / 4))},
	}),
	Name:    "T",
	Symbol:  "Td",
	IsBoxed: true,
}

var CNOT = Gate{
	Matrix: *NewMatrixFromElements([][]Complex{
		{Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
		{Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
		{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0))},
		{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0))},
	}),
	Name:    "Controlled Not",
	Symbol:  "CNOT",
	IsBoxed: false,
}

func ROT(angle float64, symbol string) Gate {
	return Gate{
		Matrix: *NewMatrixFromElements([][]Complex{
			{Complex(complex(1, 0)), Complex(complex(0, 0))},
			{Complex(complex(0, 0)), Complex(cmplx.Exp(1i * complex(angle, 0)))},
		}),
		Name:    "ROT",
		Symbol:  symbol,
		IsBoxed: false,
	}
}

func SWAP(gap int) Gate {
	bits := gap + 2
	m := ConstructNIdentity(bits).Matrix
	length := int(math.Pow(2, float64(bits)))
	offset := length/2 - 1

	for i := 1; i < length/2; i += 2 {
		m.Set(i, i, NewComplex(0, 0))
		m.Set(i+offset, i, NewComplex(1, 0))
		m.Set(i, i+offset, NewComplex(1, 0))
	}

	for i := length / 2; i < length; i += 2 {
		m.Set(i, i, NewComplex(0, 0))

	}

	return Gate{
		Matrix:  m,
		Name:    "SWAP",
		Symbol:  "X",
		IsBoxed: false,
	}
}
