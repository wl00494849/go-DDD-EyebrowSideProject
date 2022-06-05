package repository

import "go-DDD/domain/entity"

type SubcategoryRepo interface {
	GetSubCategory(code string) *[]entity.Subcategory
	CreateSubCategory()
}
