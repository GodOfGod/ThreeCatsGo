package main

import (
	api "ThreeCatsGo/api"
	middleware "ThreeCatsGo/middle_ware"
	"ThreeCatsGo/tools"
	"fmt"

	"github.com/gin-gonic/gin"

	dbCon "ThreeCatsGo/database"

	flag "github.com/spf13/pflag"

	globalvar "ThreeCatsGo/global_var"
)

func main() {
	var env *string = flag.String("env", "dev", "current env, dev or prod")
	flag.Parse()
	globalvar.SetEnv(*env)
	router := gin.Default()

	// 静态资源目录前缀为 /h5，资源存储在 ./static
	staticDir := "./static"
	// 添加中间件处理 .gz 文件
	router.Use(middleware.GzipServer(staticDir))
	router.Static("/h5", staticDir)

	router.Use(middleware.VerifyToken())
	router.Static("/assets", "save_assets")
	router.Static("/favicon.ico", "./static/favicon.ico")

	db := dbCon.ConnectDB(*env)

	router.GET("/questionnaire/v1", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	api.HandleRouter(router, db)

	if *env == "dev" {
		fmt.Println(tools.ColoredStr("Server is running in dev environment").Red())
		router.Run("localhost:8080")
	} else if *env == "prod" {
		fmt.Println(tools.ColoredStr("Server is running in prod environment").Red())
		router.RunTLS("172.16.2.91:443", "cert/cert.pem", "cert/cert.key")
	} else {
		panic(tools.ColoredStr("wrong env arg, you should assign dev or prod to the flag env").Red())
	}
}
