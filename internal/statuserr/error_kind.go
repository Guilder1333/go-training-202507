package statuserr

type ErrorKind int

const (
	KindUnknown ErrorKind = iota
	KindInvalidRequest
	KindUserNotFound
	KindCreateUserFailed
)
