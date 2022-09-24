package customErrors

type FileIsBrokenError struct {
	*CustomError
}

func CreateFileIsBrokenError(message string) *FileIsBrokenError {
	return &FileIsBrokenError{&CustomError{errorMessage: message}}
}
