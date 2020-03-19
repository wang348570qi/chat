package main

import (
	"fmt"
	"chat/client/process"
	"os"
)

//定义变量，用户id，用户密码,用户名

var userId int
var userPwd string
var userName string

func main() {
	//接受用户输入的key
	var key int
	for true {
		fmt.Println("-------欢迎登陆多人聊天-------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择（1-3）")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入密码")
			fmt.Scanf("%s\n", &userPwd)
			up := &process.UserProcess{}
			up.Login(userId, userPwd)

		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名")
			fmt.Scanf("%s\n", &userName)
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
			//loop = false
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("您的输入有误，请重新输入")
		}

	}

}
