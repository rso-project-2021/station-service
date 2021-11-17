package server

import (
	"station-service/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(ginMode string) *gin.Engine {

	router := gin.Default()
	gin.SetMode(ginMode)

	// Expose router paths.
	v1 := router.Group("v1")
	{
		station := new(controllers.StationController)
		v1.GET("/stations/:id", station.GetByID)
		v1.GET("/stations", station.GetAll)
		v1.POST("/stations", station.Create)
		v1.PUT("/stations/:id", station.Update)
		v1.DELETE("/stations/:id", station.Delete)
	}

	return router
}
