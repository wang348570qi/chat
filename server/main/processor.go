package main

import (
	"fmt"
	"io"
	"chat/common/message"
	"chat/server/process"
	"chat/server/utils"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	fmt.Println("mes=", mes)
	switch mes.Type {
	case message.LoginMesType:
		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		smsProcess := &process.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在,无法处理...")
	}
	return
}

func (this *Processor) process() (err error) {

	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出,服务端也退出")
				return err
			} else {
				fmt.Println("readPkg err=", err)
				return err
			}
		}
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
