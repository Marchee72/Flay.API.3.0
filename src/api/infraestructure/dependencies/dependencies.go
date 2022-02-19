package dependencies

import (
	"fmt"

	database "flay-api-v3.0/src/api/config"
	"flay-api-v3.0/src/api/infraestructure/entrypoints"
	"flay-api-v3.0/src/api/infraestructure/entrypoints/handlers"
)

type HandlerContainer struct {
	Login entrypoints.Handler
}

func Start() *HandlerContainer {
	err := database.Connect()
	if err != nil {
		panic(fmt.Sprintf("ERROR. Cannot connect to database: (%s)", err.Error()))
	}

	apiHandlers := HandlerContainer{}

	apiHandlers.Login = &handlers.Login{}

	return &apiHandlers
}
