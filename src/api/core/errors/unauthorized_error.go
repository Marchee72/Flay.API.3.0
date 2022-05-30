package errors

type UnauthorizedError struct {
	message string
	cause   []string
}

func NewUnauthorizedError(message string) UnauthorizedError {
	return UnauthorizedError{
		message: message,
	}
}

func NewUnauthorizedErrorWithCause(message string, cause ...string) UnauthorizedError {
	return UnauthorizedError{
		message: message,
		cause:   cause,
	}
}

func (e UnauthorizedError) Error() string {
	return e.message
}

func (e UnauthorizedError) Cause() []string {
	return e.cause
}
