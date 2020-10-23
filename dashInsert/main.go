package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(dashInsert("847345920559889"))
}

func dashInsert(number string) string {
	var result []string
	numberSplit := strings.Split(number, "")

	for i := 0; i < len(numberSplit); i++ {
		var numberIntNext int64
		result = append(result, numberSplit[i])

		numberInt, _ := strconv.ParseInt(numberSplit[i], 10, 32)
		if i+1 < len(numberSplit) {
			numberIntNext, _ = strconv.ParseInt(numberSplit[i+1], 10, 32)
		}
		odd, even := evenOdd(int(numberInt))
		oddNew, evenNew := evenOdd(int(numberIntNext))

		if odd == true && oddNew == true {
			result = append(result, "-")
		}

		if even == true && evenNew == true {
			result = append(result, "+")
		}
	}

	newResult := strings.Join(result, "")

	return newResult
}

func evenOdd(data int) (bool, bool) {
	var odd bool = false
	var even bool = false

	if data%2 == 0 {
		even = true
	} else {
		odd = true
	}

	return odd, even
}
