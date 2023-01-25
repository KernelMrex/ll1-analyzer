package recdescentparser

import (
	"github.com/golang-collections/collections/stack"
	"io"
)

func Process(in io.Reader) bool {
	if err := ruleProg(NewReader(in), stack.New()); err != nil {
		return false
	}
	return true
}
