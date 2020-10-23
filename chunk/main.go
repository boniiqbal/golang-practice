package main

import "fmt"

func main() {
	arrayInt := chunk([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4)

	fmt.Println(arrayInt)
}

func chunk(array []int, number int) []interface{} {
	var newArray []int
	var newArrayAgain []interface{}
	var index int

	for index < len(array) {
		if index+number < len(array) {
			newArray = array[index : index+number]
		} else {
			newArray = array[index:]
		}

		newArrayAgain = append(newArrayAgain, newArray)
		index += number
	}

	return newArrayAgain
}
