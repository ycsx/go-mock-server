package apiadder

import (
	"encoding/json"
	"io"
	"net/http"
	"project1/filereader"
)

type Func func(w http.ResponseWriter, r *http.Request)

type Ret struct {
	Code   int           `json:"code"`
	Params string        `json:"params"`
	Msg    string        `json:"msg"`
	Data   []interface{} `json:"data"`
}

func ApiAdder(m filereader.Response) Func {
	return func(w http.ResponseWriter, r *http.Request) {
		ret := new(Ret)
		ret.Code = m.Code
		ret.Params = m.Params
		ret.Msg = m.Msg
		for _, datum := range m.Data {
			ret.Data = append(ret.Data, datum)
		}
		ret_json, _ := json.Marshal(ret)
		io.WriteString(w, string(ret_json))
	}
}
