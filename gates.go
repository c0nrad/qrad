package qrad

import (
	"math"
	"math/cmplx"
)

type Gate struct {
	Matrix Matrix

	Symbol string
}

func (g Gate) Operands() int {
	return int(math.Log2(float64(g.Matrix.Height)))
}

func (g Gate) IsControl() bool {
	return g.Symbol == "C" && g.Matrix.Height == 0
}

var NotGate = Gate{
	Matrix: *NewMatrixFromElements([][]Complex{
		[]Complex{Complex(complex(0, 0)), Complex(complex(1, 0))},
		[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0))},
	}),
	Symbol: "X",
}
var X = NotGate
var PauliXGate = NotGate

var HadamardGate = Gate{
	Matrix: *NewMatrixFromElements([][]Complex{
		[]Complex{Complex(complex(1/math.Sqrt(2), 0)), Complex(complex(1/math.Sqrt(2), 0))},
		[]Complex{Complex(complex(1/math.Sqrt(2), 0)), Complex(complex(-1/math.Sqrt(2), 0))},
	}),
	Symbol: "H",
}

var H = HadamardGate

var IdentityGate = Gate{
	Matrix: *NewMatrixFromElements([][]Complex{
		[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(1, 0))},
	}),
	Symbol: "I",
}
var I = IdentityGate

func ConstructNIdentity(s int) Gate {
	out := NewMatrix()
	for i := 0; i < s; i++ {
		out.TensorProduct(*out, I.Matrix)
	}
	return Gate{Matrix: *out, Symbol: "I"}
}

var PauliZGate = Gate{
	Matrix: *NewMatrixFromElements([][]Complex{
		[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(-1, 0))},
	}),
}

var SGate = Gate{
	Matrix: *NewMatrixFromElements([][]Complex{
		[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(cmplx.Exp(1i * math.Pi / 2))},
	}),
}

var SdGate = Gate{
	Matrix: *NewMatrixFromElements([][]Complex{
		[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(cmplx.Exp(-1i * math.Pi / 2))},
	}),
}

var TGate = Gate{
	Matrix: *NewMatrixFromElements([][]Complex{
		[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(cmplx.Exp(1i * math.Pi / 4))},
	}),
}

var TdGate = Gate{
	Matrix: *NewMatrixFromElements([][]Complex{
		[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(cmplx.Exp(-1i * math.Pi / 4))},
	}),
}

var CNOT = Gate{
	Matrix: *NewMatrixFromElements([][]Complex{
		[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0))},
	}),
}

var CCNotGate = Gate{
	Matrix: *NewMatrixFromElements([][]Complex{
		[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0))},
		[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0))},
	}),
}

var ToffoliGate = CCNotGate

func ExtendGate(gIndex, total int, gate Gate) Gate {
	outGate := NewMatrix()
	for i := 0; i < total; i++ {
		if i == gIndex {
			outGate.TensorProduct(*outGate, gate.Matrix)
		} else {
			outGate.TensorProduct(*outGate, IdentityGate.Matrix)
		}
	}

	return Gate{Matrix: *outGate}
}

func ExtendGateFill(indexes []int, total int, gate Gate) Gate {
	outGate := NewMatrix()
	for i := 0; i < total; i++ {

		isMatch := false
		for _, e := range indexes {
			if i == e {
				isMatch = true
			}
		}

		if isMatch {
			outGate.TensorProduct(*outGate, gate.Matrix)
		} else {
			outGate.TensorProduct(*outGate, IdentityGate.Matrix)
		}
	}

	return Gate{Matrix: *outGate}
}

func ExtendControlGate(cIndex, gIndex, total int, gate Gate) Gate {
	// fmt.Println("ExtendControlGate", cIndex, gIndex, total, gate)
	zero := NewMatrixFromElements([][]Complex{
		{Complex(complex(1, 0)), Complex(complex(0, 0))}})
	zerot := NewMatrix().Transpose(*zero)

	one := NewMatrixFromElements([][]Complex{
		{Complex(complex(0, 0)), Complex(complex(1, 0))}})
	onet := NewMatrix().Transpose(*one)

	zeroMatrix := NewMatrix().TensorProduct(*zero, *zerot)
	oneMatrix := NewMatrix().TensorProduct(*one, *onet)

	identityMatrix := ConstructNIdentity(gate.Operands())

	outControl := NewMatrix()
	outGate := NewMatrix()
	for i := 0; i < total; i++ {
		if i == cIndex {
			outControl.TensorProduct(*outControl, *zeroMatrix)
		} else {
			outControl.TensorProduct(*outControl, identityMatrix.Matrix)
		}

		if i == cIndex {
			outGate.TensorProduct(*outGate, *oneMatrix)
		} else if i == gIndex {
			outGate.TensorProduct(*outGate, gate.Matrix)
		} else {
			outGate.TensorProduct(*outGate, identityMatrix.Matrix)
		}
	}

	// fmt.Println("After ExtenControlGate")
	// outControl.PPrint()
	// outGate.PPrint()

	outControl.Add(*outControl, *outGate)
	return Gate{Matrix: *outControl, Symbol: "C" + gate.Symbol}
}

func ExtendControlControlGate(c0Index, c1Index, gIndex, total int, gate Gate) Gate {
	zero := NewMatrixFromElements([][]Complex{
		[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0))}})
	zerot := NewMatrix().Transpose(*zero)

	one := NewMatrixFromElements([][]Complex{
		[]Complex{Complex(complex(0, 0)), Complex(complex(1, 0))}})
	onet := NewMatrix().Transpose(*one)

	zeroMatrix := NewMatrix().TensorProduct(*zero, *zerot)
	oneMatrix := NewMatrix().TensorProduct(*one, *onet)

	out00Control := NewMatrix()
	out10Control := NewMatrix()
	out01Control := NewMatrix()

	outGate := NewMatrix()
	for i := 0; i < total; i++ {
		if i == c0Index {
			out00Control.TensorProduct(*out00Control, *zeroMatrix)
		} else if i == c1Index {
			out00Control.TensorProduct(*out00Control, *zeroMatrix)
		} else {
			out00Control.TensorProduct(*out00Control, IdentityGate.Matrix)
		}

		if i == c0Index {
			out10Control.TensorProduct(*out10Control, *oneMatrix)
		} else if i == c1Index {
			out10Control.TensorProduct(*out10Control, *zeroMatrix)
		} else {
			out10Control.TensorProduct(*out10Control, IdentityGate.Matrix)
		}

		if i == c0Index {
			out01Control.TensorProduct(*out01Control, *zeroMatrix)
		} else if i == c1Index {
			out01Control.TensorProduct(*out01Control, *oneMatrix)
		} else {
			out01Control.TensorProduct(*out01Control, IdentityGate.Matrix)
		}

		if i == c0Index {
			outGate.TensorProduct(*outGate, *oneMatrix)
		} else if i == c1Index {
			outGate.TensorProduct(*outGate, *oneMatrix)
		} else if i == gIndex {
			outGate.TensorProduct(*outGate, gate.Matrix)
		} else {
			outGate.TensorProduct(*outGate, IdentityGate.Matrix)
		}
	}

	out := NewMatrix()
	out.Add(*out00Control, *outGate).Add(*out, *out01Control).Add(*out, *out10Control)

	// out00Control.PPrint()
	// out01Control.PPrint()
	// out10Control.PPrint()
	// outGate.PPrint()

	return Gate{Matrix: *out}
}
