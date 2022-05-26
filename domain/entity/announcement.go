package entity

import "time"

type Announcement struct {
	NewID       string
	Title       string
	Content     string
	Statu       bool
	Shelf_Time  time.Time
	Create_Time time.Time
	IsTop       bool
}
