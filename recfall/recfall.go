package recfall

import "io"

func IsValid(in io.Reader) bool {

	return true
}

func ReadChar(in io.Reader) byte {
	ch := make([]byte, 1)
	if _, err := in.Read(ch); err != nil {
		panic(err)
	}
	return ch[0]
}
