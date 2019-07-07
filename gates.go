package qrad

import "math"

var NotGate = NewCMatrixFromElements([][]Complex{
	[]Complex{Complex(complex(0, 0)), Complex(complex(1, 0))},
	[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0))},
})
var PauliXGate = NotGate

var HadmardGate = NewCMatrixFromElements([][]Complex{
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

var CNotGate = NewCMatrixFromElements([][]Complex{
	[]Complex{Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
	[]Complex{Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0)), Complex(complex(0, 0))},
	[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0))},
	[]Complex{Complex(complex(0, 0)), Complex(complex(0, 0)), Complex(complex(1, 0)), Complex(complex(0, 0))},
})
