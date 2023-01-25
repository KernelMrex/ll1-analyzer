package recdescentparser

import (
	"errors"
	"github.com/golang-collections/collections/stack"
)

func ruleProg(in Reader, stack *stack.Stack) error {
	in.SkipSpaces()
	if in.Next() != 'p' {
		return errors.New("unexpected char")
	}
	if in.Next() != 'r' {
		return errors.New("unexpected char")
	}
	if in.Next() != 'o' {
		return errors.New("unexpected char")
	}
	if in.Next() != 'g' {
		return errors.New("unexpected char")
	}

	in.SkipSpaces()

	if in.Next() != 'i' {
		return errors.New("unexpected char")
	}
	if in.Next() != 'd' {
		return errors.New("unexpected char")
	}

	in.SkipSpaces()

	if err := ruleVar(in, stack); err != nil {
		return err
	}

	in.SkipSpaces()

	if in.Next() != 'b' {
		return errors.New("unexpected char")
	}
	if in.Next() != 'e' {
		return errors.New("unexpected char")
	}
	if in.Next() != 'g' {
		return errors.New("unexpected char")
	}
	if in.Next() != 'i' {
		return errors.New("unexpected char")
	}
	if in.Next() != 'n' {
		return errors.New("unexpected char")
	}

	in.SkipSpaces()

	if err := ruleListSt(in, stack); err != nil {
		return err
	}

	if stack.Pop().(byte) != 'e' {
		return errors.New("unexpected char")
	}
	if in.Next() != 'n' {
		return errors.New("unexpected char")
	}
	if in.Next() != 'd' {
		return errors.New("unexpected char")
	}

	return nil
}

func ruleVar(in Reader, stack *stack.Stack) error {
	if in.Next() != 'v' {
		return errors.New("unexpected char")
	}

	if in.Next() != 'a' {
		return errors.New("unexpected char")
	}

	if in.Next() != 'r' {
		return errors.New("unexpected char")
	}

	in.SkipSpaces()

	if err := ruleIdList(in, stack); err != nil {
		return err
	}

	if stack.Len() == 0 {
		return errors.New("unexpected end of file")
	}

	if stack.Pop().(byte) != ':' {
		return errors.New("unexpected char")
	}

	in.SkipSpaces()

	if err := ruleType(in, stack); err != nil {
		return err
	}

	return nil
}

func ruleIdList(in Reader, stack *stack.Stack) error {
	if in.Next() != 'i' {
		return errors.New("unexpected char")
	}
	if in.Next() != 'd' {
		return errors.New("unexpected char")
	}

	in.SkipSpaces()

	if err := ruleIdListA(in, stack); err != nil {
		return err
	}

	return nil
}

func ruleIdListA(in Reader, stack *stack.Stack) error {
	ch := in.Next()

	in.SkipSpaces()

	if ch == ',' {

		if in.Next() != 'i' {
			return errors.New("unexpected char")
		}
		if in.Next() != 'd' {
			return errors.New("unexpected char")
		}

		in.SkipSpaces()

		if err := ruleIdListA(in, stack); err != nil {
			return err
		}
	} else {
		stack.Push(ch)
	}

	return nil
}

func ruleType(in Reader, _ *stack.Stack) error {
	ch := in.Next()

	if ch == 'i' {
		if in.Next() != 'n' {
			return errors.New("unexpected char")
		}
		if in.Next() != 't' {
			return errors.New("unexpected char")
		}
	} else if ch == 'f' {
		if in.Next() != 'l' {
			return errors.New("unexpected char")
		}
		if in.Next() != 'o' {
			return errors.New("unexpected char")
		}
		if in.Next() != 'a' {
			return errors.New("unexpected char")
		}
		if in.Next() != 't' {
			return errors.New("unexpected char")
		}
	} else if ch == 'b' {
		if in.Next() != 'o' {
			return errors.New("unexpected char")
		}
		if in.Next() != 'o' {
			return errors.New("unexpected char")
		}
		if in.Next() != 'l' {
			return errors.New("unexpected char")
		}
	} else if ch == 's' {
		if in.Next() != 't' {
			return errors.New("unexpected char")
		}
		if in.Next() != 'r' {
			return errors.New("unexpected char")
		}
		if in.Next() != 'i' {
			return errors.New("unexpected char")
		}
		if in.Next() != 'n' {
			return errors.New("unexpected char")
		}
		if in.Next() != 'g' {
			return errors.New("unexpected char")
		}
	} else {
		return errors.New("unexpected char")
	}

	return nil
}

func ruleListSt(in Reader, stack *stack.Stack) error {
	ch := in.Next()
	stack.Push(ch)
	if err := ruleSt(in, stack); err != nil {
		return err
	}

	in.SkipSpaces()

	stack.Push(in.Next())
	if err := ruleListStA(in, stack); err != nil {
		return err
	}

	return nil
}

func ruleListStA(in Reader, stack *stack.Stack) error {
	ch := stack.Peek().(byte)

	if ch != 'r' && ch != 'w' && ch != 'i' { // Check if FIRST-relation vector for ST does not match
		return nil
	}

	if err := ruleSt(in, stack); err != nil {
		return err
	}

	in.SkipSpaces()

	stack.Push(in.Next())
	if err := ruleListStA(in, stack); err != nil {
		return err
	}

	return nil
}

func ruleSt(in Reader, stack *stack.Stack) error {
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

func ruleRead(in Reader, stack *stack.Stack) error {
	ch := stack.Pop().(byte)
	if ch != 'r' {
		return errors.New("unexpected char")
	}
	if in.Next() != 'e' {
		return errors.New("unexpected char")
	}
	if in.Next() != 'a' {
		return errors.New("unexpected char")
	}
	if in.Next() != 'd' {
		return errors.New("unexpected char")
	}
	if in.Next() != '(' {
		return errors.New("unexpected char")
	}
	in.SkipSpaces()
	if err := ruleIdList(in, stack); err != nil {
		return err
	}
	in.SkipSpaces()
	if stack.Pop().(byte) != ')' {
		return errors.New("unexpected char")
	}
	if in.Next() != ';' {
		return errors.New("unexpected char")
	}
	return nil
}

func ruleWrite(in Reader, stack *stack.Stack) error {
	if stack.Pop().(byte) != 'w' {
		return errors.New("unexpected char")
	}
	if in.Next() != 'r' {
		return errors.New("unexpected char")
	}
	if in.Next() != 'i' {
		return errors.New("unexpected char")
	}
	if in.Next() != 't' {
		return errors.New("unexpected char")
	}
	if in.Next() != 'e' {
		return errors.New("unexpected char")
	}
	if in.Next() != '(' {
		return errors.New("unexpected char")
	}
	in.SkipSpaces()
	if err := ruleIdList(in, stack); err != nil {
		return err
	}
	in.SkipSpaces()
	if stack.Pop().(byte) != ')' {
		return errors.New("unexpected char")
	}
	if in.Next() != ';' {
		return errors.New("unexpected char")
	}
	return nil
}

func ruleAssign(in Reader, stack *stack.Stack) error {
	if stack.Pop().(byte) != 'i' {
		return errors.New("unexpected char")
	}
	if in.Next() != 'd' {
		return errors.New("unexpected char")
	}
	in.SkipSpaces()
	if in.Next() != ':' {
		return errors.New("unexpected char")
	}
	if in.Next() != '=' {
		return errors.New("unexpected char")
	}
	if err := ruleExp(in, stack); err != nil {
		return err
	}
	if stack.Pop().(byte) != ';' {
		return errors.New("unexpected char")
	}
	return nil
}

func ruleExp(in Reader, stack *stack.Stack) error {
	if err := ruleT(in, stack); err != nil {
		return err
	}

	in.SkipSpaces()
	if err := ruleExpA(in, stack); err != nil {
		return err
	}

	return nil
}

func ruleT(in Reader, stack *stack.Stack) error {
	in.SkipSpaces()
	if err := ruleF(in, stack); err != nil {
		return err
	}

	in.SkipSpaces()
	return ruleTA(in, stack)
}

func ruleExpA(in Reader, stack *stack.Stack) error {
	ch := stack.Pop().(byte)
	if ch == ';' || ch == ')' {
		stack.Push(ch)
		return nil
	}

	if ch != '+' {
		return errors.New("unexpected operation")
	}
	in.SkipSpaces()
	if err := ruleT(in, stack); err != nil {
		return errors.New("invalid arg")
	}

	in.SkipSpaces()
	return ruleExpA(in, stack)
}

func ruleF(in Reader, stack *stack.Stack) error {
	switch in.Next() {
	case '-':
		in.SkipSpaces()
		return ruleF(in, stack)
	case '(':
		if err := ruleExp(in, stack); err != nil {
			return err
		}
		if stack.Pop().(byte) != ')' {
			return errors.New("unexpected char")
		}
		return nil
	case 'i':
		if in.Next() != 'd' {
			return errors.New("unexpected char")
		}
		return nil
	case 'n':
		if in.Next() != 'u' {
			return errors.New("unexpected char")
		}
		if in.Next() != 'm' {
			return errors.New("unexpected char")
		}
		return nil
	default:
		return errors.New("unexpected char")
	}
}

func ruleTA(in Reader, stack *stack.Stack) error {
	ch := in.Next()
	if ch == ';' || ch == ')' || ch == '+' {
		stack.Push(ch)
		return nil
	}
	if ch != '*' {
		return errors.New("unexpected operation")
	}
	in.SkipSpaces()
	if err := ruleF(in, stack); err != nil {
		return errors.New("invalid arg")
	}

	in.SkipSpaces()
	return ruleTA(in, stack)
}
