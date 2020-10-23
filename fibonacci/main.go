package main

import "fmt"

func main() {
	// fmt.Println(fibonacciSequence(4))
	// fmt.Println(fibonacci(4))
	fmt.Println(fibonacciRange([]int{15, 1, 3}))
}

func fibonacci(n int) int {
	a := 0
	b := 1

	for i := 0; i < n; i++ {
		temp := a
		a = b
		b = temp + a
	}
	return a
}

func fibonacciRange(data []int) int {
	var total int
	for _, v := range data {
		total += v
	}

	var i int
	var fibo int
	for total > fibonacci(i) {
		fibo = fibonacci(i)
		i++
	}

	selisih := total - fibo

	return selisih
}

// func fibonacciSequence(n int) []int {
// 	var result []int

// 	for i := 0; i < n; i++ {
// 		result = append(result, fibonacci(i))
// 	}

// 	return result
// }
