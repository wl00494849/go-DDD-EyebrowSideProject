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
	GetNewDetail(newID string) *restdto.AnnouncementDetail
	CreateNew(new entity.Announcement) error
	DeleteNew(newID string) error
	GetListNew() *[]restdto.AnnouncementList
}

func CreateNewLastestNewApp(rl repository.AnnouncementRepo) *LastestNewApp {
	return &LastestNewApp{
		rl: rl,
	}
}

func (r *LastestNewApp) GetNewDetail(newID string) *restdto.AnnouncementDetail {
	data := r.rl.GetNewDetail(newID)
	return &restdto.AnnouncementDetail{
		Id:          data.Id,
		Title:       data.Title,
		Content:     data.Content,
		Create_Time: data.Create_Time,
	}
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
		ls.IsTop = v.IsTop
		rto = append(rto, ls)
	}

	return &rto
}
