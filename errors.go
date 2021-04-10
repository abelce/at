package at

import "abelce/at/errors"

// BusinessError business error
type BusinessError interface {
	error
	// Code error code
	Code() string
}

type businessError struct {
	code    string
	message string
}

func (e *businessError) Code() string {
	return e.code
}

func (e *businessError) Error() string {
	return e.message
}

func NewError(code string, message string) BusinessError {
	return &businessError{
		code:    code,
		message: message,
	}
}

func Ensure(e *error, messages ...string) bool {
	return errors.EnsureWithCaller(e, 1, messages...)
}

func Extract(e *error) bool {
	err := *e
	if atError, ok := err.(errors.Error); ok {
		cause := atError.Cause()
		*e = cause
		return true
	}
	return false
}

//get the original error
func RawError(err error) error {
	if err == nil {
		return err
	}

	if atError, ok := err.(errors.Error); ok {
		return atError.Cause()
	}

	return err
}
