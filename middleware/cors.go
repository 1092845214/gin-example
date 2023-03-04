package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {

	cfg := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTION"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Token", "Authorization "},
		AllowCredentials: true,
		//AllowOrigins:     nil,
		//AllowAllOrigins:        true,  // 设置 AllowCredentials 后这个不生效
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 20 * 24 * 60 * 60,
	}
	return cors.New(cfg)
}
