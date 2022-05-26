package repository

import "go-DDD/domain/entity"

type AnnouncementRepo interface {
	GetNewDetail(newId string) *entity.Announcement
	CreateNew(new *entity.Announcement) error
	DeleteNew(newId string) error
	GetNewList() *[]entity.Announcement
}
