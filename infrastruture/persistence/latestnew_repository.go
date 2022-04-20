package persistence

import (
	"database/sql"
	"fmt"
	"go-DDD/domain/entity"
	"go-DDD/domain/repository"
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

func (r *LatestNewsRepo) CreateNew(new *entity.LatestNews) error {
	stmt, _ := r.db.Prepare("Insert lastestNew set id=?,title=?,contant=?,createTime=?")
	_, err := stmt.Exec(new.NewID, new.Title, new.Content, new.CreateTime)
	return err
}

func (r *LatestNewsRepo) DeleteNew(newId string) error {
	_, err := r.db.Exec("delete from lastestNew where id=?", newId)
	return err
}

func (r *LatestNewsRepo) GetNewList() *[]entity.LatestNews {
	sli := make([]entity.LatestNews, 0)
	rows, _ := r.db.Query("Select id,title,createTime from lastestNew")
	defer rows.Close()
	for rows.Next() {
		var data entity.LatestNews
		rows.Scan(&data.NewID, &data.Title, &data.CreateTime)
		sli = append(sli, data)
	}
	return &sli
}
