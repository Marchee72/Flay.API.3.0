package errors

type InternalServerError struct {
	message string
	cause   []string
}

func NewInternalServerError(message string) InternalServerError {
	return InternalServerError{
		message: message,
	}
}

func NewInternalServerErrorWithCause(message string, cause ...string) InternalServerError {
	return InternalServerError{
		message: message,
		cause:   cause,
	}
}

func (e InternalServerError) Error() string {
	return e.message
}

func (e InternalServerError) Cause() []string {
	return e.cause
}
