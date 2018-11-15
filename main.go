package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"text/template"
)

/*User 模板输出*/
type User struct {
	Username string
	Gender   string
	Fruit    string
}

/*GetGenderByNum 模板会调用的函数*/
func GetGenderByNum(num string) string {
	if num == "1" {
		return string("male")
	} else if num == "2" {
		return string("female")
	} else {
		return string("unknown gender")
	}
}

func unknown(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "501 unknown") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()
		if len(r.Form["username"][0]) == 0 {
			//为空的处理
			fmt.Println("username cannot be empty")
		}
		slice := []string{"apple", "pear", "banana"}

		v := r.Form.Get("fruit")
		for _, item := range slice {
			if item == v {
				fmt.Println("fruit", item)
			}
		}
		slice = []string{"1", "2"}

		for _, v := range slice {
			if v == r.Form.Get("gender") {
				fmt.Println("gender get", reflect.TypeOf(v))
			}
		}
		fmt.Println("r", r)
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		user := User{Username: r.Form["username"][0],
			Fruit:  r.Form.Get("fruit"),
			Gender: r.Form.Get("gender")}

		t := template.New("foo")
		t = t.Funcs(template.FuncMap{"gender": GetGenderByNum})
		t, _ = t.Parse(`{{define "T"}}
		<table border="1">
		<tr>
			<th>username</th>
			<th>gender</th>
			<th>favourite fruit</th>
		</tr>
		<tr>
			<td>{{.Username}}</td>
			<td>{{.Gender|gender}}</td>
			<td>{{.Fruit}}</td>
		</tr>
	</table> 
			{{end}}
			`)
		err := t.ExecuteTemplate(w, "T", user)
		if err != nil {
			fmt.Println("err", err)
		}

	}
}

func main() {
	wd, err := os.Getwd()
	fmt.Println("wd", wd)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", unknown)    //设置访问的路由
	http.HandleFunc("/login", login) //设置访问的路由
	fsh := http.FileServer(http.Dir(wd))
	http.Handle("/static", http.StripPrefix("/static", fsh))

	err = http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
