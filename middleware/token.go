package middleware

import "github.com/kataras/iris"

type TokenMiddleware struct{}

func (m *TokenMiddleware) Serve(c *iris.Context) {
	if c.RequestHeader("Authorization") == "sdakfjdsakdvmwiehfg==" {
		c.Next()
	} else {
		c.Text(403, "Please login first!!")
	}
}
