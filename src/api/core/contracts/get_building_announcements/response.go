package get_building_announcements

import (
	"time"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/contracts/common"
	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Response struct {
	Announcements []Announcement `json:"announcements"`
}

type Announcement struct {
	ID       primitive.ObjectID `json:"id"`
	User     common.UserLw      `json:"user"`
	Building common.BuildingLw  `json:"building"`
	Title    string             `json:"title"`
	Message  string             `json:"message"`
	Date     time.Time          `json:"date"`
	Severity constants.Severity `json:"severity"`
}

func (resp *Response) AddAnnouncement(a entities.Announcement) {
	announcement := Announcement{
		ID:       a.ID,
		User:     common.NewUserLw(a.User),
		Building: common.NewBuildingLw(a.Building),
		Title:    a.Title,
		Message:  a.Message,
		Date:     a.Date,
		Severity: a.Severity,
	}
	resp.Announcements = append(resp.Announcements, announcement)
}
