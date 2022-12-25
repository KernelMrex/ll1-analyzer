package ll1

import (
	"errors"
	"github.com/golang-collections/collections/stack"
	"io"
)

func ruleProg(in io.Reader, stack *stack.Stack) error {
	if ReadChar(in) != 'P' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != 'R' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != 'O' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != 'G' {
		return errors.New("unexpected char")
	}

	if ReadChar(in) != ' ' {
		return errors.New("unexpected char")
	}

	if ReadChar(in) != 'i' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != 'd' {
		return errors.New("unexpected char")
	}

	if ReadChar(in) != ' ' {
		return errors.New("unexpected char")
	}

	if err := ruleVar(in, stack); err != nil {
		return err
	}

	if ReadChar(in) != ' ' {
		return errors.New("unexpected char")
	}

	if ReadChar(in) != 'b' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != 'e' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != 'g' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != 'i' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != 'n' {
		return errors.New("unexpected char")
	}

	if ReadChar(in) != ' ' {
		return errors.New("unexpected char")
	}

	if err := ruleListSt(in, stack); err != nil {
		return err
	}

	if stack.Pop().(byte) != ' ' {
		return errors.New("unexpected char")
	}

	if ReadChar(in) != 'e' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != 'n' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != 'd' {
		return errors.New("unexpected char")
	}

	return nil
}

func ruleVar(in io.Reader, stack *stack.Stack) error {
	if ReadChar(in) != 'V' {
		return errors.New("unexpected char")
	}

	if ReadChar(in) != 'A' {
		return errors.New("unexpected char")
	}

	if ReadChar(in) != 'R' {
		return errors.New("unexpected char")
	}

	if ReadChar(in) != ' ' {
		return errors.New("unexpected char")
	}

	if err := ruleIdList(in, stack); err != nil {
		return err
	}

	if stack.Len() == 0 {
		return errors.New("unexpected end of file")
	}
	ch := stack.Pop().(byte)

	if ch != ' ' {
		return errors.New("unexpected char")
	}

	if ReadChar(in) != ':' {
		return errors.New("unexpected char")
	}

	if ReadChar(in) != ' ' {
		return errors.New("unexpected char")
	}

	if err := ruleType(in, stack); err != nil {
		return err
	}

	return nil
}

func ruleIdList(in io.Reader, stack *stack.Stack) error {
	if ReadChar(in) != 'i' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != 'd' {
		return errors.New("unexpected char")
	}

	stack.Push(ReadChar(in))
	if err := ruleIdListA(in, stack); err != nil {
		return err
	}

	return nil
}

func ruleIdListA(in io.Reader, stack *stack.Stack) error {
	ch := stack.Pop().(byte)

	if ch == ',' {
		if ReadChar(in) != ' ' {
			return errors.New("unexpected char")
		}
		if ReadChar(in) != 'i' {
			return errors.New("unexpected char")
		}
		if ReadChar(in) != 'd' {
			return errors.New("unexpected char")
		}

		stack.Push(ReadChar(in))
		if err := ruleIdListA(in, stack); err != nil {
			return err
		}
	} else {
		stack.Push(ch)
	}

	return nil
}

func ruleType(in io.Reader, _ *stack.Stack) error {
	ch := ReadChar(in)

	if ch == 'i' {
		if ReadChar(in) != 'n' {
			return errors.New("unexpected char")
		}
		if ReadChar(in) != 't' {
			return errors.New("unexpected char")
		}
	} else if ch == 'f' {
		if ReadChar(in) != 'l' {
			return errors.New("unexpected char")
		}
		if ReadChar(in) != 'o' {
			return errors.New("unexpected char")
		}
		if ReadChar(in) != 'a' {
			return errors.New("unexpected char")
		}
		if ReadChar(in) != 't' {
			return errors.New("unexpected char")
		}
	} else if ch == 'b' {
		if ReadChar(in) != 'o' {
			return errors.New("unexpected char")
		}
		if ReadChar(in) != 'o' {
			return errors.New("unexpected char")
		}
		if ReadChar(in) != 'l' {
			return errors.New("unexpected char")
		}
	} else if ch == 's' {
		if ReadChar(in) != 't' {
			return errors.New("unexpected char")
		}
		if ReadChar(in) != 'r' {
			return errors.New("unexpected char")
		}
		if ReadChar(in) != 'i' {
			return errors.New("unexpected char")
		}
		if ReadChar(in) != 'n' {
			return errors.New("unexpected char")
		}
		if ReadChar(in) != 'g' {
			return errors.New("unexpected char")
		}
	} else {
		return errors.New("unexpected char")
	}

	return nil
}

func ruleExp(in io.Reader, stack *stack.Stack) error {
	if err := ruleT(in, stack); err != nil {
		return err
	}

	return ruleExpA(in, stack)
}

func ruleT(in io.Reader, stack *stack.Stack) error {
	if err := ruleF(in, stack); err != nil {
		return err
	}

	return ruleTA(in, stack)
}

func ruleExpA(in io.Reader, stack *stack.Stack) error {
	if ch := ReadChar(in); ch == ' ' {
		return nil
	} else {
		stack.Push(ch)
	}

	if stack.Pop().(byte) != '+' {
		return errors.New("invalid arg")
	}

	if err := ruleT(in, stack); err != nil {
		return err
	}

	return ruleExpA(in, stack)
}

func ruleF(in io.Reader, stack *stack.Stack) error {
	switch ReadChar(in) {
	case '-':
		return ruleF(in, stack)
	case '(':
		if err := ruleExp(in, stack); err != nil {
			return err
		}

		if ch := ReadChar(in); ch != ')' {
			return errors.New("invalid arg")
		}
		return nil
	case 'i':
		if ch := ReadChar(in); ch != 'd' {
			return errors.New("invalid arg")
		}
		return nil
	case 'n':
		if ch := ReadChar(in); ch != 'u' {
			return errors.New("invalid arg")
		}
		if ch := ReadChar(in); ch != 'm' {
			return errors.New("invalid arg")
		}
		return nil
	default:
		return nil
	}
}

func ruleTA(in io.Reader, stack *stack.Stack) error {
	if ch := ReadChar(in); ch == ' ' {
		return nil
	} else {
		stack.Push(ch)
	}

	if stack.Pop().(byte) != '+' {
		return errors.New("invalid arg")
	}

	err := ruleF(in, stack)
	if err != nil {
		return errors.New("invalid arg")
	}

	return ruleTA(in, stack)
}

func ruleListSt(in io.Reader, stack *stack.Stack) error {
	ch := ReadChar(in)
	stack.Push(ch)
	if err := ruleSt(in, stack); err != nil {
		return err
	}

	if ReadChar(in) != ' ' {
		return errors.New("unexpected char")
	}

	stack.Push(ReadChar(in))
	if err := ruleListStA(in, stack); err != nil {
		return err
	}

	return nil
}

func ruleListStA(in io.Reader, stack *stack.Stack) error {
	ch := stack.Peek().(byte)

	if ch != 'r' && ch != 'w' && ch != 'i' { // Check if FIRST-relation vector for ST does not match
		return nil
	}

	if err := ruleSt(in, stack); err != nil {
		return err
	}

	stack.Push(ReadChar(in))
	if err := ruleListStA(in, stack); err != nil {
		return err
	}

	return nil
}

func ruleSt(in io.Reader, stack *stack.Stack) error {
	ch := stack.Peek().(byte)

	if ch == 'r' {
		if err := ruleRead(in, stack); err != nil {
			return err
		}
	} else if ch == 'w' {
		if err := ruleWrite(in, stack); err != nil {
			return err
		}
	} else if ch == 'i' {
		if err := ruleAssign(in, stack); err != nil {
			return err
		}
	} else {
		return errors.New("unexpected char")
	}

	return nil
}

func ruleRead(in io.Reader, stack *stack.Stack) error {
	ch := stack.Pop().(byte)
	if ch != 'r' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != 'e' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != 'a' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != 'd' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != '(' {
		return errors.New("unexpected char")
	}
	if err := ruleIdList(in, stack); err != nil {
		return err
	}
	if stack.Pop().(byte) != ')' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != ';' {
		return errors.New("unexpected char")
	}
	return nil
}

func ruleWrite(in io.Reader, stack *stack.Stack) error {
	ch := stack.Pop().(byte)
	if ch != 'w' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != 'r' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != 'i' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != 't' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != 'e' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != '(' {
		return errors.New("unexpected char")
	}
	if err := ruleIdList(in, stack); err != nil {
		return err
	}
	if ReadChar(in) != ')' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != ';' {
		return errors.New("unexpected char")
	}
	return nil
}

func ruleAssign(in io.Reader, stack *stack.Stack) error {
	ch := stack.Pop().(byte)
	if ch != 'i' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != 'd' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != ' ' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != ':' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != '=' {
		return errors.New("unexpected char")
	}
	if ReadChar(in) != ' ' {
		return errors.New("unexpected char")
	}
	if err := ruleExp(in, stack); err != nil {
		return err
	}
	if ReadChar(in) != ';' {
		return errors.New("unexpected char")
	}
	return nil
}
