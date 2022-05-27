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
func f1(w http.ResponseWriter, r *http.Request) {

	//定义一个函数kua
	kua := func (name string) (string, error) {
		return name + "真帅啊", nil 
	}
	// 定义模板
	t := template.New("f.tmpl")// template.New("f") ——> 创建一个名字f的模板对象,名字一定要和模板的名字对应上
	t.Funcs(template.FuncMap{
		"kua" : kua,
	})

	// 解析模板 
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed err:%v",err)
	}

	// 渲染模板
	name := "小王子"
	t.Execute(w, name)
}

func demo1(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./t.tmpl","./ul.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed err:%v",err)
		return 
	}
	//渲染模板
	name := "小王子"
	t.Execute(w, name)
}
func main() {
	http.HandleFunc("/",f1)
	http.HandleFunc("/templateDemo",demo1)
	err := http.ListenAndServe(":9000",nil)
	if err != nil {
		fmt.Printf("http start failed err %v",err)
		return 
	}
}