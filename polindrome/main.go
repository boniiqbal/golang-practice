package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(polindrome("tpennept"))
}

func polindrome(str string) bool {
	var result bool = true
	strSplit := strings.Split(str, "")

	for i := 0; i < len(strSplit); i++ {
		return strSplit[i] == strSplit[len(strSplit)-i-1]
	}

	return result
}

// func polindrome(text string) bool {
// 	var result bool = false
// 	word := strings.Split(text, "")

// 	for i := 0; i < len(word)-1; i++ {
// 		return word[i] == word[len(word)-i-1]
// 	}

// 	return result
// }
