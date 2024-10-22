package main

import (
	api "ThreeCatsGo/api"
	middleware "ThreeCatsGo/middle_ware"
	"ThreeCatsGo/tools"

	"github.com/gin-gonic/gin"

	dbCon "ThreeCatsGo/database"

	flag "github.com/spf13/pflag"
)

func main() {
	router := gin.Default()

	router.Use(middleware.VerifyToken())

	router.Static("/assets", "save_assets")

	db := dbCon.ConnectDB()

	api.HandleRouter(router, db)

	var env *string = flag.String("env", "dev", "current env, dev or prod")
	flag.Parse()
	if *env == "dev" {
		router.Run("localhost:8080")
	} else if *env == "prod" {
		router.RunTLS("172.16.2.91:443", "cert.pem", "cert.key")
	} else {
		panic(tools.ColoredStr("wrong env arg, you should assign dev or prod to the flag env").Red())
	}
}
