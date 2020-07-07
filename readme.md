# qrad Quantum Computing Simulator

qrad is a quantum simulator heavily based on [Quantum Computing for Computer Scientist](https://www.amazon.com/Quantum-Computing-Computer-Scientists-Yanofsky/dp/0521879965/), [Quantum Computing for the Determined](https://www.youtube.com/watch?v=X2q1PuI2RFI&list=PL1826E60FD05B44E4), with inspiration from [quirk](https://algassert.com/quirk), [qiskit](https://qiskit.org/), and [cirq](https://github.com/quantumlib/Cirq).

It doesn't do anything particularly unique/special. It's just a tool for me to learn how quantum algorithms work. Written in Go.

![Quantum Fourier Transform](/images/qft.gif)

### Algorithms

* [Bell](/circuits/bell): Create the bell states
* [Superdense](/circuits/superdense): Transmit two bits of classical information in one entangled qubit
* [Teleportation](/circuits/teleportation): Move the state of a qubit to another qubit (destroying the original qubit in the process)
* [Deutsch-Josza](/circuits/deutsch-josza): Determine if a blackbox function is balanced/constant
* [Quantum Fourier Transform](/circuits/qft): Take the QFT/iQFT of some qubits with a neat interface

### Example

```go
c := qrad.NewCircuit([]int{0, 0, 0, 0})
c.Append(qrad.H, []int{0, 1, 2, 3})

v := float64(13)

c.Append(qrad.ROT(v*math.Pi/8, "13PI/8"), []int{0})
c.Append(qrad.ROT(v*math.Pi/4, "13PI/4"), []int{1})
c.Append(qrad.ROT(v*math.Pi/2, "13PI/2"), []int{2})
c.Append(qrad.ROT(v*math.Pi, "13PI"), []int{3})

c.AppendBarrier()

qrad.ApplyInverseQFT(c, 0, 3)
c.Execute()
if c.Measure() != 13 {
    panic("Circuit should have measured 13")
}
```

## Blog Posts

* [Quantum Teleportation](https://blog.c0nrad.io/posts/quantum-teleportation/)
* [Quantum Fourier Transform](https://blog.c0nrad.io/posts/qft/)

### License

Copyright 2020 Stuart Larsen

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.