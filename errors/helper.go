package errors

import (
	"strings"

	"github.com/abelce/at/errors/internal"
	"github.com/abelce/at/logging"
)

// Ensure wrap a raw error and trace
func Ensure(pErr *error, messages ...string) bool {
	return EnsureWithCaller(pErr, 1, messages...)
}

// EnsureWithCaller wrap a raw error and trace
func EnsureWithCaller(pErr *error, callerSkip int, messages ...string) bool {
	if pErr == nil || *pErr == nil {
		return false
	}

	switch (*pErr).(type) {
	case Error: // Pass has handled err
		return true
	}

	msg := strings.Join(messages, ", ")
	vtErr := &internal.VTError{
		Message: msg,
		Raw:     *pErr,
	}
	if AllowStackTrace {
		vtErr.ST = internal.Callers(3 + callerSkip).StackTrace()
	}
	if AllowAutoLogging {
		content := vtErr.Error()
		if content != "" {
			content = content + ", "
		}
		logging.Tracef("%sStackframe: %v", content, vtErr.ST)
	}

	*pErr = vtErr
	return true
}
