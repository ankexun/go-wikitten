package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
*index页面
*显示config.ini指定的Default页面
 */
func Index(c *gin.Context) {
	fileName := "./" + Dir + "/" + Default
	content, err := Parse(fileName)

	if err != nil {
		log.Fatalf("解析文件出错: %v", err)
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":             Title,
		"USE_DARK_THEME":    UseDarkTheme,
		"USE_WIKITTEN_LOGO": UseWikittenLogo,
		"content":           content,
		"tree":              JsonData,
	})
}

/*
* POST
* @name,需要Parse的文件名
* @isDir,是否是目录
 */
func GetByName(c *gin.Context) {
	var fileName string
	name := c.Request.FormValue("name")
	isDir := c.Request.FormValue("isDir")
	// fmt.Printf("GetByName:%v\n", isDir)
	if isDir == "false" {
		if name == "" {
			fileName = "./" + Dir + "/" + Default
		} else {
			fileName = name
		}
		content, err := Parse(fileName)
		if err != nil {
			log.Fatalf("解析文件出错: %v", err)
		}

		c.JSON(http.StatusOK, gin.H{
			"content": content,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"content": "",
		})
	}
}

// 测试一些前端代码用
func Test(c *gin.Context) {
	//对树状目录排序,注意全局变量Data
	// tree, _ := json.Marshal(JsonData)
	c.HTML(http.StatusOK, "test.tmpl", gin.H{
		"title": "test",
		"tree":  JsonData,
	})
}
