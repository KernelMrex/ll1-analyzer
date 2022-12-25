package ll1

import (
	"github.com/golang-collections/collections/stack"
	"io"
)

func Process(in io.Reader) bool {
	if err := ruleProg(in, stack.New()); err != nil {
		return false
	}
	return true
}
