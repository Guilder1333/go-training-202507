package statuserr

import "errors"

type msgError struct {
	msg     string
	wrapped error
}

func (e *msgError) Error() string {
	return e.wrapped.Error()
}

func (e *msgError) Unwrap() error {
	return e.wrapped
}

func SetMessage(err error, msg string) error {
	return &msgError{
		msg:     msg,
		wrapped: err,
	}
}

func GetMessage(err error) (string, bool) {
	var me *msgError
	if errors.As(err, &me) {
		return me.msg, true
	}
	return "", false
}
