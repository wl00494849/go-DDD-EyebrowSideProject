package entity

import "time"

type LatestNews struct {
	NewID      string
	Content    string
	CreateTime time.Duration
	Images     [][]byte
}
