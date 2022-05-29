package repository

type SysCodeRepo interface {
	GetSysCode(types string, code string) string
	GetSysCodes(types string) []string
}
