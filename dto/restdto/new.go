package restdto

import (
	"time"
)

type LatestNewsList struct {
	NewID      string
	Title      string
	CreateTime time.Time
}
