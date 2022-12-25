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
