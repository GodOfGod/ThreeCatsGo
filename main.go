package main

import (
	api "ThreeCatsGo/api"
	middleware "ThreeCatsGo/middle_ware"
	"ThreeCatsGo/tools"
	"fmt"

	"github.com/gin-gonic/gin"

	dbCon "ThreeCatsGo/database"

	flag "github.com/spf13/pflag"
)

func main() {
	var env *string = flag.String("env", "dev", "current env, dev or prod")

	router := gin.Default()

	router.Use(middleware.VerifyToken())

	router.Static("/assets", "save_assets")

	db := dbCon.ConnectDB(*env)

	api.HandleRouter(router, db)

	flag.Parse()
	if *env == "dev" {
		fmt.Println(tools.ColoredStr("Server is running in dev environment").Red())
		router.Run("localhost:8080")
	} else if *env == "prod" {
		fmt.Println(tools.ColoredStr("Server is running in prod environment").Red())
		router.RunTLS("172.16.2.91:443", "cert.pem", "cert.key")
	} else {
		panic(tools.ColoredStr("wrong env arg, you should assign dev or prod to the flag env").Red())
	}
}
