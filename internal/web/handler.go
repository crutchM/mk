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
		//не логично не логично
		auth.POST("/sign-in", h.signUp)
		auth.POST("/sign-up", h.signIn)
	}
	api := router.Group("/notes", h.userIdentity)
	{
		api.GET("/", h.getNotes)
		//а где метод у точки?
		api.GET("/:id")
		//а как мы через гет запрос запись создадим?
		api.GET("/", h.createNote)
	}
	return router
}
