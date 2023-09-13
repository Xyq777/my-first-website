package main

import "os"
func main() {

		_, err := os.Open("user.txt")
		if err != nil {
			println("err：你为什么不输入账号密码！！ (＃｀д´)ﾉ")
		}

	}
