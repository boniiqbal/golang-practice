package main

import (
	"fmt"
	"strings"
)

func main() {
	word := maxChar("abbccc")

	fmt.Println(word)
}

func maxChar(word string) string {
	var max int
	var maxChar string

	wordSplit := strings.Split(word, "")

	char := make(map[string]int)

	for _, v := range wordSplit {
		if char[v] == 0 {
			char[v] = 1
		} else {
			char[v]++
		}
	}

	for i, v := range char {
		if v > max {
			maxChar = i
			max = v
		}
	}

	return maxChar
}
