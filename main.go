package main

import (
	"flag"
	"fmt"
	"strings"
)

func readInput() string {
	flag.Parse()
	return strings.Join(flag.Args(), " ")
}

func countWords(text string) int {
	if text == "" {
		return 0
	}
	return strings.Count(text, " ") + 1
}

func main() {
	fmt.Println(
		countWords(
			readInput(),
		),
	)
}
