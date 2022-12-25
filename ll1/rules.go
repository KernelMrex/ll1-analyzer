package ll1

import (
	"errors"
	"github.com/golang-collections/collections/stack"
	"io"
	"strings"
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

	err := ruleListSt(in, stack)
	if err != nil {
		return err
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

	if stack.Len() == 0 {
		return errors.New("unexpected end of file")
	}
	ch := stack.Pop().(byte)

	if ch != ':' {
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
	err := ruleT(in, stack)
	if err != nil {
		return err
	}

	return ruleExpA(in, stack)
}

func ruleT(in io.Reader, stack *stack.Stack) error {
	err := ruleF(in, stack)
	if err != nil {
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
	err := ruleSt(in, stack)
	if err != nil {
		return err
	}

	return ruleListStA(in, stack)
}

func ruleListStA(in io.Reader, stack *stack.Stack) error {
	if ch := ReadChar(in); ch == ' ' {
		return nil
	} else {
		stack.Push(ch)
	}

	err := ruleSt(in, stack)
	if err != nil {
		return err
	}

	return ruleListStA(in, stack)
}

func ruleSt(in io.Reader, stack *stack.Stack) error {
	var word []byte
	for ch := ReadChar(in); ch != '(' && ch != ':'; ch = ReadChar(in) {
		word = append(word, ch)
	}

	strWord := string(word)
	strWord = strings.ToLower(strWord)

	switch strWord {
	case "read":
		return ruleRead(in, stack)
	case "write":
		return ruleWrite(in, stack)
	case "id":
		return ruleAssign(in, stack)
	default:
		return errors.New("excepted read, write or assign")
	}
}

func ruleRead(in io.Reader, stack *stack.Stack) error {
	err := ruleIdList(in, stack)
	if err != nil {
		return err
	}

	if ch := ReadChar(in); ch != ')' {
		return errors.New("excepted read")
	}
	return nil
}

func ruleWrite(in io.Reader, stack *stack.Stack) error {
	err := ruleIdList(in, stack)
	if err != nil {
		return err
	}

	if ch := ReadChar(in); ch != ')' {
		return errors.New("excepted assign")
	}
	return nil
}

func ruleAssign(in io.Reader, stack *stack.Stack) error {
	if ch := ReadChar(in); ch != '=' {
		return errors.New("excepted assign")
	}

	if ch := ReadChar(in); ch != ' ' {
		return errors.New("excepted assign")
	}

	return ruleExp(in, stack)
}
