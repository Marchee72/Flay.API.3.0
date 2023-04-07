package errors

import (
	"flay-api-v3.0/src/api/infraestructure/api_errors"
)

func GetCommonsApiError(err error) *api_errors.APIError {
	switch err.(type) {
	case BadRequestError:
		return newBadRequestApiError(err)
	case UnauthorizedError:
		return newUnauthorizedError(err)
	case ResourceNotFoundError:
		return newResourceNotFoundError(err)
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

func newUnauthorizedError(err error) *api_errors.APIError {
	apiError := api_errors.NewUnauthorizedError(err.Error())
	apiError.Cause = err.(UnauthorizedError).Cause()
	return apiError
}

func newResourceNotFoundError(err error) *api_errors.APIError {
	apiError := api_errors.NewResourceNotFound(err.Error())
	apiError.Cause = err.(ResourceNotFoundError).Cause()
	return apiError
}
