package save_announcement

import store "flay-api-v3.0/src/api/infraestructure/repositories"

type Implementation struct {
	AnnouncementRepository *store.AnnouncementRepository
}
