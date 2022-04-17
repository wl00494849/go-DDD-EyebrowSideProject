package entity

import "time"

type LatestNews struct {
	NewID      int
	Title      string
	Content    string
	CreateTime time.Time
	Images     [][]byte
}
