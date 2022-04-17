package persistence

import (
	"database/sql"
	"fmt"
	"go-DDD/domain/entity"
	"go-DDD/domain/repository"
	"time"
)

type LatestNewsRepo struct {
	db *sql.DB
}

func NewLatestNewRepo(db *sql.DB) *LatestNewsRepo {
	return &LatestNewsRepo{db}
}

//implements newRepo interface
var _ repository.LatestNewsRepo = &LatestNewsRepo{}

func (r *LatestNewsRepo) GetNewDetail(newID string) *entity.LatestNews {
	fmt.Println(newID)
	news := &entity.LatestNews{}
	row := r.db.QueryRow("Select * from lastestNew where id=?", newID)
	err := row.Scan(&news.NewID, &news.Title, &news.Content, &news.CreateTime)
	if err != nil {
		panic(err)
	}
	return news
}

func (r *LatestNewsRepo) SetNew(new *entity.LatestNews) error {
	stmt, _ := r.db.Prepare("Insert lastestNew set title=?,contant=?,createTime=?")
	_, err := stmt.Exec(new.Title, new.Content, time.Now().Local())
	return err
}
