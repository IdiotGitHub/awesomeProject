package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
go-web处理http请求的时候可以直接使用内置函数http.HandleFunc(string, func)来进行处理，他会自动的将func转为一个处理器，
	或者使用自定义处理器，那么就要这个自定义处理器实现ServeHTTP(w http.ResponseWriter, r *http.ReadRequest){}这个接口
如果需要配置服务器属性，可以创建一个http.Server{}来进行详细配置，然后使用这个结构体直接调用ListenAndServe()方法来启动服务器
获取请求行:r.URL.Path
获取请求参数:r.URL.RawQuery
	r.Form也可以获取请求参数，包括url参数和form表单，所有的参数值都可以得到，form表单中的参数值要在url的前面。使用之前需要执行r.ParseForm()
	Form contains the parsed form data, including both the URL field's query parameters and the PATCH,
	POST, or PUT form data. This field is only available after ParseForm is called. The HTTP client ignores
	Form and uses Body instead
r.PostForm获取form表单中的参数，这个字段仅支持enctype='application/x-wwww-form-urlencoded'编码
r.Multipart获取表单enctype='multipart/form-data'的表单参数（上传文件时使用)
获取请求头:r.Header--类型是map[string][]string
获取请求体r.Body(只有post请求才有请求体)
	len := r.ContentLength//获取请求体长度
	body := make([]byte, len)//创建一个byte数组用于存放请求体内容
	string(body)

*/
type MyHandler struct {
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprint(w, "hello world by myself")
	if err != nil {
		return
	}
}
func main() {
	myHandler := &MyHandler{}
	http.Handle("/helloWorld", myHandler)
	http.HandleFunc("/hello", HelloWorld)
	server := http.Server{
		Addr:        ":8080",
		Handler:     myHandler,
		ReadTimeout: 2 * time.Second,
	}
	err := server.ListenAndServe()
	//err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
func HelloWorld(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprint(w, "hello world by system")
	if err != nil {
		return
	}
}
