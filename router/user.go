package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yangkaiyue/gin-exp/api"
	"net/http"
)

// initUserRoutes 初始化用户相关路由
func initUserRoutes() {

	appendRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {

		// 用户路由子组(建不建都行)
		rgPublicUser := rgPublic.Group("/user")
		{
			// 登录路由
			var user api.User
			rgPublicUser.POST("/login/:user", user.Login)
		}

		rgAuthUser := rgAuth.Group("/user")
		{
			rgAuthUser.GET("/list", func(ctx *gin.Context) {
				ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
					"data": []map[string]any{
						{"id": "张三"},
						{"id": "李四"},
						{"id": "王五"},
					},
				})
			})

			rgAuthUser.GET("/:id", func(ctx *gin.Context) {
				ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
					"data": map[string]any{
						"id":   1,
						"name": "Aki",
					},
				})
			})
		}
	})
}
