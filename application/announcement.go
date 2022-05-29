package application

import (
	"go-DDD/domain/entity"
	"go-DDD/domain/repository"
	"go-DDD/dto/restdto"
	"time"
)

type LastestNewApp struct {
	rl repository.AnnouncementRepo
}

var _ LastestNewAppInterFace = &LastestNewApp{}

type LastestNewAppInterFace interface {
	GetNewDetail(newID string) *entity.Announcement
	CreateNew(new entity.Announcement) error
	DeleteNew(newID string) error
	GetListNew() *[]restdto.AnnouncementList
}

func CreateNewLastestNewApp(rl repository.AnnouncementRepo) *LastestNewApp {
	return &LastestNewApp{
		rl: rl,
	}
}

func (r *LastestNewApp) GetNewDetail(newID string) *entity.Announcement {
	return r.rl.GetNewDetail(newID)
}

func (r *LastestNewApp) CreateNew(new entity.Announcement) error {
	new.Shelf_Time = time.Now().Local()
	new.Create_Time = time.Now().Local()
	err := r.rl.CreateNew(&new)
	return err
}

func (r *LastestNewApp) DeleteNew(newId string) error {
	return r.rl.DeleteNew(newId)
}

func (r *LastestNewApp) GetListNew() *[]restdto.AnnouncementList {
	rto := make([]restdto.AnnouncementList, 0)
	datas := r.rl.GetNewList()
	for _, v := range *datas {
		var ls restdto.AnnouncementList
		ls.Id = v.Id
		ls.Title = v.Title
		ls.CreateTime = v.Create_Time
		rto = append(rto, ls)
	}

	return &rto
}
