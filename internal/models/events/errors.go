package events

import "fmt"

type ProcessingError struct {
	Message   string
	NeedRetry bool
	Cause     error
}

func NewProcessingError(msg string, needRetry bool, cause error) *ProcessingError {
	return &ProcessingError{
		Message:   msg,
		NeedRetry: needRetry,
		Cause:     cause,
	}
}

func (e *ProcessingError) Error() string {
	needRetry := "false"
	if e.NeedRetry {
		needRetry = "true"
	}

	if e.Cause != nil {
		return fmt.Sprintf("%s, need retry: %s, caused by: %v", e.Message, needRetry, e.Cause)
	}
	return fmt.Sprintf("%s, need retry: %s", e.Message, needRetry)
}

func (e *ProcessingError) Unwrap() error {
	return e.Cause
}
