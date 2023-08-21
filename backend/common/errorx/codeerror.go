package errorx

import (
	"errors"
	"fmt"
)

type (
	CodeError interface {
		error
		Status() int
		Code() int
	}

	codeError struct {
		code   int
		status int
		desc   string
	}
)

func (err *codeError) Code() int {
	return err.code
}

func (err *codeError) Status() int {
	return err.status
}

func (err *codeError) Error() string {
	return err.desc
}

func (err *codeError) String() string {
	return fmt.Sprintf("Status: %d, Code: %d, Desc: %s", err.status, err.code, err.desc)
}

func NewCodeError(code, status int, desc string) CodeError {
	return &codeError{
		code:   code,
		status: status,
		desc:   desc,
	}
}

func FromError(err error) (CodeError, bool) {
	if err == nil {
		return nil, false
	}
	var ce CodeError
	if ok := errors.As(err, &ce); ok {
		return ce, ok
	}
	return nil, false
}

func Error400(desc string) CodeError {
	return NewCodeError(400, 400, desc)
}

func Error500(desc string) CodeError {
	return NewCodeError(500, 500, desc)
}
