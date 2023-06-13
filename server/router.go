package server

import (
	"go-boiler-plate/config"
	"go-boiler-plate/controller"
	"go-boiler-plate/db"
	"go-boiler-plate/middleware"
	"go-boiler-plate/service"

	"github.com/gin-gonic/gin"
)

func (*HTTPServer) setUpRoutes(router *gin.Engine, dbStore db.Store, config *config.AppConfig) {
	userController := controller.NewUserController(&service.UserServiceImpl{Store: dbStore, Conf: config}, config)
	userLoginSignUp(router, *userController, config)
	userInfo(router, *userController, config)

}

func userLoginSignUp(router *gin.Engine, userController controller.UserController, c *config.AppConfig) {
	p := router.Group("/v1")
	p.POST("/signup", userController.SignUp)
	p.POST("/login", userController.SignIn)
}

func userInfo(router *gin.Engine, userController controller.UserController, c *config.AppConfig) {
	p := router.Group("/v1", middleware.JwtAuthMiddleware())
	p.GET("/all_users", userController.GetAllUser)
}
