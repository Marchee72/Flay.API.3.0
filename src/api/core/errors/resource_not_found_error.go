package errors

type ResourceNotFoundError struct {
	message string
	cause   []string
}

func NewResourceNotFoundError(message string) ResourceNotFoundError {
	return ResourceNotFoundError{
		message: message,
	}
}

func NewResourceNotFoundErrorWithCause(message string, cause ...string) ResourceNotFoundError {
	return ResourceNotFoundError{
		message: message,
		cause:   cause,
	}
}

func (e ResourceNotFoundError) Error() string {
	return e.message
}

func (e ResourceNotFoundError) Cause() []string {
	return e.cause
}
