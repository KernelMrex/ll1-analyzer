package recfall

import (
	"errors"
	"github.com/golang-collections/collections/stack"
	"io"
)

func fnProg(in io.Reader, stack *stack.Stack) error {
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

	if err := fnVar(in, stack); err != nil {
		return err
	}

	err := parseListSt(in, stack)
	if err != nil {
		return err
	}

	return nil
}

func fnVar(in io.Reader, stack *stack.Stack) error {
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

	if err := fnIdList(in, stack); err != nil {
		return err
	}

	return nil
}

func fnIdList(in io.Reader, stack *stack.Stack) error {
	if ReadChar(in) != 'i' {
		return errors.New("unexpected char")
	}

	if ReadChar(in) != 'd' {
		return errors.New("unexpected char")
	}

	stack.Push(ReadChar(in))
	if err := fnIdListA(in, stack); err != nil {
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

	if err := fnType(in, stack); err != nil {
		return err
	}

	return nil
}

func fnIdListA(in io.Reader, stack *stack.Stack) error {
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
		if err := fnIdListA(in, stack); err != nil {
			return err
		}
	} else {
		stack.Push(ch)
	}

	return nil
}

func fnType(in io.Reader, stack *stack.Stack) error {
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
