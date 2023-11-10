package dependencies

import (
	"fmt"

	database "flay-api-v3.0/src/api/config"
	"flay-api-v3.0/src/api/core/usecases/book_common_space"
	"flay-api-v3.0/src/api/core/usecases/get_admin_announcements"
	"flay-api-v3.0/src/api/core/usecases/get_announcement"
	"flay-api-v3.0/src/api/core/usecases/get_building_announcements"
	"flay-api-v3.0/src/api/core/usecases/get_building_bookings"
	"flay-api-v3.0/src/api/core/usecases/get_building_expenses"
	"flay-api-v3.0/src/api/core/usecases/get_buildings"
	"flay-api-v3.0/src/api/core/usecases/get_file"
	"flay-api-v3.0/src/api/core/usecases/get_unit_expenses"
	"flay-api-v3.0/src/api/core/usecases/get_user_basic_info"
	"flay-api-v3.0/src/api/core/usecases/get_user_bookings"
	"flay-api-v3.0/src/api/core/usecases/get_user_building"
	login "flay-api-v3.0/src/api/core/usecases/login"
	"flay-api-v3.0/src/api/core/usecases/save_announcement"
	"flay-api-v3.0/src/api/core/usecases/save_expense_file"

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

	fileRepository := store.FileRepository{
		Files: dbContainer.Files,
	}
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

	expenseRepository := store.ExpenseRepository{
		Expenses: dbContainer.Expenses,
		Files:    &fileRepository,
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
		UserRepository:    &userRepository,
	}

	getUserBuildingUseCase := get_user_building.Implementation{
		BuildingRepository: &buildingRepository,
		UserRepository:     &userRepository,
	}

	getBuildingBookingUseCase := get_building_bookings.Implementation{
		BookingRepository: &bookingRepository,
		UserRepository:    &userRepository,
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

	getAnnouncementUseCase := get_announcement.Implementation{
		AnnouncementRepository: &announcementRepository,
	}

	saveExpenseFileUseCase := save_expense_file.Implementation{
		ExpenseRepository: &expenseRepository,
	}

	getUnitExpenses := get_unit_expenses.Implementation{
		ExpensesRepository: &expenseRepository,
	}

	getFile := get_file.Implementation{
		FileRepository: &fileRepository,
	}

	getAdminAnnouncements := get_admin_announcements.Implementation{
		AnnouncementRepositorty: &announcementRepository,
		BuildingRepository:      &buildingRepository,
	}

	getBuildings := get_buildings.Implementation{
		BuildingRepository: &buildingRepository,
	}

	getBuildingExpenses := get_building_expenses.Implementation{
		ExpenseRepository: &expenseRepository,
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

	apiHandlers.GetUserBasicInfo = &handlers.GetUserBasicInfo{
		GetUserBasicInfoUseCase: getUserBasicInfoUseCase,
	}

	apiHandlers.GetAnnouncement = &handlers.GetAnnouncement{
		GetAnnouncementUseCase: getAnnouncementUseCase,
	}

	apiHandlers.SaveExpenseFile = &handlers.SaveExpenseFile{
		SaveExpenseFileUseCase: saveExpenseFileUseCase,
	}

	apiHandlers.GetUnitExpenses = &handlers.GetUnitExpenses{
		GetUnitExpensesUseCase: getUnitExpenses,
	}
	apiHandlers.GetFile = &handlers.GetFile{
		GetFileUseCase: getFile,
	}

	apiHandlers.GetAdminAnnouncements = &handlers.GetAdminAnnouncements{
		GetAdminAnnouncementUseCase: getAdminAnnouncements,
	}

	apiHandlers.GetBuildings = &handlers.GetBuildings{
		GetBuildingsUseCase: getBuildings,
	}

	apiHandlers.GetBuildingExpenses = &handlers.GetBuildingExpenses{
		GetBuildingExpensesUseCase: &getBuildingExpenses,
	}

	return &apiHandlers
}
