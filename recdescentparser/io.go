package recdescentparser

import (
	"fmt"
	"io"
)

type Reader interface {
	Next() byte
	SkipSpaces()
}

type skipSpacesReader struct {
	in  io.Reader
	buf byte
}

func NewReader(reader io.Reader) Reader {
	return &skipSpacesReader{
		in:  reader,
		buf: 0,
	}
}

func (r *skipSpacesReader) Next() byte {
	ch := r.ReadChar()
	fmt.Printf("[%c] %d\n", ch, ch)
	return toLower(ch)
}

func (r *skipSpacesReader) SkipSpaces() {
Loop:
	ch := r.ReadChar()
	if isSpace(ch) {
		goto Loop
	}
	r.buf = ch
}

func (r *skipSpacesReader) ReadChar() byte {
	var ch byte
	if r.buf != 0 {
		ch = r.buf
		r.buf = 0
	} else {
		buf := make([]byte, 1)
		if _, err := r.in.Read(buf); err != nil {
			panic(err)
		}
		ch = buf[0]
	}
	return ch
}

func isSpace(ch byte) bool {
	return ch == ' ' || ch == '\n' || ch == '\t'
}

func toLower(ch byte) byte {
	if ch >= 'A' && ch <= 'Z' {
		ch = ch - 'A' + 'a'
	}
	return ch
}
