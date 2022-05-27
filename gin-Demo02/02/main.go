package main

import (
	"fmt"
	"html/template"
	"net/http"
	
	
)

type User struct {
	Name string
	Gender string 
	Age int
}
func sayhello(w http.ResponseWriter, r *http.Request) {

	// 定义模板

	// 解析模板 
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse failed err:%v",err)
	}

	// 渲染模板

	u1 := User{
		Name:   "小王子",
		Gender: "男",
		Age:    10,
	}
	m1 := map[string]interface{} {
		"name":   "小王子~",
		"gender": "男~",
		"age":   10,
	}
	hobbylist := []string {
		"篮球",
		"足球",
		"双色球",
	}
	t.Execute(w, map[string]interface{} {
		"u1" : u1,
		"m1" : m1,
		"hobby" : hobbylist,

	})
}

func main() {
	http.HandleFunc("/",sayhello)
	err := http.ListenAndServe(":9000",nil)
	if err != nil {
		fmt.Printf("http start failed err %v",err)
		return 
	}
}