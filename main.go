package main

import "github.com/yangkaiyue/gin-exp/cmd"

// @title  gin example
// @version v0.0.1
// @description gin 框架示例
// @BasePath /api/v1
func main() {

	// 最后收尾清理工作
	defer cmd.Clean()

	// 开始函数
	cmd.Start()
}
