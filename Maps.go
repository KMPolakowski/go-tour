package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	var counts map[string]int = make(map[string]int)
	var words []string = strings.Fields(s)

	for _, x := range words {
		var count int
		for _, y := range words {
			if y == x {
				count++
			}
		}

		//could also just overwrite
		_, counted := counts[x]

		if !counted {
			counts[x] = count
		}
	}

	return counts
}

func main() {
	wc.Test(WordCount)
}
