package application

import (
	"go-DDD/domain/entity"
	"go-DDD/domain/repository"
	"go-DDD/dto/restdto"
	"sort"
)

type SubcategoryApp struct {
	rs repository.SubcategoryRepo
}

type SubcategoryInterfact interface {
	GetSubCategorys(code string) *[]restdto.Subcategory
	CreateSubCategory(data *entity.Subcategory) error
	DeleteSubCategory(code string) error
}

func CreateSysCodeApp(rs repository.SubcategoryRepo) *SubcategoryApp {
	return &SubcategoryApp{rs: rs}
}

var _ SubcategoryInterfact = &SubcategoryApp{}

func (r *SubcategoryApp) GetSubCategorys(code string) *[]restdto.Subcategory {
	data := make([]restdto.Subcategory, 0)
	rd := r.rs.GetSubCategorys(code)
	for _, v := range *rd {
		var d restdto.Subcategory
		d.Id = v.Id
		d.SubCategory_Name = v.Subcategory_Name
		d.Price = v.Price
		d.Subcategory_Order = v.Subcategory_Order
		data = append(data, d)
	}
	sort.Slice(data, func(i, j int) bool { return data[i].Subcategory_Order < data[j].Subcategory_Order })
	return &data
}
func (r *SubcategoryApp) CreateSubCategory(data *entity.Subcategory) error {
	return r.rs.CreateSubCategory(data)
}
func (r *SubcategoryApp) DeleteSubCategory(code string) error {
	return r.rs.DeleteSubCategory(code)
}
