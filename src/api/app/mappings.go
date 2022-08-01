package app

import (
	"flay-api-v3.0/src/api/infraestructure/dependencies"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func configureMappings(router *gin.Engine, handlers *dependencies.HandlerContainer) {
	api := router.Group("/api")
	api.POST("login", handlers.Login.Handle)
	api.POST("booking", handlers.BookCommonSpace.Handle)
	api.GET("user/:user_id/bookings", handlers.GetUserBookings.Handle)
	api.GET("building/:building_id/bookings", handlers.GetBuildingBookings.Handle)
	api.GET("building", handlers.GetUserBuilding.Handle)

}
