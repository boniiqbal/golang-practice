package main

import "fmt"

func main() {
	fmt.Println(aritmatika([]int{2}))
}

func aritmatika(data []int) bool {
	var diff int
	var arit bool = true

	if len(data) > 1 {
		diff = data[1] - data[0]
	}

	for i := 0; i < len(data)-1; i++ {
		if diff != data[i+1]-data[i] {
			arit = false
			break
		} 
	}

	return arit
}
