package statuserr

import "errors"

type statusError struct {
	kind    ErrorKind
	wrapped error
}

func (e *statusError) Error() string {
	return e.wrapped.Error()
}

func (e *statusError) Unwrap() error {
	return e.wrapped
}

func SetKind(err error, kind ErrorKind) error {
	return &statusError{
		kind:    kind,
		wrapped: err,
	}
}

func GetKind(err error) ErrorKind {
	var se *statusError
	ok := errors.As(err, &se)
	if !ok {
		return KindUnknown
	}

	return se.kind
}
