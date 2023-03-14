package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yangkaiyue/gin-exp/global"
	"go.uber.org/zap"
	"time"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now().In(global.CstZone)
		path := c.Request.URL.Path // /name/aki
		//query := c.Request.URL.RawQuery // ?name=aki

		c.Next()

		cost := time.Since(start)
		global.Logger.Infow(c.HandlerName(),
			zap.Int("status", c.Writer.Status()),
			zap.String("cost", cost.String()),
			zap.String("client_ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
		)
	}
}
