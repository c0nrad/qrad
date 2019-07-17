package qrad

import "math"
import "math/cmplx"

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
