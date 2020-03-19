package process

import (
	"encoding/json"
	"fmt"
	"chat/common/message"
	"chat/server/model"
	"chat/server/utils"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

func (this *UserProcess) NotifyOtherOnlineUser(userId int) {
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		up.NotifyMeOnline(userId)
	}
}
func (this *UserProcess) NotifyMeOnline(userId int) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.marshal err=", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.marshal err=", err)
		return
	}
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("notifymeonline err=", err)
		return
	}

}
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.unmarshal fail err=", err)
		return
	}
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var LoginResMes message.LoginResMes

	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			LoginResMes.Code = 500
			LoginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			LoginResMes.Code = 403
			LoginResMes.Error = err.Error()
		} else {
			LoginResMes.Code = 505
			LoginResMes.Error = "服务器内部错误..."
		}
	} else {
		LoginResMes.Code = 200
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		this.NotifyOtherOnlineUser(loginMes.UserId)
		for id, _ := range userMgr.onlineUsers {
			LoginResMes.UsersId = append(LoginResMes.UsersId, id)
		}
		fmt.Println(user, "登录成功")
	}
	/* if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		LoginResMes.Code = 200
	} else {
		LoginResMes.Code = 500
		LoginResMes.Error = "该用户不存在，请注册..."
	} */

	data, err := json.Marshal(LoginResMes)
	if err != nil {
		fmt.Println("json.marshal fail", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.marshal fail", err)
		return
	}
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.unmarshal fail err=", err)
		return
	}
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 200
	}
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.marshal fail", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.marshal fail ", err)
		return
	}
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
