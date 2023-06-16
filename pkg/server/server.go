package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rookout/piper/pkg/clients"
	"github.com/rookout/piper/pkg/conf"
	"github.com/rookout/piper/pkg/server/routes"
)

var router = gin.Default()

func Start(cfg *conf.Config, clients *clients.Clients) {
	getRoutes(cfg, clients)
	router.Run()
}

func getRoutes(cfg *conf.Config, clients *clients.Clients) {
	v1 := router.Group("/")
	routes.AddHealthRoutes(cfg, v1)
	routes.AddWebhookRoutes(cfg, clients, v1)
}