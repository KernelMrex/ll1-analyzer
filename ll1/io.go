package ll1

import (
	"fmt"
	"io"
)

func ReadChar(in io.Reader) byte {
	ch := make([]byte, 1)
	if _, err := in.Read(ch); err != nil {
		panic(err)
	}

	fmt.Printf("%c %d\n", ch, ch)

	return ch[0]
}

//type Reader struct {
//	stack *stack.Stack
//	in    io.Reader
//}
//
//func NewReader(in io.Reader) *Reader {
//	return &Reader{
//		stack: stack.New(),
//		in:    in,
//	}
//}
//
//func (r *Reader) ReadChar() byte {
//	if r.stack.Len() > 0 {
//		return r.stack.Pop().(byte)
//	}
//
//	ch := make([]byte, 1)
//	if _, err := r.in.Read(ch); err != nil {
//		panic(err)
//	}
//
//	return ch[0]
//}
//
//func (r *Reader) ReadWord() string {
//	var buf []byte
//
//	chBuf := make([]byte, 1)
//	for {
//		_, err := r.in.Read(chBuf)
//		if err != nil {
//			if err == io.EOF {
//				return string(buf)
//			}
//			panic(err)
//		}
//
//		ch := chBuf[0]
//		if !isAlphaNumeric(ch) {
//			return string(buf)
//		}
//		buf = append(buf, ch)
//	}
//}
//
//func isAlphaNumeric(ch byte) bool {
//	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '0')
//}
