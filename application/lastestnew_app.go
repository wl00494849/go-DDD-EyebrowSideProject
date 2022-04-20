package application

import (
	"go-DDD/domain/entity"
	"go-DDD/domain/repository"
	"go-DDD/dto/restdto"
	"time"
)

type LastestNewApp struct {
	rl repository.LatestNewsRepo
}

var _ LastestNewAppInterFace = &LastestNewApp{}

type LastestNewAppInterFace interface {
	GetNewDetail(newID string) *entity.LatestNews
	CreateNew(new entity.LatestNews) error
	DeleteNew(newID string) error
	GetListNew() *[]restdto.LatestNewsList
}

func CreateNewLastestNewApp(rl repository.LatestNewsRepo) *LastestNewApp {
	return &LastestNewApp{
		rl: rl,
	}
}

func (r *LastestNewApp) GetNewDetail(newID string) *entity.LatestNews {
	return r.rl.GetNewDetail(newID)
}

func (r *LastestNewApp) CreateNew(new entity.LatestNews) error {
	new.NewID = generatorID()
	new.CreateTime = time.Now().Local()
	err := r.rl.CreateNew(&new)
	return err
}

func (r *LastestNewApp) DeleteNew(newId string) error {
	return r.rl.DeleteNew(newId)
}

func (r *LastestNewApp) GetListNew() *[]restdto.LatestNewsList {
	rto := make([]restdto.LatestNewsList, 0)
	datas := r.rl.GetNewList()
	for _, v := range *datas {
		var ls restdto.LatestNewsList
		ls.NewID = v.NewID
		ls.Title = v.Title
		ls.CreateTime = v.CreateTime
		rto = append(rto, ls)
	}

	return &rto
}
