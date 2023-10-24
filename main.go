package main

import (
	"github.com/AwespireTech/InterfaceForCare-Backend/config"
	"github.com/AwespireTech/InterfaceForCare-Backend/database"
	"github.com/AwespireTech/InterfaceForCare-Backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.PrintConfig()
	err := database.Init(config.DATABASE_URL)
	if err != nil {
		panic(err)
	}

	router := createRouter()
	router.Run()
}

func createRouter() *gin.Engine {
	router := gin.Default()
	river := router.Group("api")
	{
		routes.SetRiverRoutes(river)
		routes.SetUserRoutes(river)
		routes.SetEventRoutes(river)
	}
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to InterfaceForCare API",
		})
	})
	return router
}
