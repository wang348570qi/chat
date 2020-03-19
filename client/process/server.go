package process

import (
	"encoding/json"
	"fmt"
	"chat/client/utils"
	"chat/common/message"
	"net"
	"os"
)

func ShowMenu() {

	fmt.Println("-------恭喜登录成功------")
	fmt.Println("\t\t\t 1 显示在线用户列表")
	fmt.Println("\t\t\t 2 发送消息")
	fmt.Println("\t\t\t 3 信息列表")
	fmt.Println("\t\t\t 4 退出系统")
	fmt.Println("\t\t\t 请选择（1-4）")
	var key int
	var content string
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", key)
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		fmt.Println("你想对大家说什么:")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("您的输入有误，请重新输入")
	}
}

func serverProcessMes(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.Readpkg err=", err)
			return
		}
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType: //群发消息
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器端返回了未知消息类型")
		}

		//fmt.Printf("mes=%v\n", mes)
	}
}
