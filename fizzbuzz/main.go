package main

import (
	"fmt"
)

func main() {
	fizzbuzz(100)
}

func fizzbuzz(word int) {
	for i := 1; i <= word; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else{
			fmt.Println(i)
		}
	}
}
