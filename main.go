package main

import (
	"fmt"
	"ll1_analyzer/recdescentparser"
	"os"
)

func main() {
	if recdescentparser.Process(os.Stdin) {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}
}
