package main

import (
	"log"
	"net/http"

	"github.com/fsnotify/fsnotify"
	"github.com/lxmgo/config"
)

//全局变量
var Title, Default, Dir string
var UseDarkTheme, UseWikittenLogo bool

func main() {
	// 读取config.ini的配置
	config, err := config.NewConfig("config.ini")
	if err != nil {
		log.Printf("读取config.ini配置文件错误: %v", err)
	}
	host := config.String("host")
	port := config.String("port")
	Title = config.String("app_name")
	Default = config.String("default_file")
	Dir = config.String("library")
	UseDarkTheme, _ = config.Bool("use_dark_theme")
	UseWikittenLogo, _ = config.Bool("use_wikitten_logo")

	//设置缺省目录与首页
	if Dir == "" {
		Dir = "myDoc"
	}
	if Default == "" {
		Default = "index.md"
	}

	router := InitRouter()
	srv := &http.Server{
		Addr:    host + ":" + port,
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	watch, _ := fsnotify.NewWatcher()

	w := Watch{
		watch: watch,
	}

	if Dir == "" {
		w.watchDir("myDoc")
	} else {
		w.watchDir(Dir)
	}

	select {}
}
