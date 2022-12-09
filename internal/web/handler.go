package web

import (
	"github.com/gin-gonic/gin"
	"mk/internal/services"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}
	api := router.Group("/notes", h.userIdentity)
	{
		api.GET("/")
		api.GET("/:id")
		api.POST("/")
	}
	return router
}
