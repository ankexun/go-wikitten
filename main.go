package main

import (
	"log"
	"net/http"

	"github.com/fsnotify/fsnotify"
	"github.com/go-ini/ini"
)

//全局变量
var (
	Config                        *ini.File
	Title, Default, Dir           string
	UseDarkTheme, UseWikittenLogo bool
	JsonData                      interface{}
)

func main() {
	// 读取config.ini的配置
	var err error
	Config, err = ini.Load("config.ini")
	if err != nil {
		log.Fatalf("读取config.ini配置文件错误: %v", err)
	}

	host := Config.Section("").Key("host").MustString("127.0.0.1")
	port := Config.Section("").Key("port").MustString("8080")
	Title = Config.Section("").Key("app_name").String()
	Default = Config.Section("").Key("default_file").MustString("index.md")
	Dir = Config.Section("").Key("library").MustString("myDoc")
	UseDarkTheme = Config.Section("").Key("use_dark_theme").MustBool(false)
	UseWikittenLogo = Config.Section("").Key("use_wikitten_logo").MustBool(false)

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

	JsonData = w.watchDir(Dir)

	select {}
}
