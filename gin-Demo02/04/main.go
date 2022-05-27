package main

import (
	"fmt"
	"html/template"
	"net/http"	
)


func home (w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed err：%v",err)
		return
	}
	msg :=  "小王子"
	//渲染模板
	t.Execute(w,msg)
}
func index (w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed err：%v",err)
		return
	}
	msg := "小王子"
	//渲染模板
	t.Execute(w,msg)
}
func main() {
	http.HandleFunc("/index",index)
	http.HandleFunc("/home",home)
	err := http.ListenAndServe(":9000",nil)
	if err != nil {
		fmt.Printf("http start failed err %v",err)
		return 
	}
}