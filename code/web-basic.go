package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func index(rw http.ResponseWriter, req *http.Request) {

	req.ParseForm()       // 解析参数
	fmt.Println(req.Form) // 输出到服务器端的打印信息
	fmt.Println("path", req.URL.Path)
	fmt.Println("scheme", req.URL.Scheme)
	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintln(rw, "Hello, Astaxie") // 写入到的 rw 是输出到客户端的
}

func login(w http.ResponseWriter, req *http.Request) {
	fmt.Println("method:", req.Method) //获取请求的方法
	if req.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
        if(len(req.Form["username"][0]) == 0) {
            // 数据校验 
        }
		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", req.Form["username"])
		fmt.Println("password:", req.Form["password"])
	}
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/index", index)         // 设置访问路由
	err := http.ListenAndServe(":9090", nil) // 设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
