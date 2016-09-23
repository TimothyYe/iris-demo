package controllers

import (
	"os"

	"github.com/kataras/iris"
)

type DefaultController struct {
}

func (dc *DefaultController) DefaultHandler(ctx *iris.Context) {
	hostName, _ := os.Hostname()
	ctx.Text(iris.StatusOK, "Hostname is: "+hostName)
}
