package api

import (
	"station-service/config"
	"station-service/db"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	config config.Config
	store  *db.Store
	router *gin.Engine
}

func NewServer(config config.Config, store *db.Store) (*Server, error) {

	gin.SetMode(config.GinMode)
	router := gin.Default()

	server := &Server{
		config: config,
		store:  store,
	}

	// Setup routing for server.
	v1 := router.Group("v1")
	{
		v1.GET("/stations/:id", server.GetByID)
		v1.GET("/stations", server.GetAll)
		v1.POST("/stations", server.Create)
		v1.PUT("/stations/:id", server.Update)
		v1.DELETE("/stations/:id", server.Delete)
	}

	// Setup health check routes.
	health := router.Group("health")
	{
		health.GET("/live", server.Live)
		health.GET("/ready", server.Ready)
	}

	// Setup metrics routes.
	metrics := router.Group("metrics")
	{
		metrics.GET("/", func(ctx *gin.Context) {
			handler := promhttp.Handler()
			handler.ServeHTTP(ctx.Writer, ctx.Request)
		})
	}

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
