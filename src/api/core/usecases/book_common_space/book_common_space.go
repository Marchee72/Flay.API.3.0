package book_common_space

import (
	"context"

	book_common_space "flay-api-v3.0/src/api/core/contracts/book_common_space"
	"flay-api-v3.0/src/api/core/entities"
	"flay-api-v3.0/src/api/core/providers"
)

type UseCase interface {
	Execute(ctx context.Context, request book_common_space.Request) (*book_common_space.Response, error)
}

type Implementation struct {
	PenaltyRepository providers.PenaltyRepository
	BookingRepository providers.BookingRepository
}

func (useCase *Implementation) Execute(ctx context.Context, request book_common_space.Request) (*book_common_space.Response, error) {
	//Check if user has active penalties
	penalty, err := useCase.PenaltyRepository.GetActivePenalties(ctx, request.User.ID)
	if err != nil {
		return nil, err
	}
	response := book_common_space.Response{}
	if penalty != nil {
		response.SetPenalty(*penalty)
		return &response, nil
	}
	isAbailable, err := useCase.BookingRepository.IsAbailable(ctx, request.CommonSpace.ID, request.StartDate, request.EndDate)
	if err != nil {
		return nil, err
	}
	response.IsAbailable = isAbailable
	if !isAbailable {
		return &response, nil
	}
	booking := entities.Booking{
		User:        request.User,
		Building:    request.Builging,
		CommonSpace: request.CommonSpace,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
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
