package qrad

// func f4(q *Circuit) {
// 	// (a & b) & (!c & d)
// 	// (a & b) => 4
// 	q.ApplyToffoliGate(0, 1, 4)

// 	// (!c & d) => 5
// 	q.ApplyNot(2)
// 	q.ApplyToffoliGate(2, 3, 5)
// 	q.ApplyNot(2)

// 	// 4 & 5
// 	q.ApplyToffoliGate(4, 5, 6)
// }

// func f4Reverse(q *Circuit, includeFinal bool) {
// 	if includeFinal {
// 		q.ApplyToffoliGate(4, 5, 6)
// 	}
// 	q.ApplyNot(2)
// 	q.ApplyToffoliGate(2, 3, 5)
// 	q.ApplyNot(2)
// 	q.ApplyToffoliGate(0, 1, 4)
// }

// func TestOracle4Correctness(t *testing.T) {
// 	i := 0
// 	for a := 0; a < 2; a++ {
// 		for b := 0; b < 2; b++ {
// 			for c := 0; c < 2; c++ {
// 				for d := 0; d < 2; d++ {
// 					q := NewCircuit([]int{a, b, c, d, 0, 0, 0})

// 					f4(q)
// 					out := q.Measure()

// 					// fmt.Println(a, b, c, d, fmt.Sprintf("%07b", out))
// 					if a == 1 && b == 1 && c != 1 && d == 1 {
// 						if out&1 != 1 {

// 							t.Error("oracle didn't work")
// 						}
// 					} else {
// 						if out&1 != 0 {
// 							t.Error("oracle gave a yes on a bad input")
// 						}
// 					}

// 					i++
// 				}
// 			}
// 		}
// 	}
// }

// func TestOracleReverse(t *testing.T) {

// 	// oracles := []func(*Circuit){f}
// 	// reverseoracles := []func(*Circuit, bool){fReverse}

// 	for a := 0; a < 2; a++ {
// 		for b := 0; b < 2; b++ {
// 			for c := 0; c < 2; c++ {
// 				for d := 0; d < 2; d++ {
// 					q := NewCircuit([]int{a, b, c, d, 0, 0, 0})
// 					f4(q)
// 					f4Reverse(q, true)

// 					out := q.Measure()
// 					if out != (a<<6)+(b<<5)+(c<<4)+(d<<3) {
// 						t.Error("failed to reverse")
// 					}

// 					q2 := NewCircuit([]int{a, b, c, d, 0, 0, 0})
// 					f4(q2)
// 					f4Reverse(q2, false)
// 					out2 := q2.Measure()

// 					// fmt.Println(a, b, c, d, fmt.Sprintf("%07b", out2))
// 					if a == 1 && b == 1 && c != 1 && d == 1 {
// 						if out2&1 != 1 {

// 							t.Error("oracle didn't work")
// 						}
// 					} else {
// 						if out2&1 != 0 {
// 							t.Error("oracle gave a yes on a bad input")
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// }

// func Test3SatProblem(t *testing.T) {
// 	t.Skip("This will take approximately 22years to finish")

// 	totalqubits := 22
// 	q := NewCircuit(make([]int, totalqubits))

// 	// a | b | !c
// 	fmt.Println("a | b | !c")
// 	fmt.Println("q.ApplyNot(2)")
// 	q.ApplyNot(2)
// 	fmt.Println("q.Apply3OrGate(0, 1, 2, 3, 4)")
// 	q.Apply3OrGate(0, 1, 2, 3, 4)
// 	q.ApplyNot(2)

// 	// a | b | c
// 	fmt.Println("a | b | c")
// 	q.Apply3OrGate(0, 1, 2, 5, 6)

// 	// a | !b | c
// 	q.ApplyNot(1)
// 	q.Apply3OrGate(0, 1, 2, 7, 8)
// 	q.ApplyNot(1)

// 	// a | !b | !c
// 	q.ApplyGate(ExtendGateFill([]int{1, 2}, totalqubits, NotGate))
// 	q.Apply3OrGate(0, 1, 2, 9, 10)
// 	q.ApplyGate(ExtendGateFill([]int{1, 2}, totalqubits, NotGate))

// 	// !a | b | !c
// 	q.ApplyGate(ExtendGateFill([]int{0, 2}, totalqubits, NotGate))
// 	q.Apply3OrGate(0, 1, 2, 11, 12)
// 	q.ApplyGate(ExtendGateFill([]int{0, 2}, totalqubits, NotGate))

// 	// !a | b | c
// 	q.ApplyNot(0)
// 	q.Apply3OrGate(0, 1, 2, 13, 14)
// 	q.ApplyNot(0)

// 	// !a | !b | !c
// 	q.ApplyGate(ExtendGateFill([]int{0, 1, 2}, totalqubits, NotGate))
// 	q.Apply3OrGate(0, 1, 2, 15, 16)
// 	q.ApplyGate(ExtendGateFill([]int{0, 1, 2}, totalqubits, NotGate))

// 	// 4 & 6
// 	q.ApplyToffoliGate(4, 6, 17)

// 	// 8 & 10
// 	q.ApplyToffoliGate(8, 10, 18)

// 	// 12 & 14
// 	q.ApplyToffoliGate(12, 14, 19)

// 	// 17 & 18
// 	q.ApplyToffoliGate(17, 18, 20)

// 	// 19 & 16
// 	q.ApplyToffoliGate(19, 16, 21)

// 	// 20 & 21
// 	q.ApplyToffoliGate(20, 21, 22)
// 	q.Measure()
// }
