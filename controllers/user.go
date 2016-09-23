package controllers

import "github.com/kataras/iris"

type UserController struct {
}

type Profile struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (uc *UserController) GetUserProfile(ctx *iris.Context) {
	profile := Profile{Name: "隔壁老王", Age: 45}
	ctx.JSON(iris.StatusOK, profile)
}
