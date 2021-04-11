package respz

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
)

func Json(w http.ResponseWriter, resp interface{}, err error) {
	if err != nil {
		httpx.Error(w, err)
	} else {
		httpx.OkJson(w, NewOkResp(resp))
	}
}
