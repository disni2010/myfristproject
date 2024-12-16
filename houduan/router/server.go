package main

import (
	"Go_Work/demo/CreateDemo/ginvuedemo1/houduan/model"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	model.InitRedis()
	router := gin.Default()
	//读取路径里的html
	//在cmd里运行脚本必须使用相对路径 在vscode运行可以使用绝对路径
	//router.LoadHTMLGlob("../template/**/*")
	//我的解决办法就是通过filepath.Walk来搜索/templates下的以.html结尾的文件，把这些html文件都加载一个数组中，然后用LoadHTMLFiles加载
	var files []string
	filepath.Walk("../template", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return nil
	})
	router.LoadHTMLFiles(files...)

	//加载资源文件css js 图片等资源文件 前面参数指定路径 后面参数相对路径
	router.Static("/static", "../static")
	server(router)
	adminServer(router)
	router.Run(":8080")
}
