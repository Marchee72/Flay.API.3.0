package get_building_announcements

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/get_building_announcements"
	"flay-api-v3.0/src/api/core/providers"
)

type UseCase interface {
	Execute(ctx context.Context, request get_building_announcements.Request) (*get_building_announcements.Response, error)
}

type Implementation struct {
	AnnouncementRepository providers.AnnouncementRepository
}

func (usecase Implementation) Execute(ctx context.Context, request get_building_announcements.Request) (*get_building_announcements.Response, error) {
	announcements, err := usecase.AnnouncementRepository.GetBuildingAnnouncements(ctx, request.BuildingID(), request.Date, request.Date)
	if err != nil {
		return nil, err
	}
	response := get_building_announcements.Response{}
	for _, a := range announcements {
		response.AddAnnouncement(a)
	}
	return &response, nil

}
