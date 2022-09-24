package get_building_announcements

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/get_building_announcements"
)

type UseCase interface {
	Execute(ctx context.Context, request get_building_announcements.Request)
}
