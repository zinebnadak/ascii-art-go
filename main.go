// checks arguments and prints

package main

import (
	"fmt"
	"os"
	"strings"
)

func split(input string) string { // builds the list of items

	list := ""
	for index, charachter := range input {
		if input[index] == '\\' && input[index+1] == 'n' { // it could need to be converted
			list += ("\n")
		}
		list += string(charachter)

	}
	return ""
}

func test(input string) string {

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . \"your text\"")
		return ""
	} else if os.Args[1] == "" {
		return ""
	}
	for _, letterpassed := range input {
		if strings.Contains(input, "\\n") {
			split(input)
			break
		} else {
			Forumla(letterpassed)
		}
	}
	return ""
}
func Forumla(TheRune rune) int {
	var OurNumber int
	OurNumber = (int(TheRune) - 32) * 9
	var newArr []int
	newArr = append(newArr, OurNumber)
	return 1
}

// swedish sounds so oooo uuu
func main() {
	if len(os.Args) != 2 {
		return
	} else {
		input := os.Args[1]
		fmt.Print(test(input))
	}

}

//content, err := os.ReadFile("sample.txt")
//if err != nil {
//	fmt.Println("Not reading:", err)
//	return
//}
