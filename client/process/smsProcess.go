package process

import (
	"encoding/json"
	"fmt"
	"chat/client/utils"
	"chat/common/message"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("sendGroupMES JSON.marshal fail =", err.Error())
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("sendGroupmes json.marshal fail", err.Error())
		return
	}
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("sendGroupmes err=", err.Error)
		return
	}
	return

}
