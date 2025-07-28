package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/database/mysql"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

var (
	name     = "Flavia Quik"
	email    = "teteus.30.08.07@outlook.com"
)

func main() {
	loadEnv()
	logger.LoadLogger()
	password := os.Getenv("JUDGE_PASSWORD")
	if password == "" {
		logger.Error("Error trying get env", errors.New("password is nil"), zap.String("journey", "insert judge"))
		log.Fatal("password is nil")
	}
	mysql := loadMysql()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		log.Fatalf("error trying insert judge: %v", err)
	}

	id := uuid.NewString()

	query := "INSERT INTO judge (judge_id, name, email, password) VALUES (?, ?, ?, ?)"
	_, err = mysql.ExecContext(ctx, query, id, name, email, hash)
	if err != nil {
		log.Fatalf("error trying insert judge: %v", err)
	}
	logger.Info("judge inserted with success", zap.String("journey", "insert judge"))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", err)
		log.Fatal("Error loading .env file")
	}
}

func loadMysql() *sql.DB {
	user := os.Getenv("MYSQL_USER")
	host := os.Getenv("MYSQL_HOST")
	name := os.Getenv("MYSQL_NAME")
	password := os.Getenv("MYSQL_PASSWORD")
	port, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		logger.Error("Error loading database", err)
		log.Fatal("invalid port")
	}
	if user == "" || host == "" || name == "" || password == "" || port == 0 {
		logger.Error("Error loading env", err)
		log.Fatal("error loading env")
	}
	db, err := mysql.NewMysql(user, host, name, password, port).NewMysqlConnection()
	if err != nil {
		log.Fatal("Error connecting")
	}
	return db
}
