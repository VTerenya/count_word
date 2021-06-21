package main

import (
	"fmt"
	"strings"
)

func wordCount(s string) map[string]int {
	ss := strings.Split(s, " ")
	fmt.Println(ss)
	visited := make(map[string]bool)
	result := make(map[string]int)
	for _, item := range ss {
		if visited[item] == true || item==""{
			continue
		} else {
			visited[item] = true
			c:=countSubString(item,ss)
			result[item] = c
		}
	}
	return result
}

func countSubString(s string, ss []string) int {
	k:=0
	for _,item := range ss{
		if item == s{
			k++
		}
	}
	return k
}

func main() {
	m := wordCount(" as sa as fe we e e e")
	fmt.Println(m)
}


