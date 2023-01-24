package ll1

import (
	"fmt"
	"io"
)

type Reader interface {
	Next() byte
}

type skipSpacesReader struct {
	in io.Reader
}

func NewReader(reader io.Reader) Reader {
	return &skipSpacesReader{
		in: reader,
	}
}

func (r *skipSpacesReader) Next() byte {
	buf := make([]byte, 1)

	//Loop:
	if _, err := r.in.Read(buf); err != nil {
		panic(err)
	}
	ch := buf[0]

	//if isSpace(ch) {
	//	goto Loop
	//}

	fmt.Printf("[%c] %d\n", ch, ch)
	return toLower(ch)
}

//func isSpace(ch byte) bool {
//	return ch == ' ' || ch == '\n'
//}

func toLower(ch byte) byte {
	if ch >= 'A' && ch <= 'Z' {
		ch = ch - 'A' + 'a'
	}
	return ch
}
