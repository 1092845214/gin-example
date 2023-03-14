package router

import (
	"context"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/yangkaiyue/gin-exp/docs"
	"github.com/yangkaiyue/gin-exp/global"
	"github.com/yangkaiyue/gin-exp/middleware"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type IFnRegisRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	routes []IFnRegisRoute
)

// appendRoute 各子级路由调用该方法将路由添加到 routes 中
func appendRoute(fn IFnRegisRoute) {

	if fn == nil {
		return
	}
	routes = append(routes, fn)
}

// InitRouter 初始化路由
func InitRouter() {

	// 优雅关闭
	// 参考链接 https://github.com/gin-gonic/examples/tree/master/graceful-shutdown/graceful-shutdown
	// gin 官方提供的使用 context 的方式优雅退出, 有两种方式, 这里使用 with-context 的方案

	// 打开一个 channel 用于接收两个退出信号
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// 自定义 engine, 加入中间件
	r := gin.New()
	r.Use(middleware.Cors(), middleware.Recovery(), middleware.GinLogger())

	// 系统路由
	initSysRoutes(r)

	// 两个组, 一个用于鉴权的组, 一个用于非鉴权的组
	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")

	// 添加并注册路由(新加路由在这里添加)
	initBusinessRoutes()
	for i := 0; i < len(routes); i++ {
		routes[i](rgPublic, rgAuth)
	}

	// 端口设置默认值
	viper.SetDefault("server.port", 9000)
	port := viper.GetString("server.port")

	// 启动(gin框架提供的启动方式, 无法支持强制退出)
	//r.Run(viper.GetString("server.ip") + ":" + port)

	// 为了保证非优雅退出时, 有强制退出兜底, 所以要先获取 serve , 由于 gin 包装的方法无法获取, 所以自己创造
	server := &http.Server{
		Addr:    viper.GetString("server.ip") + ":" + port,
		Handler: r,
	}
	// 启动一个 goroutine 监听
	go func() {
		global.Logger.Infof("Open Listening '%v:%v'", viper.GetString("server.ip"), viper.GetString("server.port"))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Fatalf("Start Server Failed. Err: %v", err.Error())
			return
		}
	}()

	// 阻塞
	<-ctx.Done()

	global.Logger.Infoln("Exit Gin Server")

	// 设置超时时间,超时则强制退出(防止优雅退出失败)
	ctx, stopForce := context.WithTimeout(context.Background(), 5*time.Second)
	defer stopForce()
	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Fatalf("Stop Server Failed. Err: %v", err.Error())
		return
	}
}

// initSysRoutes
func initSysRoutes(r *gin.Engine) {

	// 注册 gin-swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 非 release 模式下, 开启 pprof 资源监控
	if viper.GetBool("server.debug") {
		pprof.Register(r)
	}

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.Header("Content-type", "application/json")
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// 处理路由错误
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  fmt.Sprintf("No Route: %v", c.Request.URL.Path),
		})
	})

	// 处理方法错误, 需要开启这个才有 NoMethod 的处理
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code": 405,
			"msg":  fmt.Sprintf("No Method: %v", c.Request.Method),
		})
	})
}

// initBusinessRoutes 基础路由
func initBusinessRoutes() {

	// 基础路由 --> 用户相关路由
	initUserRoutes()
}
