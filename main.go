// checks arguments and prints

package main

import (
	"fmt"
	"os"
	"strings"
)

// function BUILDS LIST OF ITEMS
func split(input string) string {  
	list := ""
	for index, charachter := range input {
		if input[index] == '\\' && input[index+1] == 'n' { //it could need to be converted
			list += ("\n")
		}
		list += string(charachter)
	}
	return list
}

// function CALCULATES THE STARTING POINT
func Forumla(TheRune rune) int {
	return (int(TheRune)-32)*9.   // maybe +1
}

// function PRINTS EACH ROW OF THE LETTERS ITEM BY ITEM THEN LETTER BY LETTER FROM THE LIST
func print_row(words []string, LinesToChooseFrom []string) { 
	for _, word := range words { // each ITEM in the list
		if word == "" { // if the item is empty, print a blank line and continue to next item
			fmt.Println()
			continue
		}
		for row := 0; row < 8; row++ { // 8 rows 
			line := "" // start the row empty 
			for _, letter := range word { // each LETTER in the item
				line += LinesToChooseFrom[Forumla(letter)+row] // where the letter starts & how far down we are
			}
			fmt.Println(line) 
		}
	} 
}

// MAIN FUNCTION
func main() {

	// CHECK ARGS
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . \"your text\"")
		return
	}

	// READ FILE
	contentFromSample, err := os.ReadFile("standard.txt") // always need the name in double quotes
	
	// ERROR HANDLING 
	if err != nil {
		fmt.Println("Error:", err)
		return
	} 

	// TAKE USERS INPUT AND START DRAWING
	LinesToChooseFrom := strings.Split(string(contentFromSample), "\n")
	input := os.Args[1]
	print_row([]string{input}, LinesToChooseFrom)
}














