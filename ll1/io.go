package ll1

import (
	"fmt"
	"io"
)

func ReadChar(in io.Reader) byte {
	buf := make([]byte, 1)

	if _, err := in.Read(buf); err != nil {
		panic(err)
	}
	ch := buf[0]

	if ch >= 'A' && ch <= 'Z' {
		ch = ch - 'A' + 'a'
	}

	fmt.Printf("[%c] %d\n", ch, ch)

	return ch
}
