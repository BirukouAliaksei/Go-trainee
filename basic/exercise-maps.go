package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	var values = strings.Fields(s)
	m := make(map[string]int)

	for _, s = range values {
		count, ok := m[s]
		if ok {
			m[s] = count + 1
		} else {
			m[s] = 1
		}
	}
	return m

}

func main() {
	wc.Test(WordCount)
}
