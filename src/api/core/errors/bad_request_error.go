package errors

type BadRequestError struct {
	message string
	cause   []string
}

func NewBadRequestError(message string) BadRequestError {
	return BadRequestError{
		message: message,
	}
}

func NewBadRequestErrorWithCause(message string, cause ...string) BadRequestError {
	return BadRequestError{
		message: message,
		cause:   cause,
	}
}

func (e BadRequestError) Error() string {
	return e.message
}

func (e BadRequestError) Cause() []string {
	return e.cause
}
