package handlers

import (
	"context"
	"log"
	"net/http"

	"flay-api-v3.0/src/api/core/constants"
	contracts "flay-api-v3.0/src/api/core/contracts/get_file"
	"flay-api-v3.0/src/api/core/errors"
	"flay-api-v3.0/src/api/core/usecases/get_file"
	"flay-api-v3.0/src/api/infraestructure/api_errors"
	"flay-api-v3.0/src/api/infraestructure/entrypoints/wrappers"
	"github.com/gin-gonic/gin"
)

type GetFile struct {
	GetFileUseCase get_file.UseCase
}

func (handler *GetFile) Handle(c *gin.Context) {
	wrappers.AuthWrapper(handler.handle, c, []constants.UserType{constants.UserRenter, constants.UserAdmin})
}

func (handler *GetFile) handle(c *gin.Context) *api_errors.APIError {
	ctx := context.Background()
	request := contracts.Request{}
	err := c.ShouldBindUri(&request)
	if err != nil {
		log.Printf("Error binding request: %s", err.Error())
		return errors.GetCommonsApiError(err)
	}
	response, err := handler.GetFileUseCase.Execute(ctx, request.ContentID())
	if err != nil {
		log.Printf("Error getting file: %s", err.Error())
		return errors.GetCommonsApiError(err)
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=myfile.txt")

	// Write the file content to the response
	c.Data(http.StatusOK, "application/octet-stream", response)
	return nil
}
