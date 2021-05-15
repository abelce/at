package errors

import "github.com/abelce/at/errors/internal"

// Error varitrip error
type Error interface {
	error
	// Cause return raw error
	Cause() error
}

// New Make a new varitrip error
func New(message string) Error {
	return &internal.VTError{
		Message: message,
		ST:      internal.Callers(1).StackTrace(),
	}
}
