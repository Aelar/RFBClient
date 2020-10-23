package core

//CustomError Standard error
type CustomError struct {
	msg string
}

// Create a function Error() string and associate it to the struct.
func (error *CustomError) Error() string {
	return error.msg
}

//MissingPwdError Send  when no RFB password  is provided but yet needed
type MissingPwdError struct {
	msg  string
	code int
}

func (error *MissingPwdError) Error() string {
	return error.msg
}

//FormatError for wrong arguments
type FormatError struct {
	msg  string
	code int
}

func (error *FormatError) Error() string {
	return error.msg
}
