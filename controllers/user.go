package controllers

import "github.com/kataras/iris"

type UserController struct {
}

type Profile struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	Result bool   `json:"result"`
	Msg    string `json:"msg"`
}

func (uc *UserController) GetUserProfile(ctx *iris.Context) {
	profile := Profile{Name: "隔壁老王", Age: 45}
	ctx.JSON(iris.StatusOK, profile)
}

func (uc *UserController) HandleLogin(ctx *iris.Context) {
	name := string(ctx.FormValue("name"))
	pwd := string(ctx.FormValue("password"))

	if name != "wang" && pwd != "123456" {
		resp := LoginResponse{Token: "", Result: false, Msg: "用户名或者密码不正确"}
		ctx.JSON(iris.StatusOK, resp)
	}

	resp := LoginResponse{Token: "sdakfjdsakdvmwiehfg==", Result: true, Msg: "登录成功"}
	ctx.JSON(iris.StatusOK, resp)
}
