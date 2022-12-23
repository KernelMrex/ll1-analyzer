package recfall

import (
	"errors"
	"fmt"
	"github.com/golang-collections/collections/stack"
	"io"
	"strings"
)

func parseListSt(in io.Reader, stack *stack.Stack) error {
	err := ParseSt(in, stack)
	if err != nil {
		return err
	}

	return parseB(in, stack)
}

func parseB(in io.Reader, stack *stack.Stack) error {
	if ch := ReadChar(in); ch == ' ' {
		return nil
	} else {
		stack.Push(ch)
	}

	err := ParseSt(in, stack)
	if err != nil {
		return err
	}

	return parseB(in, stack)
}

func ParseSt(in io.Reader, stack *stack.Stack) error {
	var word []byte
	for ch := ReadChar(in); ch != '(' && ch != ':'; ch = ReadChar(in) {
		fmt.Println(ch)
		word = append(word, ch)
	}

	strWord := string(word)
	strWord = strings.ToLower(strWord)

	switch strWord {
	case "read":
		return parseRead(in, stack)
	case "write":
		return parseWrite(in, stack)
	case "id":
		return parseAssign(in, stack)
	default:
		return errors.New("excepted read, write or assign")
	}
}

func parseRead(in io.Reader, stack *stack.Stack) error {
	err := fnIdList(in, stack)
	if err != nil {
		return err
	}

	if ch := ReadChar(in); ch != ')' {
		return errors.New("excepted read")
	}
	return nil
}

func parseWrite(in io.Reader, stack *stack.Stack) error {
	err := fnIdList(in, stack)
	if err != nil {
		return err
	}

	if ch := ReadChar(in); ch != ')' {
		return errors.New("excepted assign")
	}
	return nil
}

func parseAssign(in io.Reader, stack *stack.Stack) error {
	if ch := ReadChar(in); ch != '=' {
		return errors.New("excepted assign")
	}

	if ch := ReadChar(in); ch != ' ' {
		return errors.New("excepted assign")
	}

	return parseExp(in, stack)
}
