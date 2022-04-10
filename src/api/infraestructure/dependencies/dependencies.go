package dependencies

import (
	"fmt"

	database "flay-api-v3.0/src/api/config"
	login "flay-api-v3.0/src/api/core/usecases/login"

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

	loginRepository := store.LoginRepository{
		Users: dbContainer.Users,
	}

	//Usecase injection
	loginUseCase := &login.Implementation{
		LoginRepository: &loginRepository,
	}
	//Handlers injection
	apiHandlers := HandlerContainer{}

	apiHandlers.Login = &handlers.Login{
		LoginUseCase: loginUseCase,
	}
	apiHandlers.BookCommonSpace = &handlers.BookCommonSpace{}

	return &apiHandlers
}
