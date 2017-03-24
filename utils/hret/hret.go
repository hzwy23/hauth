package hret

import (
	"encoding/json"

	"net/http"
	"strconv"

	"github.com/hzwy23/asofdate/utils/logs"
)

type HttpOkMsg struct {
	Version    string      `json:"version"`
	Reply_code int         `json:"reply_code"`
	Reply_msg  string      `json:"reply_msg"`
	Data       interface{} `json:"data,omitempty"`
	Total      int64       `json:"total,omitempty"`
	Rows       interface{} `json:"rows,omitempty"`
}

type HttpErrMsg struct {
	Error_code    int         `json:"error_code"`
	Error_msg     string      `json:"error_msg"`
	Error_details interface{} `json:"error_details,omitempty"`
	Version       string      `json:"version"`
}

func NewHttpErrMsg(code int, msg string, details interface{}) HttpErrMsg {
	return HttpErrMsg{
		Error_code:    code,
		Error_msg:     msg,
		Error_details: details,
		Version:       "v1.0",
	}
}

func WriteHttpErrMsg(w http.ResponseWriter, herr HttpErrMsg) {
	herr.Version = "v1.0"
	ijs, err := json.Marshal(herr)
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(`{error_code:` + strconv.Itoa(herr.Error_code) + `,error_msg:"` + herr.Error_msg + `",error_details:"format json type info failed."}`))
		return
	}
	w.WriteHeader(herr.Error_code)
	w.Write(ijs)
}

func WriteJson(w http.ResponseWriter, data interface{}) {
	ijs, err := json.Marshal(data)
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(`{error_code:765,error_msg:"` + err.Error() + `",error_details:"format json type info failed."}`))
		return
	}
	if string(ijs) == "null" {
		w.Write([]byte("[]"))
		return
	}
	w.Write(ijs)
}

func WriteHttpErrMsgs(w http.ResponseWriter, code int, msg string, details ...interface{}) {
	e := HttpErrMsg{
		Error_code:    code,
		Error_msg:     msg,
		Error_details: details,
	}

	WriteHttpErrMsg(w, e)
}

func WriteHttpOkMsg(w http.ResponseWriter, ok HttpOkMsg) {
	ojs, err := json.Marshal(ok)
	if err != nil {
		logs.Error(err.Error())
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(`{error_code:` + strconv.Itoa(http.StatusExpectationFailed) + `,error_msg:"format json type info failed.",error_details:"format json type info failed."}`))
		return
	}
	w.Write(ojs)
	return
}

func WriteHttpOkMsgs(w http.ResponseWriter, v interface{}) {
	ok := HttpOkMsg{
		Version:    "v1.0",
		Reply_code: 200,
		Reply_msg:  "execute successfully.",
		Data:       v,
	}
	ojs, err := json.Marshal(ok)
	if err != nil {
		logs.Error(err.Error())
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(`{error_code:` + strconv.Itoa(http.StatusExpectationFailed) + `,error_msg:"format json type info failed.",error_details:"format json type info failed."}`))
		return
	}
	w.Write(ojs)
	return
}

func WriteBootstrapTableJson(w http.ResponseWriter, total int64, v interface{}) {
	ok := HttpOkMsg{
		Version:    "v1.0",
		Reply_code: 200,
		Reply_msg:  "execute successfully.",
		Rows:       v,
		Total:      total,
	}

	ijs, err := json.Marshal(ok)
	if err != nil {
		logs.Error(err.Error())
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(`{error_code:` + strconv.Itoa(http.StatusExpectationFailed) + `,error_msg:"format json type info failed.",error_details:"format json type info failed."}`))
		return
	}
	w.Write(ijs)
}

type HttpPanicFunc func()

// HttpPanic user for stop panic up.
func HttpPanic(f ...HttpPanicFunc) {
	if r := recover(); r != nil {
		for _, val := range f {
			val()
		}
	}
}
