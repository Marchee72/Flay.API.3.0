package dependencies

import "flay-api-v3.0/src/api/infraestructure/entrypoints"

type HandlerContainer struct {
	Login          entrypoints.Handler
	BookCommonArea entrypoints.Handler
}
