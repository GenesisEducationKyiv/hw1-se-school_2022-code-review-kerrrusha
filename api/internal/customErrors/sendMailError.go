package customErrors

type SendMailError struct {
	*CustomError
}

func CreateSendMailError(message string) *SendMailError {
	return &SendMailError{&CustomError{errorMessage: message}}
}
