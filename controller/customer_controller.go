package controller

import (
	"go-boiler-plate/config"
	"go-boiler-plate/model"
	"go-boiler-plate/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserServiceImpl
	Conf    *config.AppConfig
}

func NewUserController(s *service.UserServiceImpl, config *config.AppConfig) *UserController {
	return &UserController{
		service: s,
		Conf:    config,
	}
}

func (u *UserController) SignUp(c *gin.Context) {
	input := model.UserSignUpRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to bind input data",
		})
		return
	}
	resp, err := u.service.SignUp(c, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to signup",
		})
		return
	}
	c.JSON(http.StatusCreated, resp)

}

func (u *UserController) SignIn(c *gin.Context) {
	input := model.UserSignInRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to bind input data",
		})
		return
	}
	resp, err := u.service.SignIn(c, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to signup",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (u *UserController) GetAllUser(c *gin.Context) {
	resp, err := u.service.GetAllUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get all users",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
