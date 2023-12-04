package routes

import (
	"goClean/backend/handlers"
	"goClean/backend/middlewares"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(r *gin.Engine) {
	student := r.Group("/student")
	{
		student.POST("/register", handlers.RegisterStudent) // public
		student.POST("/login", handlers.LoginStudent)       // public
	}

	student.Use(middlewares.AuthMiddleware())
	{
		student.GET("/getByID", handlers.GetStudentById)           // private
		student.GET("/get/:room_no", handlers.GetStudentsbyRoomID) // private
		student.PUT("/update/:room_no", handlers.UpdateRoomStatus) // private
		student.POST("/addLog", handlers.AddLog)                   // private
	}
}
