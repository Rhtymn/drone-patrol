package apperror

type AppError struct {
	msg  string
	code int
}

func (e AppError) Error() string {
	return e.msg
}

func NewAppError(msg string, code int) AppError {
	return AppError{
		msg:  msg,
		code: code,
	}
}

func NewInvalidArguments(msg string) AppError {
	return AppError{
		msg:  msg,
		code: InvalidArguments,
	}
}

func NewInternal(msg string) AppError {
	return AppError{
		msg:  msg,
		code: Internal,
	}
}

func NewOutOfRangePosition(msg string) AppError {
	return AppError{
		msg:  msg,
		code: OutOfRangePosition,
	}
}

func ErrorIs(err error, code int) bool {
	e, ok := err.(AppError)
	if !ok {
		return false
	}

	return e.code == code
}
