package apperror

// AppError representing custom app error which implement built-in Error interface
type AppError struct {
	msg  string
	code int
}

// Error returning error message
func (e AppError) Error() string {
	return e.msg
}

// NewAppError will creating app error which initialized with given msg and code value
func NewAppError(msg string, code int) AppError {
	return AppError{
		msg:  msg,
		code: code,
	}
}

// NewInvalidArguments will creating app error with code Invalid Arguments
func NewInvalidArguments(msg string) AppError {
	return AppError{
		msg:  msg,
		code: InvalidArguments,
	}
}

// NewInvalidArguments will creating app error with code Internal
func NewInternal(msg string) AppError {
	return AppError{
		msg:  msg,
		code: Internal,
	}
}

// NewInvalidArguments will creating app error with code Out Of Range Position
func NewOutOfRangePosition(msg string) AppError {
	return AppError{
		msg:  msg,
		code: OutOfRangePosition,
	}
}

// NewInvalidArguments will creating app error with code Insufficient Plots
func NewInsufficientPlots() AppError {
	return AppError{
		msg:  "insufficient plots",
		code: InsufficientPlots,
	}
}

// Wrap will do nothing if given app error
// it will transform any error to app error if given error which not apperror
// transformed error will have code internal
func Wrap(err error) AppError {
	e, ok := err.(AppError)
	if !ok {
		return NewInternal(e.Error())
	}
	return e
}

// ErrorIs will checking is given error match with given code
// If it's match, it will return true. Otherwise will return false
func ErrorIs(err error, code int) bool {
	e, ok := err.(AppError)
	if !ok {
		return false
	}

	return e.code == code
}
