package main

import (
	"fmt"
	"strings"
)

func main() {
	word := reverse("kemarin")
	fmt.Println(word)
}

func reverse(word string) string {
	var reversed string

	splitWord := strings.Split(word, "")

	for i := len(splitWord) - 1; i >= 0; i-- {
		reversed += splitWord[i]
	}

	return reversed
}
