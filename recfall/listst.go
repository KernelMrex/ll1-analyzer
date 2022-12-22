package recfall

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func parseListSt() error {
	return nil
}

func ParseSt(in io.Reader) error {
	var word []byte
	for ch := ReadChar(in); ch != '(' && ch != ':'; ch = ReadChar(in) {
		fmt.Println(ch)
		word = append(word, ch)
	}

	strWord := string(word)
	strWord = strings.ToLower(strWord)

	switch strWord {
	case "read":
		return parseRead(in)
	case "write":
		return parseWrite(in)
	case "id":
		return parseAssign(in)
	default:
		return errors.New("excepted read, write or assign")
	}
}

func parseRead(in io.Reader) error {
	//err = parseIDList()

	if ch := ReadChar(in); ch != ')' {
		return errors.New("excepted read")
	}
	return nil
}

func parseWrite(in io.Reader) error {
	//err = parseIDList()

	if ch := ReadChar(in); ch != ')' {
		return errors.New("excepted assign")
	}
	return nil
}

func parseAssign(in io.Reader) error {
	if ch := ReadChar(in); ch != '=' {
		return errors.New("excepted assign")
	}

	if ch := ReadChar(in); ch != ' ' {
		return errors.New("excepted assign")
	}

	return nil
	// return parseExp
}
