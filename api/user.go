package api

import (
	"github.com/gin-gonic/gin"
)

type User struct{}

// Login
// @Summary 用户登录接口
// @Description 传入用户名,密码进行登录操作
// @Tags USER
// @Accept json
// @Product json
// @Param user path string true "用户名"
// @Success 200 {} string "登录成功"
// @Failure 400 {} string "登录失败"
// @Router /public/user/login/{user} [POST]
func (m *User) Login(ctx *gin.Context) {
	resp := &ResponseJson{
		Msg: "Login Success",
	}
	resp.Success(ctx)
}
