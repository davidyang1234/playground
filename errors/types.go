package srserr

import "fmt"

type baseError struct {
	ErrCode        string `json:"error"`
	ErrMessage     string `json:"error_description"`
	status         int
	origErr        error
	OrigErrMessage string `json:"original_error"`
}

func newBaseError(status int, code, message string, origErr error) *baseError {
	var msg = ""
	if origErr != nil {
		msg = origErr.Error()
	}
	return &baseError{
		ErrCode:        code,
		ErrMessage:     message,
		status:         status,
		origErr:        origErr,
		OrigErrMessage: msg,
	}
}

func (e *baseError) Error() string {
	if e.origErr == nil {
		return fmt.Sprintf("%s: %s", e.ErrCode, e.ErrMessage)
	}
	return fmt.Sprintf("%s: %s (%s)", e.ErrCode, e.ErrMessage, e.OrigErrMessage)
}

func (e *baseError) Status() int {
	return e.status
}

func (e *baseError) Code() string {
	return e.ErrCode
}

func (e *baseError) Message() string {
	return e.ErrMessage
}

func (e *baseError) OrigErr() error {
	return e.origErr
}

func (e *baseError) WithMessage(msg string) Error {
	return &baseError{
		e.ErrCode,
		msg,
		e.status,
		e.origErr,
		e.OrigErrMessage,
	}
}

func (e *baseError) WithError(err error) Error {
	return &baseError{
		e.ErrCode,
		e.ErrMessage,
		e.status,
		err,
		err.Error(),
	}
}
