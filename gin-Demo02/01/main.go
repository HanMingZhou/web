package main

import (
	"fmt"
	"html/template"
	"net/http"

	_"github.com/gin-gonic/gin"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	//2. 解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse failed err:%v",err)
		return
	}
	//3. 渲染模板
	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("render template err :%v",err)
		return
	}
}
func main() {
	http.HandleFunc("/",sayhello)
	err := http.ListenAndServe(":9000",nil)
	if err != nil {
		fmt.Printf("http start failed err : %v",err)
		return
	}
}