package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	mm := make(map[string]int)
	fields := strings.Split(s, " ")

	for i := 0; i < len(fields); i++ {
		v := fields[i]
		if mm[v] != 0 {
			mm[v] = mm[v] + 1
		} else {
			mm[v] = 1
		}
	}

	return mm
}

func main() {
	wc.Test(WordCount)
}
