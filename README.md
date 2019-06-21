# go-wikitten
go-Wikitten 是用golang写的，没有数据库，没有插件的高仿wikitten的#、支持Markdown语法的Wiki知识管理系统

### 部署
目前只做了windows x64的发布包,包里的文件如下:
/myDoc  			//这是存放你的md文件的地方，目录名在config.ini里定义
/static					//存放前端js和css的地方,不可以修改
config.ini				//配置文件
index.tmpl			//模版文件
tree.tmpl			//模版文件
wikitten.exe		//解压后,执行这个即可

下载wikitten.zip的压缩包,解压到一个目录下,双击wikitten.exe,会打开一个命令行窗口,不要关闭这个命令行窗口!! 打开浏览器,地址栏输入http://localhost:8080/ 就可以看到页面了.

如果需要修改端口什么的,修改config.ini,保存后重启wikitten.exe即可.

### 使用
windows下可以配合FTP使用.

