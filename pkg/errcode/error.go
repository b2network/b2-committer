package errcode

type ErrCode struct {
	Code    string // business error code
	Message string // error message
	Err     error  //  real error
}

func NewErrCodeWithMessage(code string, msg string) ErrCode {
	err := ErrCode{
		Code:    code,
		Message: msg,
	}
	return err
}

func (e ErrCode) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}
