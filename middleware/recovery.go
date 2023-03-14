package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yangkaiyue/gin-exp/global"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func(ctx *gin.Context) {
			if err := recover(); err != nil {
				global.Logger.Errorf("Server Error. Err: %v", err)
				resp := &global.ResponseJson{
					Head: global.Head{
						Msg: "Server Error",
					},
				}
				resp.ServerFail(ctx)
			}
		}(c)

		c.Next()
	}
}
