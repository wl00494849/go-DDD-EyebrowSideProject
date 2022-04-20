package entity

import "time"

type LatestNews struct {
	NewID      string
	Title      string
	Content    string
	CreateTime time.Time
	Images     [][]byte
}
