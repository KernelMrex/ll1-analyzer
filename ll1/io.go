package ll1

import (
	"io"
)

//type StackReader struct {
//	stack stack.Stack
//}
//
//func (r *StackReader) ReadChar() {
//
//}
//
//func (r *StackReader) ReadSeparator() {
//
//}

func ReadChar(in io.Reader) byte {
	ch := make([]byte, 1)
	if _, err := in.Read(ch); err != nil {
		panic(err)
	}
	return ch[0]
}
