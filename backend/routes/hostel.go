package routes

import (
	"goClean/backend/handlers"
	"goClean/backend/middlewares"

	"github.com/gin-gonic/gin"
)

func HostelRoutes(r *gin.Engine) {
	hostel := r.Group("/hostel")
	{
		hostel.POST("/login", handlers.LoginHostel) // public
	}

	hostel.Use(middlewares.AuthMiddleware())
	{
		hostel.PUT("/clear", handlers.ClearAllLogs) // private
	}
}
