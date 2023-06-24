package get_admin_announcements

import (
	"flay-api-v3.0/src/api/core/contracts/common"
	"flay-api-v3.0/src/api/core/entities"
)

type Response struct {
	Announcements []common.Announcement `json:"announcements"`
}

func (resp *Response) AddAnnouncement(a entities.Announcement) {
	announcement := common.Announcement{
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
