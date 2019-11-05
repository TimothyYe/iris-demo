package middleware

import "github.com/kataras/iris/v12"

func TokenMiddleware(ctx iris.Context) {
	if ctx.GetHeader("Authorization") == "sdakfjdsakdvmwiehfg==" {
		ctx.Next()
		return
	}

	ctx.StatusCode(iris.StatusForbidden)
	ctx.WriteString("Please login first!!")
}
