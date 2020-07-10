package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/c0nrad/qrad"
)

func main() {
	a := 2
	N := 15

	c := qrad.NewCircuit([]int{1, 0, 0, 0, 0, 0, 0, 0})
	c.Append(qrad.H, []int{4, 5, 6, 7})
	c.AppendControl(qrad.SWAP(0), []int{4}, 2)
	c.AppendControl(qrad.SWAP(0), []int{4}, 1)
	c.AppendControl(qrad.SWAP(0), []int{4}, 0)

	c.AppendControl(qrad.SWAP(1), []int{5}, 1)
	c.AppendControl(qrad.SWAP(1), []int{5}, 0)

	qrad.ApplyQFT(c, 4, 7)

	c.Draw()
	c.Execute()
	r := c.MeasureRange(4, 7)
	fmt.Println("Measured a period of ")

	if aXmodn(a, r/2, N) == (-1+N)%N {
		fmt.Println("Bummer they were congruent")
	} else {
		fmt.Println("cool, they are not congruent", aXmodn(r, a, N), N-1%N)
	}

	fmt.Println("We found the following factors")
	fmt.Println(FindFactor(r, a, N))
}

func ClassicalShor() {
	N := 35

	fmt.Println("We are trying to factor:", N)
	a := rand.Intn(N)
	divisor := GCD(N, a)
	for divisor != 1 {
		fmt.Println("Too easy, let's find something with no common factors")
		a = rand.Intn(N)
		divisor = GCD(N, a)
	}

	fmt.Println("We have choosen our guess to be", a)
	ShowPeriod(a, N)

	r := ModPeriod(a, N)
	fmt.Println("The period is ", r)

	if aXmodn(a, r/2, N) == (-1+N)%N {
		fmt.Println("BUMMER")
	} else {
		fmt.Println("cool, they are not congruent", aXmodn(r, a, N), N-1%N)
	}

	fmt.Println(FindFactor(r, a, N))
}

func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func FindFactor(r, a, N int) []int {
	f1 := GCD(Pow(a, r/2)+1, N)
	f2 := GCD(Pow(a, r/2)-1, N)

	return []int{f1, f2}
}

func ShowPeriod(a, N int) {
	for x := 0; x < 10; x++ {
		fmt.Printf("%d**%d == %d mod %d\n", a, x, int(math.Pow(float64(a), float64(x)))%N, N)
	}
}

func ModPeriod(a, N int) int {
	t := aXmodn(a, 1, N)
	for x := 2; x < N; x++ {
		if t == aXmodn(a, x, N) {
			return x - 1
		}
	}
	panic("no period")
}

func aXmodn(a, x, N int) int {
	return int(math.Pow(float64(a), float64(x))) % N
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
