package recfall

import (
	"errors"
	"io"
)

func prog(in io.Reader) error {
	if ReadChar(in) != 'P' {
		return errors.New("unexpected char")
	}

	return nil
}
