package main

import (
	"fmt"
	"time"

	"github.com/c0nrad/qrad"
)

// 1, 0, 2 1
func encodeCharacter(r rune) []int {
	out := []int{}

	for i := 0; i < 4; i++ {
		out = append(out, int(r%4))
		r >>= 2
	}
	return out
}

func decodeCharacter(in []int) string {
	sum := byte(0)

	for i := 3; i >= 0; i-- {
		sum <<= 2
		sum += byte(in[i])
	}
	return string(sum)
}

func main() {

	messageStr := "hello world!!"
	message := []int{}
	for _, c := range messageStr {
		message = append(message, encodeCharacter(rune(c))...)
	}

	out := ""
	buff := []int{}

	for _, m := range message {
		c := qrad.NewCircuit([]int{0, 0})

		fmt.Println("Entangle two qubits")
		c.Append(qrad.H, []int{0})
		c.AppendControl(qrad.X, []int{0}, 1)
		c.Draw()

		fmt.Println("Alice encodes her message")
		switch m {
		case 0:
			break
		case 1:
			c.Append(qrad.X, []int{0})
		case 2:
			c.Append(qrad.PauliZGate, []int{0})
		case 3:
			c.Append(qrad.X, []int{0})
			c.Append(qrad.PauliZGate, []int{0})
		}
		c.Draw()

		// Bob decoders the info
		fmt.Println("Bob Decodes the message")
		c.AppendControl(qrad.X, []int{0}, 1)
		c.Append(qrad.H, []int{0})
		c.Draw()
		c.Execute()

		if qrad.ReverseEndianness(c.Measure(), 2) != m {
			panic("Failed to encode information")
		}

		buff = append(buff, m)

		if len(buff) == 4 {
			out += decodeCharacter(buff)
			fmt.Println(out)
			buff = []int{}
		}
		time.Sleep(time.Millisecond * 100)
	}

	if out != messageStr {
		panic("Failed to decode message")
	}
}
