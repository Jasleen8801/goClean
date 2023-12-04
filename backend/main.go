package main

import (
	"goClean/backend/models"
	"goClean/backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	router := gin.Default()
	router.Use(cors.Default())
	routes.StudentRoutes(router)
	routes.AdminRoutes(router)
	routes.HostelRoutes(router)

	router.Run(":3000")
}
