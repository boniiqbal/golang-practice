package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(wordSplit([]string{"hellocatapplegoodbye", "apple,bat,cat,goodbye,hello,yellow,why"}))
}

type Result struct {
	Index int
	Word  string
}

func wordSplit(data []string) string {
	var result Result
	var resultArray []Result
	var newResult []string

	word := data[0]
	dictionary := data[1]

	dictionarySplit := strings.Split(dictionary, ",")

	for _, v := range dictionarySplit {
		if strings.Contains(word, v) == true {
			indexWord := strings.Index(word, v)
			result.Index = indexWord
			result.Word = v

			resultArray = append(resultArray, result)
		}
	}

	sort.Slice(resultArray, func(i int, j int) bool {
		return resultArray[i].Index < resultArray[j].Index
	})

	for _, v := range resultArray {
		newResult = append(newResult, v.Word)
	}

	return strings.Join(newResult, ",")
}
