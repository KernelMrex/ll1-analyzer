package main

import (
	"fmt"
	"ll1_analyzer/ll1"
	"os"
)

func main() {
	if ll1.Process(os.Stdin) {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}
}
