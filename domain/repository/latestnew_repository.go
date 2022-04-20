package repository

import "go-DDD/domain/entity"

type LatestNewsRepo interface {
	GetNewDetail(newId string) *entity.LatestNews
	CreateNew(new *entity.LatestNews) error
	DeleteNew(newId string) error
}
