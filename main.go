package main

import (
	api "ThreeCatsGo/api"
	middleware "ThreeCatsGo/middle_ware"
	"ThreeCatsGo/tools"

	"github.com/gin-gonic/gin"

	dbCon "ThreeCatsGo/database"
)

func main() {
	tools.CreateFolder()
	// create a router
	router := gin.Default()
	router.Use(middleware.VerifyToken())
	router.Static("/assets", "save_assets")
	// pass the router to router handler
	db := dbCon.ConnectDB()
	api.HandleRouter(router, db)
	// monitor the default port
	router.Run("localhost:8080")
}
