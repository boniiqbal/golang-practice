package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	word := reverseInt(-190000)

	fmt.Println(word)
}

func reverseInt(number int) int {
	var result []string
	numberString := strconv.FormatInt(int64(number), 10)

	numberStringSplit := strings.Split(numberString, "")

	for i := len(numberStringSplit) - 1; i >= 0; i-- {
		if numberStringSplit[0] == "-" {	
			if i != 0 {
				result = append(result, numberStringSplit[i])
			}
		} else {
			result = append(result, numberStringSplit[i])
		}
	}

	intNumber, _ := strconv.ParseInt(strings.Join(result, ""), 10, 32)

	if number < 0 {
		return int(intNumber) * -1
	} else {
		return int(intNumber)
	}
}
