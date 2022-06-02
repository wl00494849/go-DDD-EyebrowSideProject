package persistence

import (
	"database/sql"
	"go-DDD/domain/entity"
	"go-DDD/domain/repository"
)

type SysCodeRepo struct {
	db *sql.DB
}

func NewSysCodeRepo(db *sql.DB) *SysCodeRepo {
	return &SysCodeRepo{db}
}

var _ repository.SysCodeRepo = &SysCodeRepo{}

func (r *SysCodeRepo) GetSysCode(types string, code string) *entity.Sys_code {
	sc := &entity.Sys_code{}
	row := r.db.QueryRow("Select SC_Type,SC_Code,SC_Desc,SC_Order from Sys_Code where SC_Type=? amd SC_Code=?", types, code)
	err := row.Scan(&sc.SC_Type, &sc.SC_Code, &sc.SC_Desc, &sc.SC_Order)
	if err != nil {
		panic(err)
	}
	return sc
}
func (r *SysCodeRepo) GetSysCodes(types string) *[]entity.Sys_code {
	sli := make([]entity.Sys_code, 0)
	rows, _ := r.db.Query("Select SC_Type,SC_Code,SC_Desc,SC_Order from Sys_Code where SC_Type=?")
	for rows.Next() {
		var sc entity.Sys_code
		rows.Scan(&sc.SC_Type, &sc.SC_Code, &sc.SC_Type, &sc.SC_Order)
		sli = append(sli, sc)
	}
	return &sli
}
