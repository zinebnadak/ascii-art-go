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
			list += ("\\n")
		}
		list += string(charachter)
	}
	return list
}

// function CALCULATES THE STARTING POINT
func Forumla(TheRune rune) int {
	return (int(TheRune)-32)*9 + 1 // maybe +1 // yes it is 
	// run this command please 
	// head -3 standard.txt | cat -e
	// this one would read the first 3 lines in the sample
	// the first one has no nothing literally if u clicked on it u will be at the start 
	// but starting from the second line u will see that there are spaces check them first 
	// by doing + 1 im just skipping this lines "first one the one i dont need " because the drawing starts after it 

}

// function PRINTS EACH ROW OF THE LETTERS ITEM BY ITEM THEN LETTER BY LETTER FROM THE LIST
func print_row(words []string, LinesToChooseFrom []string) {
	for _, word := range words { // each ITEM in the list
		if word == "" { // if the item is empty, print a blank line and continue to next item
			fmt.Println()
			continue
		}
		for row := 0; row < 8; row++ { // 8 rows
			line := ""                    // start the row empty
			for _, letter := range word { // each LETTER in the item
				line += LinesToChooseFrom[Forumla(letter)+row] // where the letter starts & how far down we are,
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
	LinesToChooseFrom := strings.Split(string(contentFromSample), "\n")
	// let me expalin this here
	// without split it and with
	// with split lets say that our sample has those values
	// A-row1 A-row2 B-row1 B-row2
	// after doing split it would be line that
	// A-row 1
	// A-row 2
	// B-row 1
	// B-row 2
	// the new line is my cut point to seperate them

	// without split means that everything would be in one string
	// so once we loop it would like that
	// 'B'
	// '-'
	// 'r'
	// 'o'
	// ' '
	// '1' etc etc
	// even if you would work with it like that because this is what u want, it wont compile why
	// because line in the row function is a string and we cannot store a byte there "byte is the value of whats given to us from read"
	// even you forced it by casting it to a string it would give you a random character how and why
	// A-row1⏎A-row2⏎B-row1⏎B-row2 consider that this is what is in our sample
	// position:  0 1 2 3 4 5 6 7 8 9
	//character:  A - r o w 1 ⏎ A - r
	// numbers are the same in each of them split and without but the main differnce that without split
	// with split our list has in list[2] it has B-row1
	//  without split the text[2] would get u this 'r'
	// so if im passing "hello world"
	// i want this  hello all in one line so i have to do list to make it "hello" and store it in line number 1
	// if i wont use split he will take each letter and save it in it own line
	// so hello would be
	// h
	// e
	// l
	// l
	// o

	// TAKE USERS INPUT AND START DRAWING
	input := os.Args[1]
	if input == "" {
		return
	} // both will handle if it starts with \n or nothing before going through anything to save time
	if input == "\\n" {
		fmt.Println()
		return
	}
	word := strings.Split(input, "\\n")
	print_row(word, LinesToChooseFrom)
}


// WP ZEYNAB 