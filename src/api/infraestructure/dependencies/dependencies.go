package dependencies

import (
	"flay-api-v3.0/src/api/infraestructure/entrypoints"
	"flay-api-v3.0/src/api/infraestructure/entrypoints/handlers"
)

type HandlerContainer struct {
	Login entrypoints.Handler
}

func Start() *HandlerContainer {
	apiHandlers := HandlerContainer{}

	apiHandlers.Login = &handlers.Login{}

	return &apiHandlers
}
