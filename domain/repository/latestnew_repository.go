package repository

import "go-DDD/domain/entity"

type LatestNewsRepo interface {
	GetNew(newId string)
	GetNewsList()
	SetNew(new *entity.LatestNews)
}
