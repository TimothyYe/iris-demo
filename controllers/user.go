package controllers

import (
	"github.com/kataras/iris/v12"
)

type UserController struct {
}

type Profile struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Code  string `json:"code"`
	Msg   string `json:"msg"`
}

func (uc *UserController) GetUserProfile(ctx iris.Context) {
	profile := Profile{Name: "隔壁老王", Age: 45}
	ctx.JSON(profile)
}

func (uc *UserController) HandleLogin(ctx iris.Context) {
	name := string(ctx.FormValue("name"))
	pwd := string(ctx.FormValue("password"))

	ctx.Application().Logger().Debugf("Name:" + name)
	ctx.Application().Logger().Debugf("Pwd:" + pwd)

	if name == "wang" && pwd == "123456" {
		resp := LoginResponse{Token: "sdakfjdsakdvmwiehfg==", Code: "00", Msg: "登录成功"}
		ctx.JSON(resp)
		return
	}

	resp := LoginResponse{Token: "", Code: "01", Msg: "用户名或者密码不正确"}
	ctx.JSON(resp)
}
