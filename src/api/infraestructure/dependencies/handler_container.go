package dependencies

import "flay-api-v3.0/src/api/infraestructure/entrypoints"

type HandlerContainer struct {
	Login                    entrypoints.Handler
	BookCommonSpace          entrypoints.Handler
	GetUserBookings          entrypoints.Handler
	GetUserBuilding          entrypoints.Handler
	GetBuildingBookings      entrypoints.Handler
	SaveAnnouncement         entrypoints.Handler
	GetAnnouncement          entrypoints.Handler
	GetBuildingAnnouncements entrypoints.Handler
	GetUserBasicInfo         entrypoints.Handler
	SaveExpenseFile          entrypoints.Handler
	GetUnitExpenses          entrypoints.Handler
	GetFile                  entrypoints.Handler
}
