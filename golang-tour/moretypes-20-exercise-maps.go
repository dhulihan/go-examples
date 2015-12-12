package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	occurranceMap := map[string]int{}
	phraseAry := strings.Fields(s) 
	fmt.Println(phraseAry)
	for _, word := range phraseAry {
	 occurranceMap[word]++ 
	}
	return occurranceMap
}

func main() {
	wc.Test(WordCount)
}