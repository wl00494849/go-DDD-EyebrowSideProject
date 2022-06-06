package persistence

import (
	"database/sql"
	"go-DDD/domain/entity"
	"go-DDD/domain/repository"
)

type SubcategoryRepo struct {
	db *sql.DB
}

func NewSubcategoryRepo(db *sql.DB) *SubcategoryRepo {
	return &SubcategoryRepo{db}
}

var _ repository.SubcategoryRepo = &SubcategoryRepo{}

func (r *SubcategoryRepo) GetSubCategorys(code string) *[]entity.Subcategory {
	sli := make([]entity.Subcategory, 0)
	rows, _ := r.db.Query("Select Id,SubCategory_Name,Price,SubCategory_Order from SubCategory where Category_Id=?", code)
	for rows.Next() {
		var data entity.Subcategory
		rows.Scan(&data.Id, &data.Subcategory_Name, &data.Price, &data.Subcategory_Order)
		sli = append(sli, data)
	}
	return &sli
}

func (r *SubcategoryRepo) CreateSubCategory(data *entity.Subcategory) error {
	stmt, _ := r.db.Prepare("Insert SubCategory set Category_Id=?,SubCategory_Name=?,Price=?,SubCategory_Order=?")
	_, err := stmt.Exec(data.Category_Id, data.Subcategory_Name, data.Price, data.Subcategory_Order)
	return err
}

func (r *SubcategoryRepo) DeleteSubCategory(code string) error {
	_, err := r.db.Exec("delete from SubCategory where Id=?", code)
	return err
}
