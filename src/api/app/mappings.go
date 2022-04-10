package app

import (
	"flay-api-v3.0/src/api/infraestructure/dependencies"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func configureMappings(router *gin.Engine, handlers *dependencies.HandlerContainer) {
	api := router.Group("/api")
	api.POST("login", handlers.Login.Handle)
	api.POST("book", handlers.BookCommonSpace.Handle)
}
