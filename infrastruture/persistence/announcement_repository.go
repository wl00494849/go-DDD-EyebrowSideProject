package persistence

import (
	"database/sql"
	"fmt"
	"go-DDD/domain/entity"
	"go-DDD/domain/repository"
)

type AnnouncementRepo struct {
	db *sql.DB
}

func NewAnnouncementRepo(db *sql.DB) *AnnouncementRepo {
	return &AnnouncementRepo{db}
}

//implements newRepo interface
var _ repository.AnnouncementRepo = &AnnouncementRepo{}

func (r *AnnouncementRepo) GetNewDetail(newID string) *entity.Announcement {
	fmt.Println(newID)
	news := &entity.Announcement{}
	row := r.db.QueryRow("Select * from Announcement where Id=?", newID)
	err := row.Scan(&news.Id, &news.Title, &news.Content, &news.Create_Time)
	if err != nil {
		panic(err)
	}
	return news
}

func (r *AnnouncementRepo) CreateNew(new *entity.Announcement) error {
	fmt.Printf("%+v", new)
	stmt, _ := r.db.Prepare("Insert Announcement set Title=?,Content=?,Shelf_Time=?,Create_Time=?,IsTop=?,Statu=?")
	_, err := stmt.Exec(new.Title, new.Content, new.Shelf_Time, new.Create_Time, new.IsTop, new.Statu)
	return err
}

func (r *AnnouncementRepo) DeleteNew(newId string) error {
	_, err := r.db.Exec("delete from Announcement where Id=?", newId)
	return err
}

func (r *AnnouncementRepo) GetNewList() *[]entity.Announcement {
	sli := make([]entity.Announcement, 0)
	rows, _ := r.db.Query("Select Id,Title,Create_Time from Announcement")
	defer rows.Close()
	for rows.Next() {
		var data entity.Announcement
		rows.Scan(&data.Id, &data.Title, &data.Create_Time)
		sli = append(sli, data)
	}
	return &sli
}
