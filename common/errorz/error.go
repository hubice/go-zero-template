package errorz

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (c *CodeError) Error() string {
	return string(c.Code) + c.Msg
}

func (c *CodeError) RawData() CodeError {
	return *c
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}
