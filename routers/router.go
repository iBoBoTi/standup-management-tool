package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/standup-management-tool/handlers"
	"github.com/iBoBoTi/standup-management-tool/repository"
	"github.com/iBoBoTi/standup-management-tool/server"
	"github.com/iBoBoTi/standup-management-tool/service"
)

const (
	EnvironmentProduction  = "production"
	EnvironmentDevelopment = "development"
	EnvironmentTesting     = "testing"
)

// SetupRouter registers all the HTTP routes in the system
// if you want to move out some routes, you can accept *gin.Engine as an argument
func SetupRouter(srv *server.Server) {
	if srv.GetConfig().Environment == EnvironmentDevelopment ||
		srv.GetConfig().Environment == EnvironmentTesting {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(server.CustomLogger(srv.Logger), gin.Recovery()).Use(server.CORS())

	v1 := router.Group("/api/v1")

	v1.GET("/health-check", handlers.NewHealthController(srv).HealthCheck)

	authHandler := handlers.NewAuthHandler(srv, service.NewAuthService(srv.TokenMaker, repository.NewEmployeeRepository(srv.DB.GormDB)))
	v1.POST("/signup", authHandler.SignUp)
	v1.POST("/login", authHandler.Login)

	v1.Use(srv.ApplyAuthentication())
	employeeHandler := handlers.NewEmployeeHandler(srv, service.NewEmployeeService(repository.NewEmployeeRepository(srv.DB.GormDB)))
	v1.POST("/employees", employeeHandler.AdminCreateEmployee)

	srv.Router = router

}
