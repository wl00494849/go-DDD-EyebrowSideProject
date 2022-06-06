package repository

import "go-DDD/domain/entity"

type SubcategoryRepo interface {
	GetSubCategorys(code string) *[]entity.Subcategory
	CreateSubCategory(data *entity.Subcategory) error
	DeleteSubCategory(code string) error
	// UploadSubCategory(*[]entity.Subcategory) error
}
