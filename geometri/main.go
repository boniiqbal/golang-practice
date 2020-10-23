package main

import "fmt"

func main() {
	fmt.Println(geometri([]int{2, 6, 18, 54}))
}

func geometri(data []int) (bool, bool, bool) {
	var geometriCheck bool = true
	var aritmatikaCheck bool = true
	var notCheck bool = false
	var differentGeo int
	var differentAri int

	if len(data) > 1 {
		differentGeo = data[1] / data[0]
		differentAri = data[0] - data[1]
	}

	for i := 0; i < len(data)-1; i++ {
		if differentGeo != data[i+1]/data[i] {
			geometriCheck = false
			break
		} 
	}

	for i := 0; i < len(data)- 1; i++ {
		if differentAri != data[i]-data[i+1] {
			aritmatikaCheck = false
			break
		} 
	}

	if aritmatikaCheck == false && geometriCheck == false {
		notCheck = true
	}

	return aritmatikaCheck, geometriCheck, notCheck
}
