package main

import (
	"fmt"
)

func main() {
	a := Fpb(34, 72)
	b := Kpk(34, 72)
	fmt.Println(a, b)
}

func Fpb(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func Kpk(a, b int, integers ...int) int {
	result := a * b / Fpb(a, b)

	for i := 0; i < len(integers); i++ {
		result = Kpk(result, integers[i])
	}

	return result
}
