package main

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//发布编译时设置release模式
	//gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	// 设置静态资源
	router.Static("static", "./static")
	router.Static(Dir, Dir)

	router.LoadHTMLGlob("*.tmpl")
	router.GET("/", Index)
	router.GET("/test", Test)

	// 未知路由处理
	router.NoRoute(func(context *gin.Context) {
		context.String(404, "Not router")
	})
	// 未知调用方式
	router.NoMethod(func(context *gin.Context) {
		context.String(404, "Not method")
	})

	return router
}
