package main
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"project1/filereader"
)

type Func func(w http.ResponseWriter,r *http.Request)

func main() {
	fmt.Println("正在初始化...")
	a := make(map[string]Func)
	apilist := filereader.New()
	for _, test := range apilist {
		fmt.Printf("正在创建接口 => %v\n", test.Api)
		a[test.Api] = ApiAdder(test.Response)
	}
	for k,v := range a{
		http.HandleFunc(k, v)
	}
	fmt.Println("初始化完毕...")
	fmt.Println("服务已起动... 请访问 http://localhost:1234")
	http.ListenAndServe(":1234", nil)
}

type Data struct {
	Key string
	Value string
}

type T struct {
	Url  string      `json:"url"`
	Func interface{} `json:"func"`
}

type Ret struct {
	Code int
	Params string
	Msg string
	Data []interface{}
}
func ApiAdder(m filereader.Response) Func{
	return func(w http.ResponseWriter, r *http.Request) {

		ret := new(Ret)
		ret.Code = m.Code
		ret.Params = m.Params
		ret.Msg = m.Msg
		for _, datum := range m.Data {
			ret.Data = append(ret.Data,datum)
		}
		ret_json,_ := json.Marshal(ret)
		io.WriteString(w, string(ret_json))

	}
}
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{Key:"1",Value:"1"}
	ret := new(Ret)

	ret.Code = 0
	ret.Params = "1"
	ret.Msg = "success"

	ret.Data = append(ret.Data,data)
	ret.Data = append(ret.Data,data)
	ret.Data = append(ret.Data,data)
	ret_json,_ := json.Marshal(ret)
	io.WriteString(w, string(ret_json))
}
func HelloHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World222")
}