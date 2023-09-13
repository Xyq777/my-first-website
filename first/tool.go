package main

import "strings"

/*
	func generateHTML(writer http.ResponseWriter, data interface{}, filename ...string) {
		templates := template.Must(template.ParseFiles(filename...))
		err := templates.ExecuteTemplate(writer, "layout", templates)
		checker(err)
	}
*/
func register(username, password string) string {

	if username != "" && password != "" {

		var cunzai = false
		if strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") && len(password) >= 8 {

			_, cunzai = yonghu[username]
			if !cunzai {
				yonghu[username] = password //添加用户数据
			}
			if !cunzai {
				res := "注册成功"
				return res
			} else {
				res := "用户名已存在"
				return res
			}
		}
	}
	return ""
}

func denglu(username, password string) string {
	if username != "" && password != "" {

		var cunzai, pipei = false, false
		//检验用户是否存在
		_, cunzai = yonghu[username]
		//检查登录是否成功

		if password == yonghu[username] {
			pipei = true

		}

		if cunzai {
			if pipei {
				res := "登陆成功"
				return res

			} else {
				res := "密码不匹配"
				return res
			}
		} else {
			res := "用户名不存在"
			return res
		}
	}
	return ""
}

/*type data struct {
	cunzai bool
	pipei  bool
}
*/
