package entity

import "time"

type Announcement struct {
	Id          int
	Title       string
	Content     string
	Statu       bool
	Shelf_Time  time.Time `time_format:"2006-01-02"`
	Create_Time time.Time
	IsTop       bool
}
