package errors

import (
	"flay-api-v3.0/src/api/infraestructure/api_errors"
)

func GetCommonsApiError(err error) *api_errors.APIError {
	switch err.(type) {
	case BadRequestError:
		return newBadRequestApiError(err)
	default:
		return newInternalServerError(err)
	}

}

func newBadRequestApiError(err error) *api_errors.APIError {
	apiError := api_errors.NewBadRequestError(err.Error())
	apiError.Cause = err.(BadRequestError).Cause()
	return apiError
}

func newInternalServerError(err error) *api_errors.APIError {
	apiError := api_errors.NewInternalServerError(err.Error())
	apiError.Cause = err.(InternalServerError).Cause()
	return apiError
}
