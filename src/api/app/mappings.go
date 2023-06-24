package app

import (
	"flay-api-v3.0/src/api/infraestructure/dependencies"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func configureMappings(router *gin.Engine, handlers *dependencies.HandlerContainer) {
	api := router.Group("/api")
	api.POST("login", handlers.Login.Handle)
	api.GET("user/basic-info", handlers.GetUserBasicInfo.Handle)
	api.POST("booking", handlers.BookCommonSpace.Handle) //TODO: change url to building/:building_id/booking
	api.GET("user/:user_id/bookings", handlers.GetUserBookings.Handle)
	api.GET("building", handlers.GetUserBuilding.Handle)
	api.GET("building/:building_id/bookings", handlers.GetBuildingBookings.Handle)
	api.POST("building/:building_id/announcement", handlers.SaveAnnouncement.Handle)
	api.GET("building/:building_id/announcements", handlers.GetBuildingAnnouncements.Handle)
	api.GET("announcements", handlers.GetAdminAnnouncements.Handle)
	api.GET("announcements/:announcement_id", handlers.GetAnnouncement.Handle)
	api.POST("expense", handlers.SaveExpenseFile.Handle)
	api.GET("expense/:unit", handlers.GetUnitExpenses.Handle)
	api.GET("file/:content_id", handlers.GetFile.Handle)
}
