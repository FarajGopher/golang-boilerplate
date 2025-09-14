package routes

import (
	"golang-boilerplate/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Group routes under /api
	api := r.Group("/api")
	authController := controller.NewAuthController()
	api.POST("/signup", authController.Signup)
}
