package restdto

import (
	"time"
)

type AnnouncementList struct {
	Id         int
	Title      string
	CreateTime time.Time
}
type AnnouncementCreate struct {
	Title      string
	Content    string
	Statu      bool
	Shelf_Time string
	IsTop      bool
}
