package mysql

import (
	"database/sql"
	"strconv"
	"time"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
)

func (m *mysql) NewMysqlConnection() (*sql.DB, error) {
	port := strconv.Itoa(m.port)
	db, err := sql.Open("mysql", m.user+":"+m.password+"@tcp("+m.host+":"+port+")/"+m.name)
	if err != nil {
		logger.Error("MYSQL DB CONNECT ERROR!!!", err, zap.String("journey", "databaseConnect"))
		return nil, err
	}
	if err := db.Ping(); err != nil {
		logger.Error("MYSQL DB CONNECT ERROR!!!", err, zap.String("journey", "databaseConnect"))
		return nil, err
	}
	logger.Info("MYSQL DB IS RUNNING!!!")
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Minute * 5)
	return db, nil
}
