package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	model_util "github.com/matheuswww/quikworkout-games-backend/src/model/util"
	"github.com/matheuswww/quikworkout-games-backend/src/routes"
)

func main() {
	loadEnv()
	logger.LoadLogger()
	logger.Info("About to start user application")
	file, err := os.OpenFile("./log/log.txt",  os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("error trying open log")
	}
	defer file.Close()
	multiWriter := io.MultiWriter(os.Stdout, file)
	gin.DefaultWriter = multiWriter
	gin.DefaultErrorWriter = multiWriter
	router := gin.Default()
	corsConfig := loadCors()
	mysql := loadMysql()
	model_util.InitDb(mysql)
	router.Use(cors.New(*corsConfig))
	router.Static("/images", "./images")
	router.GET("/pdf/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		filename = filepath.Base(filename)
		filepath := filepath.Join("./pdf", filename)
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Arquivo não encontrado"})
				return
		}
		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Header("Content-Type", "application/octet-stream")
		c.File(filepath)
	})
	routes.InitRoutes(&router.RouterGroup, mysql)
	chanError := make(chan error)
	go graceFullyShutdown(router, "8081", chanError)
	if err := <-chanError; err != nil {
		log.Fatal(err)
	}
}

func graceFullyShutdown(handler http.Handler, addr string, chanError chan error) {
	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", addr),
		Handler:           handler,
		ReadHeaderTimeout: 30 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      1 * time.Minute,
		IdleTimeout:       2 * time.Minute,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			chanError <- fmt.Errorf("an erro ocurred while trying to start application on port %s. Error %v", addr, err)
			return
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	
	go checkLogs()
	<-ctx.Done()
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	logger.Info("Received a shutdown signal, quiting...")

	defer func() {
		stop()
		cancel()
		close(chanError)
	}()

	err := server.Shutdown(ctxTimeout)
	if err != nil {
		logger.Error("Error trying shutdown", err)
	}
	logger.Info("Shutdown completed")
}
