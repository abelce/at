package internal

import "fmt"

type VTError struct {
	Message string
	Raw     error
	ST      StackTrace
}

func (e *VTError) Error() string {
	cause := e.Cause()
	msg := e.Message
	if cause != nil {
		causeErr := cause.Error()
		if msg == "" {
			msg = causeErr
		} else if causeErr != "" {
			msg = msg + ", " + causeErr
		}
	}
	return msg
}

// Format implement fmt.PrintX interface
func (e *VTError) Format(s fmt.State, verb rune) {
	if e.ST == nil {
		fmt.Fprintf(s, "%s", e.Error())
		return
	}
	e.ST.Format(s, verb)
}

func (e *VTError) StackTrace() StackTrace {
	type stackTracer interface {
		StackTrace() StackTrace
	}

	if e.ST != nil {
		return e.ST
	}

	if e.Raw != nil {
		err := e.Raw
		if st, ok := err.(stackTracer); ok {
			return st.StackTrace()
		}
	}

	return nil
}

func (e *VTError) Cause() error {
	type causer interface {
		Cause() error
	}

	err := e.Raw
	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}
