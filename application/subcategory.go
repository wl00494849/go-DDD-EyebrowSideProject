package application

import (
	"go-DDD/domain/entity"
	"go-DDD/domain/repository"
)

type SubcategoryApp struct {
	rs repository.SysCodeRepo
}

type SubcategoryInterfact interface {
	GetSubcategorys(code string) *[]entity.Sys_code
}

func CreateSysCodeApp(rs repository.SysCodeRepo) *SubcategoryApp {
	return &SubcategoryApp{rs: rs}
}

var _ SubcategoryInterfact = &SubcategoryApp{}

func (r *SubcategoryApp) GetSubcategorys(code string) *[]entity.Sys_code {
	return r.rs.GetSysCodes(code)
}
