package respz

type OkResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewOkResp(data interface{}) *OkResp {
	return &OkResp{
		Code: 0,
		Msg:  "ok",
		Data: data,
	}
}
