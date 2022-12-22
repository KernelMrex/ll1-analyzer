package main

import (
	"fmt"
	"os"
	"rec_fall/recfall"
)

func main() {
	if recfall.IsValid(os.Stdin) {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}
}
