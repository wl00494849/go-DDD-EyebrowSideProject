package restdto

import (
	"time"
)

type AnnouncementList struct {
	NewID      string
	Title      string
	CreateTime time.Time
}
