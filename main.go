// checks arguments and prints

package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . \"your text\"")
		return
	}
	input := os.Args[1]
	for _, char := range input {
		Forumla(char)
	}
}

func Forumla(TheRune rune) int {
	ourRune := TheRune
	OurNumber := (ourRune - 32) * 9
}

