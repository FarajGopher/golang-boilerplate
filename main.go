package main

import (
	"golang-boilerplate/config"
	"golang-boilerplate/database"
	"golang-boilerplate/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	
	config.InitEnvConfig()
	database.ConnectDB() 

	

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.SetTrustedProxies(nil)
	r.Run(":" + config.EnvConfigs.ServerPort)

}