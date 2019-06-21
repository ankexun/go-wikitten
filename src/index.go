package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var fileName string
	name := c.Query("name")
	if name == "" {
		fileName = "./" + Dir + "/" + Default
	} else {
		fileName = "./" + Dir + "/" + name
	}
	content, lang, err := Parse(fileName)
	if err != nil {
		log.Printf("解析文件出错: %v", err)
	}
	tree := sortData()
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":             Title,
		"USE_DARK_THEME":    UseDarkTheme,
		"USE_WIKITTEN_LOGO": UseWikittenLogo,
		"content":           content,
		"tree":              tree,
		"lang":              lang,
	})
}

func Test(c *gin.Context) {
	c.HTML(http.StatusOK, "test.tmpl", gin.H{
		"title": "test",
	})
}
