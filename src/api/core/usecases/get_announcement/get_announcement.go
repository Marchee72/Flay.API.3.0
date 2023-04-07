package get_announcement

import (
	"context"
	"fmt"

	"flay-api-v3.0/src/api/core/contracts/common"
	"flay-api-v3.0/src/api/core/contracts/get_announcement"
	"flay-api-v3.0/src/api/core/errors"
	"flay-api-v3.0/src/api/core/providers"
)

type UseCase interface {
	Execute(ctx context.Context, request get_announcement.Request) (*get_announcement.Response, error)
}

type Implementation struct {
	AnnouncementRepository providers.AnnouncementRepository
}

func (useCase Implementation) Execute(ctx context.Context, request get_announcement.Request) (*get_announcement.Response, error) {
	announcement, err := useCase.AnnouncementRepository.GetAnnouncement(ctx, request.BuildingID())
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	if announcement == nil {
		return nil, errors.NewResourceNotFoundError(fmt.Sprintf("Announcement with ID: %s not found", request.AnnouncementRawID))
	}
	response := get_announcement.Response{
		User:     common.NewUserLw(announcement.User),
		Building: common.NewBuildingLw(announcement.Building),
		Title:    announcement.Title,
		Message:  announcement.Message,
		Date:     announcement.Date,
		Severity: announcement.Severity,
	}
	return &response, nil
}
