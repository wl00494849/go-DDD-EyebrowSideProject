package persistence

import (
	"database/sql"
	"go-DDD/domain/repository"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Repositories struct {
	AnnouncementRepo repository.AnnouncementRepo
	db               *sql.DB
}

func NewRepositories(us, pwd, addr, dbName string) (*Repositories, error) {
	config := mysql.Config{
		User:                 us,
		Passwd:               pwd,
		Addr:                 addr,
		Net:                  "tcp",
		DBName:               dbName,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	d, _ := sql.Open("mysql", config.FormatDSN())
	d.SetMaxOpenConns(5)
	//max idle connection
	d.SetMaxIdleConns(2)
	//idle time delete connection
	d.SetConnMaxLifetime(time.Hour)
	//Connection Test
	err := d.Ping()
	if err != nil {
		return nil, err
	}

	return &Repositories{
		AnnouncementRepo: NewAnnouncementRepo(d),
		db:               d,
	}, nil
}

func (r *Repositories) Close() {
	r.db.Close()
}
