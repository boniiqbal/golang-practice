package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	word := stringSplit([]string{"hellocatapplebat", "apple,cat,bat,goodbye,hello,yellow,why"})

	fmt.Println(word)
}

type Result struct {
	Index int
	Word  string
}

func stringSplit(data []string) string {
	var result Result
	var resultArray []Result
	var newResult []string

	dictionary := strings.Split(data[1], ",")
	word := data[0]

	for _, v := range dictionary {
		if strings.Contains(word, v) == true {
			getIndex := strings.Index(word, v)
			result.Index = getIndex
			result.Word = v
	
			resultArray = append(resultArray, result)
		}
	}

	sort.Slice(resultArray, func(i, j int) bool {
		return resultArray[i].Index < resultArray[j].Index
	})

	for _, v := range resultArray {
		newResult = append(newResult, v.Word)
	}

	return strings.Join(newResult, ",")
}
