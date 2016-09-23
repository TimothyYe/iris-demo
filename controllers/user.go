package controllers

import (
	"fmt"

	"github.com/kataras/iris"
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

func (uc *UserController) GetUserProfile(ctx *iris.Context) {
	profile := Profile{Name: "隔壁老王", Age: 45}
	ctx.JSON(iris.StatusOK, profile)
}

func (uc *UserController) HandleLogin(ctx *iris.Context) {
	name := string(ctx.FormValue("name"))
	pwd := string(ctx.FormValue("password"))

	fmt.Println("Name:" + name)
	fmt.Println("Pwd:" + pwd)

	if name == "wang" && pwd == "123456" {
		resp := LoginResponse{Token: "sdakfjdsakdvmwiehfg==", Code: "00", Msg: "登录成功"}
		ctx.JSON(iris.StatusOK, resp)
		return
	}

	resp := LoginResponse{Token: "", Code: "01", Msg: "用户名或者密码不正确"}
	ctx.JSON(iris.StatusOK, resp)
}
