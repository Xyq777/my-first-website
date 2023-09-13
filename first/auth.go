package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

var yonghu = make(map[string]string)

/*type data struct {
	username string `json:"username"`
	password string `json:"password"`
}*/

type Result struct {
	Msg string `json:"msg"`
}

func regReceive(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")
	err := request.ParseMultipartForm(1 << 10)
	checker(err)

	username := request.FormValue("username")
	password := request.FormValue("password")
	println("username=", username)
	if res := register(username, password); res != "" {

		result := Result{Msg: res}

		data, err := json.Marshal(result)
		if err != nil {
			println("编码失败")

		} else {
			_, err := writer.Write(data)

			checker(err)
		}

	}
}

func reg(writer http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/layout.html",
		"templates/reg.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	err := templates.ExecuteTemplate(writer, "layout", templates)
	checker(err)

}

//遍历数据库

/*
	func index(writer http.ResponseWriter, request *http.Request) {
		files := []string{
			"templates/layout.html",
			"templates/content.html",
		}
		templates := template.Must(template.ParseFiles(files...))
		err := templates.ExecuteTemplate(writer, "layout", templates)
		checker(err)

}
*/
func login(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json,text/html")
	files := []string{
		"templates/layout.html",
		"templates/login.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	err := templates.ExecuteTemplate(writer, "layout", templates)
	checker(err)

	err = request.ParseForm()
	checker(err)

	username := request.FormValue("name")
	password := request.FormValue("password")

	if res := denglu(username, password); res != "" {
		data, err := json.Marshal(res)
		if err != nil {
			println("编码失败")

		} else {
			_, err := writer.Write(data)
			checker(err)
		}

	}

}

func logReceive(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")
	err := request.ParseMultipartForm(1 << 10)
	checker(err)

	username := request.FormValue("username")
	password := request.FormValue("password")
	if res := denglu(username, password); res != "" {

		result := Result{Msg: res}

		data, err := json.Marshal(result)
		if err != nil {
			println("编码失败")

		} else {
			_, err := writer.Write(data)

			checker(err)
		}

	}
}
func checker(err error) {
	if err != nil {
		println(err)

	}
}
