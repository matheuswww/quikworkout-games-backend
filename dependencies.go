package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"

	"github.com/joho/godotenv"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/database/mysql"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", err)
		log.Fatal("Error loading .env file")
	}
}

func loadCors() *cors.Config {
	config := cors.DefaultConfig()
	origin_1 := os.Getenv("CORSORIGIN_1")
	origin_2 := os.Getenv("CORSORIGIN_2")
	if origin_1 == "" || origin_2 == "" {
		log.Fatal("corsorigin not found")
	}
	config.AllowOrigins = []string{origin_1,origin_2}
	config.AllowMethods = []string{"POST", "GET", "DELETE", "PATCH", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Cookie"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	return &config
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
		logger.Error("Error loading database database", err)
		log.Fatal("Error connecting")
	}
	return db
}
