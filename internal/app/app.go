package app

import (
	"fmt"
	"net/http"

	"github.com/MrOldOwl/react_server/internal/app/configuration"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func Run() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		api.GET("/setting", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"setting": "ok",
			})
		})
	}

	json_config, _ := configuration.OpenFile("./configuration.json")
	router.Run(fmt.Sprintf(":%v", json_config.Port))
}
