package main

import (
	"fmt"
	"os"
	"strings"
)

func wordCount(inputString string) map[string]int {
	words := strings.Fields(inputString)
	fmt.Println(words)
	result := make(map[string]int)
	for _, word := range words {
		_, isExist := result[word]
		if isExist{
			result[word]+=1
		}else{
			result[word] = 1
		}
	}
	return result
}

func main() {
	argsWithProg := os.Args[1:]
	inputString:=strings.Join(argsWithProg," ")
	m := wordCount(inputString)
	fmt.Println(m)
}


