package recfall

import (
	"container/list"
	"errors"
	"io"
)

func fnProg(in io.Reader, stack *list.List) error {
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

	return nil
}

func fnVar(in io.Reader, stack *list.List) error {
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

func fnIdList(in io.Reader, stack *list.List) error {
	if ReadChar(in) != 'i' {
		return errors.New("unexpected char")
	}

	if ReadChar(in) != 'd' {
		return errors.New("unexpected char")
	}

	if ReadChar(in) != ' ' {
		return errors.New("unexpected char")
	}

	if err := fnIdListA(in, stack); err != nil {
		return err
	}

	e := stack.Front()
	ch := e.Value
	stack.Remove(e)

	if ch != ':' {
		return errors.New("unexpected char")
	}

	if err := fnType(in, stack); err != nil {
		return err
	}

	return nil
}

func fnIdListA(in io.Reader, stack *list.List) error {
	ch := ReadChar(in)

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
	} else {
		stack.PushFront(ch)
	}

	return nil
}

func fnType(in io.Reader, stack *list.List) error {
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
