package controllers

import (
	"os"

	"github.com/kataras/iris/v12"
)

func DefaultHandler(ctx iris.Context) {
	hostName, err := os.Hostname()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.Writef("Hostname is: %s\n", hostName)
}
