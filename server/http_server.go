package server

import (
	"go-boiler-plate/config"
	"go-boiler-plate/db"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type HTTPServer struct {
	Router *gin.Engine
}

func Init(config *config.AppConfig, dbStore db.Store) (*HTTPServer, error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			nSubString := 2
			name := strings.SplitN(fld.Tag.Get("json"), ",", nSubString)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
	server := &HTTPServer{}
	gin.SetMode(config.GinMode)
	router := gin.Default()
	// Set the middleware for the health checks for k8s HTTP probes.
	// router.Use(middleware.HealthCheck())
	server.setUpRoutes(router, dbStore, config)
	server.Router = router
	return server, nil
}

func (server *HTTPServer) Start(address string) error {
	return server.Router.Run(address)

}
