package errorz

import (
	"net/http"
)

func ErrorHandler(err error) (i int, i2 interface{}) {
	switch e := err.(type) {
	case *CodeError:
		return http.StatusOK, e.RawData()
	default:
		return http.StatusOK, CodeError{
			Code: 9999,
			Msg:  e.Error(),
		}
	}
}
