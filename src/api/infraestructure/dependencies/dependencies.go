package dependencies

import (
	"fmt"

	database "flay-api-v3.0/src/api/config"
	"flay-api-v3.0/src/api/core/usecases/book_common_space"
	"flay-api-v3.0/src/api/core/usecases/get_building_announcements"
	"flay-api-v3.0/src/api/core/usecases/get_building_bookings"
	"flay-api-v3.0/src/api/core/usecases/get_user_basic_info"
	"flay-api-v3.0/src/api/core/usecases/get_user_bookings"
	"flay-api-v3.0/src/api/core/usecases/get_user_building"
	login "flay-api-v3.0/src/api/core/usecases/login"
	"flay-api-v3.0/src/api/core/usecases/save_announcement"

	"flay-api-v3.0/src/api/infraestructure/entrypoints/handlers"
	store "flay-api-v3.0/src/api/infraestructure/repositories"
)

func Start() *HandlerContainer {
	//DB configs
	db, err := database.Connect()
	if err != nil {
		panic(fmt.Sprintf("ERROR. Cannot connect to database: (%s)", err.Error()))
	}

	//DB injection
	dbContainer := NewCollectionContainer(db)

	userRepository := store.UserRepository{
		Users: dbContainer.Users,
	}

	bookingRepository := store.BookingRepository{
		Bookings: dbContainer.Bookings,
	}
	// issueRepository := store.IssueRepository{
	// 	Issues: dbContainer.Issues,
	// }
	penaltyRepository := store.PenaltyRepository{
		Penalties: dbContainer.Penalties,
	}

	buildingRepository := store.BuildingRepository{
		Buildings: dbContainer.Buildings,
	}

	announcementRepository := store.AnnouncementRepository{
		Announcements: dbContainer.Announcements,
	}

	apartmentRepository := store.ApartmentRepository{
		Apartments: dbContainer.Apartments,
	}

	//Usecase injection
	loginUseCase := &login.Implementation{
		UserRepository: &userRepository,
	}

	bookCommonSpaceUseCase := &book_common_space.Implementation{
		PenaltyRepository: &penaltyRepository,
		BookingRepository: &bookingRepository,
		UserRepository:    &userRepository,
	}

	getUserBookingsUseCase := get_user_bookings.Implementation{
		BookingRepository: &bookingRepository,
	}

	getUserBuildingUseCase := get_user_building.Implementation{
		BuildingRepository: &buildingRepository,
		UserRepository:     &userRepository,
	}

	getBuildingBookingUseCase := get_building_bookings.Implementation{
		BookingRepository: &bookingRepository,
	}

	saveAnnouncementUseCase := save_announcement.Implementation{
		AnnouncementRepository: &announcementRepository,
		UserRepository:         &userRepository,
		BuildingRepository:     &buildingRepository,
	}

	getBuildingAnnouncementsUseCase := get_building_announcements.Implementation{
		AnnouncementRepository: &announcementRepository,
	}

	getUserBasicInfoUseCase := get_user_basic_info.Implementation{
		BuildingRepository:  &buildingRepository,
		UserRepository:      &userRepository,
		ApartmentRepository: &apartmentRepository,
	}
	//Handlers injection
	apiHandlers := HandlerContainer{}

	apiHandlers.Login = &handlers.Login{
		LoginUseCase: loginUseCase,
	}

	apiHandlers.BookCommonSpace = &handlers.BookCommonSpace{
		BookCommonSpaceUseCase: bookCommonSpaceUseCase,
	}

	apiHandlers.GetUserBookings = &handlers.GetUserBookings{
		GetUserBookingsUseCase: getUserBookingsUseCase,
	}

	apiHandlers.GetUserBuilding = &handlers.GetUserBuilding{
		GetUserBuildingUseCase: getUserBuildingUseCase,
	}

	apiHandlers.GetBuildingBookings = &handlers.GetBuildingBookings{
		GetBuildingBookingsUseCase: getBuildingBookingUseCase,
	}

	apiHandlers.SaveAnnouncement = &handlers.SaveAnnouncement{
		SaveAnnouncementUseCase: saveAnnouncementUseCase,
	}

	apiHandlers.GetBuildingAnnouncements = &handlers.GetBuildingAnnouncements{
		GetBuildingAnnouncementsUseCase: getBuildingAnnouncementsUseCase,
	}

	return &apiHandlers
}
