package persistence

import (
	"database/sql"
	"go-DDD/domain/repository"
)

type SysCodeRepo struct {
	db *sql.DB
}

func NewSysCodeRepo(db *sql.DB) *SysCodeRepo {
	return &SysCodeRepo{db}
}

var _ repository.SysCodeRepo = &SysCodeRepo{}

func (r *SysCodeRepo) GetSysCode(types string, code string) string {

	return ""
}
func (r *SysCodeRepo) GetSysCodes(types string) []string {
	data := make([]string, 0)
	return data
}
