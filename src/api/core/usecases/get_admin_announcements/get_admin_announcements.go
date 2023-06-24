package get_admin_announcements

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/get_admin_announcements"
	"flay-api-v3.0/src/api/core/providers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UseCase interface {
	Execute(ctx context.Context, adminId primitive.ObjectID) (*get_admin_announcements.Response, error)
}

type Implementation struct {
	AnnouncementRepositorty providers.AnnouncementRepository
	BuildingRepository      providers.BuildingRepository
}

func (usecase Implementation) Execute(ctx context.Context, adminId primitive.ObjectID) (*get_admin_announcements.Response, error) {
	buildings, err := usecase.BuildingRepository.GetBuildingsByAdmin(ctx, adminId)
	if err != nil {
		return nil, err
	}
	response := get_admin_announcements.Response{}
	for _, building := range buildings {
		ann, err := usecase.AnnouncementRepositorty.GetBuildingAnnouncements(ctx, building.ID, nil, nil)
		if err != nil {
			return nil, err
		}
		for _, a := range ann {
			response.AddAnnouncement(a)
		}
	}
	return &response, nil

}
