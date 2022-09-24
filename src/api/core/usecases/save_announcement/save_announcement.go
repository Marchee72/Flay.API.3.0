package save_announcement

import (
	"context"
	"fmt"

	"flay-api-v3.0/src/api/core/contracts/save_announcement"
	"flay-api-v3.0/src/api/core/entities"
	"flay-api-v3.0/src/api/core/errors"
	"flay-api-v3.0/src/api/core/providers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UseCase interface {
	Execute(ctx context.Context, request save_announcement.Request, userID primitive.ObjectID) error
}
type Implementation struct {
	AnnouncementRepository providers.AnnouncementRepository
	UserRepository         providers.UserRepository
	BuildingRepository     providers.BuildingRepository
}

func (usecase Implementation) Execute(ctx context.Context, request save_announcement.Request, userID primitive.ObjectID) error {
	user, err := usecase.UserRepository.GetUserById(ctx, userID)
	if err != nil {
		return err
	}
	building, err := usecase.BuildingRepository.GetBuildingById(ctx, request.BuildingID)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}
	if building == nil {
		return errors.NewBadRequestError(fmt.Sprintf("building not found. building_id: %s", request.BuildingID.String()))
	}
	announcement := entities.Announcement{
		User:     user.ToLw(),
		Building: building.ToLw(),
		Message:  request.Message,
		Date:     request.Date,
		Severity: request.Severity,
	}
	err = usecase.AnnouncementRepository.Save(ctx, announcement)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}
	return nil
}
