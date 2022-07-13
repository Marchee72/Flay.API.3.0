package book_common_space

import (
	"context"

	book_common_space "flay-api-v3.0/src/api/core/contracts/book_common_space"
	"flay-api-v3.0/src/api/core/entities"
	"flay-api-v3.0/src/api/core/errors"
	"flay-api-v3.0/src/api/core/providers"
	"flay-api-v3.0/src/api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UseCase interface {
	Execute(ctx context.Context, request book_common_space.Request, userID primitive.ObjectID) (*book_common_space.Response, error)
}

type Implementation struct {
	PenaltyRepository providers.PenaltyRepository
	BookingRepository providers.BookingRepository
	UserRepository    providers.UserRepository
}

func (useCase *Implementation) Execute(ctx context.Context, request book_common_space.Request, userID primitive.ObjectID) (*book_common_space.Response, error) {
	//Check if user has active penalties
	user, err := useCase.UserRepository.GetUserById(ctx, userID)
	if err != nil || user == nil {
		return nil, errors.NewInternalServerError("error searching user with provided credentials")
	}
	penalty, err := useCase.PenaltyRepository.GetActivePenalties(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	response := book_common_space.Response{}
	if penalty != nil {
		response.SetPenalty(*penalty)
		return &response, nil
	}
	date, err := utils.Date(request.Date)
	if err != nil {
		return nil, err
	}

	isAbailable, err := useCase.BookingRepository.IsAbailable(ctx, user.Building.ID, request.CommonSpace, date, request.Shift)
	response.IsAbailable = isAbailable
	if !isAbailable {
		return &response, nil
	}
	if err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}
	booking := entities.Booking{
		User:        user.ToLw(),
		Building:    *user.Building,
		CommonSpace: request.CommonSpace,
		Date:        date,
		Shift:       request.Shift,
	}
	err = useCase.BookingRepository.Save(ctx, booking)
	if err != nil {
		return nil, err
	}
	response.SetBooking(booking)
	return &response, nil

	//Check if common space is abailable that day at that timeSs
	//Book space or return error if its not abailable
}
