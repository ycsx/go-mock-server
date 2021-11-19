package main

import (
	"fmt"
	"net/http"
	"project1/apiadder"
	"project1/filereader"
)

func main() {
	fmt.Println("正在初始化...")
	a := make(map[string]apiadder.Func)
	apilist := filereader.New()
	for _, test := range apilist {
		fmt.Printf("正在创建接口 => %v\n", test.Api)
		a[test.Api] = apiadder.ApiAdder(test.Response)
	}
	for k, v := range a {
		http.HandleFunc(k, v)
	}
	fmt.Println("初始化完毕...")
	fmt.Println("服务已起动... 请访问 http://localhost:1234")
	http.ListenAndServe(":1234", nil)
}
