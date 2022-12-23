package recfall

import (
	"errors"
	"github.com/golang-collections/collections/stack"
	"io"
)

func parseExp(in io.Reader, stack *stack.Stack) error {
	err := parseT(in, stack)
	if err != nil {
		return err
	}

	return parseC(in, stack)
}

func parseT(in io.Reader, stack *stack.Stack) error {
	err := parseF(in, stack)
	if err != nil {
		return err
	}

	return parseD(in, stack)
}

func parseC(in io.Reader, stack *stack.Stack) error {
	if ch := ReadChar(in); ch == ' ' {
		return nil
	} else {
		stack.Push(ch)
	}

	if stack.Pop().(byte) != '+' {
		return errors.New("invalid arg")
	}

	err := parseT(in, stack)
	if err != nil {
		return err
	}

	return parseC(in, stack)
}

func parseF(in io.Reader, stack *stack.Stack) error {
	ch := ReadChar(in)
	switch ch {
	case '-':
		return parseF(in, stack)
	case '(':
		err := parseExp(in, stack)
		if err != nil {
			return err
		}

		if ch = ReadChar(in); ch != ')' {
			return errors.New("invalid arg")
		}
		return nil
	case 'i':
		if ch = ReadChar(in); ch != 'd' {
			return errors.New("invalid arg")
		}
		return nil
	case 'n':
		if ch = ReadChar(in); ch != 'u' {
			return errors.New("invalid arg")
		}
		if ch = ReadChar(in); ch != 'm' {
			return errors.New("invalid arg")
		}
		return nil
	default:
		return nil
	}
}

func parseD(in io.Reader, stack *stack.Stack) error {
	if ch := ReadChar(in); ch == ' ' {
		return nil
	} else {
		stack.Push(ch)
	}

	if stack.Pop().(byte) != '+' {
		return errors.New("invalid arg")
	}

	err := parseF(in, stack)
	if err != nil {
		return errors.New("invalid arg")
	}

	return parseD(in, stack)
}
