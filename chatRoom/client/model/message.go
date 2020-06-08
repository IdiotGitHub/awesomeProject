package model

const LoginMesType = "LoginMes"
const LoginResultMesType = "LoginResultMes"
const RegisterMesType = "RegisterMes"
const RegisterResultMesType = "RegisterResultMes"

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息内容
}

type LoginMes struct {
	UserId  string `json:"user_id"`
	UserPwd string `json:"user_pwd"`
	//UserName string `json:"user_name"`
}

type LoginResultMes struct {
	Code    int      `json:"code"` //登录状态码 500未注册  200登录成功
	UsersId []string `json:"users_id"`
	Error   string   `json:"error"` //错误内容，
}
type RegisterMes struct {
	UserId   string `json:"user_id"`
	UserPwd  string `json:"user_pwd"`
	UserName string `json:"user_name"`
}

type RegisterResultMes struct {
	Code  int    `json:"code"`  //登录状态码 500未注册  200登录成功
	Error string `json:"error"` //错误内容，
}
