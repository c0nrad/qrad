package qrad

import (
	"math"
	"math/cmplx"
)

var NotGate = NewCMatrixFromElements([][]Complex{
	[]Complex{Complex(complex(0, 0)), Complex(complex(1, 0))},
	[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0))},
})
var PauliXGate = NotGate

var HadamardGate = NewCMatrixFromElements([][]Complex{
	[]Complex{Complex(complex(1/math.Sqrt(2), 0)), Complex(complex(1/math.Sqrt(2), 0))},
	[]Complex{Complex(complex(1/math.Sqrt(2), 0)), Complex(complex(-1/math.Sqrt(2), 0))},
})

var IdentityGate = NewCMatrixFromElements([][]Complex{
	[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0))},
	[]Complex{Complex(complex(0, 0)), Complex(complex(1, 0))},
})

var PauliZGate = NewCMatrixFromElements([][]Complex{
	[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0))},
	[]Complex{Complex(complex(0, 0)), Complex(complex(-1, 0))},
})

var SGate = NewCMatrixFromElements([][]Complex{
	[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0))},
	[]Complex{Complex(complex(0, 0)), Complex(cmplx.Exp(1i * math.Pi / 2))},
})

var SdGate = NewCMatrixFromElements([][]Complex{
	[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0))},
	[]Complex{Complex(complex(0, 0)), Complex(cmplx.Exp(-1i * math.Pi / 2))},
})

var TGate = NewCMatrixFromElements([][]Complex{
	[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0))},
	[]Complex{Complex(complex(0, 0)), Complex(cmplx.Exp(1i * math.Pi / 4))},
})

var TdGate = NewCMatrixFromElements([][]Complex{
	[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0))},
	[]Complex{Complex(complex(0, 0)), Complex(cmplx.Exp(-1i * math.Pi / 4))},
})

var CNotGate = NewCMatrixFromElements([][]Complex{
	[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
	[]Complex{Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
	[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0))},
	[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0))},
})

var CCNotGate = NewCMatrixFromElements([][]Complex{
	[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
	[]Complex{Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
	[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
	[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
	[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
	[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
	[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0))},
	[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0))},
})

var ToffoliGate = CCNotGate

func ExtendGate(gIndex, total int, gate *CMatrix) *CMatrix {

	outGate := NewCMatrix()
	for i := 0; i < total; i++ {
		if i == gIndex {
			outGate.TensorProduct(*outGate, *gate)
		} else {
			outGate.TensorProduct(*outGate, *IdentityGate)
		}
	}

	return outGate
}

func ExtendGateFill(indexes []int, total int, gate *CMatrix) *CMatrix {
	outGate := NewCMatrix()
	for i := 0; i < total; i++ {

		isMatch := false
		for _, e := range indexes {
			if i == e {
				isMatch = true
			}
		}

		if isMatch {
			outGate.TensorProduct(*outGate, *gate)
		} else {
			outGate.TensorProduct(*outGate, *IdentityGate)
		}
	}

	return outGate
}

func ExtendControlGate(cIndex, gIndex, total int, gate *CMatrix) *CMatrix {
	zero := NewCMatrixFromElements([][]Complex{
		[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0))}})
	zerot := NewCMatrix().Transpose(*zero)

	one := NewCMatrixFromElements([][]Complex{
		[]Complex{Complex(complex(0, 0)), Complex(complex(1, 0))}})
	onet := NewCMatrix().Transpose(*one)

	zeroMatrix := NewCMatrix().TensorProduct(*zero, *zerot)
	oneMatrix := NewCMatrix().TensorProduct(*one, *onet)

	outControl := NewCMatrix()
	outGate := NewCMatrix()
	for i := 0; i < total; i++ {
		if i == cIndex {
			outControl.TensorProduct(*outControl, *zeroMatrix)
		} else {
			outControl.TensorProduct(*outControl, *IdentityGate)
		}

		if i == cIndex {
			outGate.TensorProduct(*outGate, *oneMatrix)
		} else if i == gIndex {
			outGate.TensorProduct(*outGate, *gate)
		} else {
			outGate.TensorProduct(*outGate, *IdentityGate)
		}
	}

	outControl.Add(*outControl, *outGate)
	return outControl
}

func ExtendControlControlGate(c0Index, c1Index, gIndex, total int, gate *CMatrix) *CMatrix {
	zero := NewCMatrixFromElements([][]Complex{
		[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0))}})
	zerot := NewCMatrix().Transpose(*zero)

	one := NewCMatrixFromElements([][]Complex{
		[]Complex{Complex(complex(0, 0)), Complex(complex(1, 0))}})
	onet := NewCMatrix().Transpose(*one)

	zeroMatrix := NewCMatrix().TensorProduct(*zero, *zerot)
	oneMatrix := NewCMatrix().TensorProduct(*one, *onet)

	out00Control := NewCMatrix()
	out10Control := NewCMatrix()
	out01Control := NewCMatrix()

	outGate := NewCMatrix()
	for i := 0; i < total; i++ {
		if i == c0Index {
			out00Control.TensorProduct(*out00Control, *zeroMatrix)
		} else if i == c1Index {
			out00Control.TensorProduct(*out00Control, *zeroMatrix)
		} else {
			out00Control.TensorProduct(*out00Control, *IdentityGate)
		}

		if i == c0Index {
			out10Control.TensorProduct(*out10Control, *oneMatrix)
		} else if i == c1Index {
			out10Control.TensorProduct(*out10Control, *zeroMatrix)
		} else {
			out10Control.TensorProduct(*out10Control, *IdentityGate)
		}

		if i == c0Index {
			out01Control.TensorProduct(*out01Control, *zeroMatrix)
		} else if i == c1Index {
			out01Control.TensorProduct(*out01Control, *oneMatrix)
		} else {
			out01Control.TensorProduct(*out01Control, *IdentityGate)
		}

		if i == c0Index {
			outGate.TensorProduct(*outGate, *oneMatrix)
		} else if i == c1Index {
			outGate.TensorProduct(*outGate, *oneMatrix)
		} else if i == gIndex {
			outGate.TensorProduct(*outGate, *gate)
		} else {
			outGate.TensorProduct(*outGate, *IdentityGate)
		}
	}

	out := NewCMatrix()
	out.Add(*out00Control, *outGate).Add(*out, *out01Control).Add(*out, *out10Control)

	// out00Control.PPrint()
	// out01Control.PPrint()
	// out10Control.PPrint()
	// outGate.PPrint()

	return out
}
