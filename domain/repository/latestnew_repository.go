package repository

import "go-DDD/domain/entity"

type LatestNewsRepo interface {
	GetNewDetail(newId string) *entity.LatestNews
	SetNew(new *entity.LatestNews) error
}
