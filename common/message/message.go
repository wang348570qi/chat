package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}
type LoginResMes struct {
	Code    int `json:"code"`
	UsersId []int
	Error   string `json:"error"`
}
type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type NotifyUserStatusMes struct {
	UserId int `json:"userId`
	Status int `json:"status"`
}
type SmsMes struct {
	Content string `json:"content"`
	User
}
