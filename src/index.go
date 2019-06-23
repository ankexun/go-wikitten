package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//index页面
// Get传参数name,需要Parse的文件名
func Index(c *gin.Context) {
	var fileName string
	name := c.Query("name")
	if name == "" {
		fileName = "./" + Dir + "/" + Default
	} else {
		fileName = "./" + Dir + "/" + name
	}
	content, err := Parse(fileName)
	if err != nil {
		log.Printf("解析文件出错: %v", err)
	}

	//对树状目录排序,注意全局变量Data
	tree := sortData()

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":             Title,
		"USE_DARK_THEME":    UseDarkTheme,
		"USE_WIKITTEN_LOGO": UseWikittenLogo,
		"content":           content,
		"tree":              tree,
	})
}

// 测试一些前端代码用
func Test(c *gin.Context) {
	c.HTML(http.StatusOK, "test.tmpl", gin.H{
		"title": "test",
	})
}
