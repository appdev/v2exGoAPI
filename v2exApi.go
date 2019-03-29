package main

import (
	"os"
	"io"
	"github.com/gin-gonic/gin"
	"./routers"
)

func main() {
	engine := gin.Default()
	//设置日志
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stderr)
	//日志
	engine.Use(gin.Logger())
	//加载路由文件
	routers.LoadRouters(engine)
	//绑定端口
	engine.Run(":8080")
}

