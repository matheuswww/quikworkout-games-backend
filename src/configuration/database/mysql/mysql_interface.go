package mysql

import (
	"database/sql"
)

func NewMysql(user, host, name, passoword string, port int) Mysql {
	return &mysql{
		user,
		host,
		name,
		passoword,
		port,
	}
}

type Mysql interface {
	NewMysqlConnection() (*sql.DB, error)
}

type mysql struct {
	user     string
	host     string
	name     string
	password string
	port     int
}
