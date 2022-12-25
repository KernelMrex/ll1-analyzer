package main

import (
	"fmt"
	"os"
	"rec_fall/ll1"
)

func main() {
	if ll1.Process(os.Stdin) {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}
}
