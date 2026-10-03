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

}
