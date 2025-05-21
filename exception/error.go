package exception

import "fmt"

type CustomError struct {
	code    interface{}
	message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("[Error]: Code: %v, Message: %s", e.code, e.message)
}

func (e *CustomError) GetCode() interface{} {
	return e.code
}

func (e *CustomError) GetMessage() string {
	return e.message
}

func (e *CustomError) SetMessage(msg string, args ...interface{}) *CustomError {
	e.message = fmt.Sprintf(msg, args...)
	return e
}
func (e *CustomError) SetMessageArgs(args ...interface{}) *CustomError {
	return e.SetMessage(e.GetMessage(), args...)
}

func NewCustomError(code interface{}, err interface{}, args ...interface{}) *CustomError {
	var message string
	switch v := err.(type) {
	case string:
		message = v
	case error:
		message = v.Error()
	default:
		return UnknownError
	}
	return &CustomError{
		code:    code,
		message: fmt.Sprintf(message, args...),
	}
}
