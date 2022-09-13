package model

type CustomError struct {
	errorMessage string
}

func (err *CustomError) GetMessage() string {
	return err.errorMessage
}
