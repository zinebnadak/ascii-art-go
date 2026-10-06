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
		for i := 0; i <= len(list); i++ {
			for _, char := range list {
				Forumla(char)
			}
		}
	}
	return ""
}

func eachletter(input string) string {

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . \"your text\"")
		return ""
	} else if os.Args[1] == "" {
		return ""
	}
	var startNumbers []int
	for _, letterpassed := range input {
		if strings.Contains(input, "\\n") {
			split(input)
			break
		} else {
			startNumbers = append(startNumbers, Forumla(letterpassed)) // this one i made just to make sure that we will run and save the formulated number for each and every letter compare the formula and this function with the one before
		}
	}
	return ""
}
func Forumla(TheRune rune) int {
	var OurNumber int
	OurNumber = (int(TheRune) - 32) * 9
	return OurNumber
}

func lineAssemblying(newArr []int, LinesToChooseFrom []string) string {

}

// swedish sounds so oooo uuu
func main() {
	if len(os.Args) != 2 {
		return
	}
	contentFromSample, err := os.ReadFile("sample.txt") // always need the name in double quotes
	if err != nil {
		fmt.Print("STFU")
		return
	} // we need to split because the read gives you the result as 1 bulk string
	LinesToChooseFrom := strings.Split(string(contentFromSample), "\n")
	input := os.Args[1]
	fmt.Print(eachletter(input))

}

//content, err := os.ReadFile("sample.txt")
//if err != nil {
//	fmt.Println("Not reading:", err)
//	return
//}
