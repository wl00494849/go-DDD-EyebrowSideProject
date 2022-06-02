package repository

import "go-DDD/domain/entity"

type SysCodeRepo interface {
	GetSysCode(types string, code string) *entity.Sys_code
	GetSysCodes(types string) *[]entity.Sys_code
}
