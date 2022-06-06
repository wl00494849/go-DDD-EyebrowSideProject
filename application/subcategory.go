package application

import (
	"go-DDD/domain/entity"
	"go-DDD/domain/repository"
)

type SubcategoryApp struct {
	rs repository.SubcategoryRepo
}

type SubcategoryInterfact interface {
	GetSubCategorys(code string) *[]entity.Subcategory
	CreateSubCategory(data *entity.Subcategory) error
	DeleteSubCategory(code string) error
}

func CreateSysCodeApp(rs repository.SubcategoryRepo) *SubcategoryApp {
	return &SubcategoryApp{rs: rs}
}

var _ SubcategoryInterfact = &SubcategoryApp{}

func (r *SubcategoryApp) GetSubCategorys(code string) *[]entity.Subcategory {
	return r.rs.GetSubCategorys(code)
}
func (r *SubcategoryApp) CreateSubCategory(data *entity.Subcategory) error {
	return r.rs.CreateSubCategory(data)
}
func (r *SubcategoryApp) DeleteSubCategory(code string) error {
	return r.rs.DeleteSubCategory(code)
}
