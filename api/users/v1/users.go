package v1

import "github.com/gogf/gf/v2/frame/g"

type RegisterReq struct {
	g.Meta   `path:"users1/register1"  method:"get" summary:"Register"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
type RegisterRes struct {
}

type DelUserReq struct {
	g.Meta `path:"users1/deluser"  method:"get" summary:"删除用户"`
	Id     int    `json:"id" v:"required#ID不能为空"`
	Token  string `json:"token" v:"required#Token不能为空"`
}
type DelUserRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
type LoginReq struct {
	g.Meta   `path:"/login" method:"get" summary:"用户登录"`
	Username string `json:"username" v:"required#用户名不能为空"`
	Password string `json:"password" v:"required#密码不能为空"`
}

type LoginRes struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type LogoutReq struct {
	g.Meta   `path:"/logout" method:"get" summary:"用户登出"`
	Username string `json:"username" v:"required#用户名不能为空"`
}

type LogoutRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type GetUserWithInfoReq struct {
	g.Meta `path:"/user/detail" method:"get" summary:"获取用户详细信息"`
	UserId uint `json:"userId" v:"required#用户ID不能为空"`
}

type GetUserWithInfoRes struct {
	User     interface{} `json:"user"`
	UserInfo interface{} `json:"userInfo"`
}
