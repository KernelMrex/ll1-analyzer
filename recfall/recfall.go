package recfall

import (
	"container/list"
	"io"
)

func IsValid(in io.Reader) bool {
	if err := fnProg(in, list.New()); err != nil {
		return false
	}

	return true
}

func ReadChar(in io.Reader) byte {
	ch := make([]byte, 1)
	if _, err := in.Read(ch); err != nil {
		panic(err)
	}
	return ch[0]
}
