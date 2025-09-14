package controller

import (
	"golang-boilerplate/dto"
	"golang-boilerplate/repository"
	"golang-boilerplate/service"
	"golang-boilerplate/utils"

	"github.com/gin-gonic/gin"
)

// Controller interface
type AuthController interface {
	Signup(c *gin.Context)
}

// Implementation
type authController struct {
	service service.AuthService
}

func NewAuthController() AuthController {
	repo := repository.NewAuthRepository()
	svc := service.NewAuthService(repo)
	return &authController{service: svc}
}

func (ctrl *authController) Signup(c *gin.Context) {
	var input dto.SignupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	user, err := ctrl.service.Signup(input)
	if err != nil {
		utils.ErrorResponse(c, 500, err.Error())
		return
	}

	utils.SuccessResponse(c, 201, utils.USER_CREATED_SUCCESS, user)
}
