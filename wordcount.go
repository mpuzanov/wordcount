package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		s := os.Args[1]
		fmt.Println(wordcount(s))
	}
}

func wordcount(s string) int {
	return len(strings.Fields(s))
}
