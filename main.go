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


func print_row(words []string, LinesToChooseFrom []string) { // formula's result gets used directly
	for _, word := range words { // each item in the list
		if word == "" { // just one empty line
			fmt.Println()
			continue
		}
		for row := 0; row < 8; row++ { // 8 rows 
			line := "" // start the row empty 
			for _, letter := range word { // each letter in this item
				line += LinesToChooseFrom[Forumla(letter)+row] // where the letter starts & how far down we are
			}
			fmt.Println(line) // row is full with all letters so print it
		}
	} 
}


// MAIN FUNCTION
func main() {

	// CHECK ARGS
	if len(os.Args) != 2 {
		return
	}

	// READ FILE
	contentFromSample, err := os.ReadFile("sample.txt") // always need the name in double quotes
	
	// ERROR HANDLING 
	if err != nil {
		fmt.Print("STFU")
		return
	} 

	// TAKE USERS INPUT AND START DRAWING
	LinesToChooseFrom := strings.Split(string(contentFromSample), "\n")
	input := os.Args[1]
	print_row([]string{input}, LinesToChooseFrom)
}



//content, err := os.ReadFile("sample.txt")
//if err != nil {
//	fmt.Println("Not reading:", err)
//	return
//}











