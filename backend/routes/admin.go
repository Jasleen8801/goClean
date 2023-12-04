package routes

import (
	"goClean/backend/handlers"
	"goClean/backend/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine) {
	admin := r.Group("/admin")
	{
		admin.POST("/register", handlers.RegisterAdmin) // public
		admin.POST("/login", handlers.LoginAdmin)       // public
	}

	admin.Use(middlewares.AuthMiddleware())
	{
		admin.POST("/addHostel", handlers.AddHostel) // private
	}
}
