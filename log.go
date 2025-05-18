package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	server "github.com/matheuswww/quikworkout-games-backend/log"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/email"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"go.uber.org/zap"
)


var (
	logPath      = "./log/log.txt"
	lastPositionPath = "./log/last_position.txt"
	lastPosition  int64
	maxLogSizeMB = 5

	lastNoti *time.Time = nil

	to = "teteus.30.08.07@outlook.com"
)

func checkLogs() {
	logger.Info("Init check logs", zap.String("journey", "checklogs"))
	rotateLog()
	lastPosition = loadLastPosition()
	info, err := os.Stat(logPath)
	if err != nil {
		log.Fatal("erro trying get stat", err, zap.String("journey", "checklogs"))
	}
	if lastPosition > info.Size() {
		saveLastPosition(info.Size())
	}
	for {
		file, err := os.OpenFile(logPath, os.O_RDONLY, 0644)
		if err != nil {
			if os.IsNotExist(err) {
				create()
				continue
			}
			logger.Error("error trying to open file", err)
			sendNotification("Houve um erro ao tentar abrir o arquivo de log")
			saveLastPosition(0)
			continue
		}
		_, err = file.Seek(lastPosition, 0)
		if err != nil {
			logger.Error("error seeking file", err)
			sendNotification("Houve um erro ao tentar buscar no arquivo de log")
			file.Close()
			saveLastPosition(0)
			continue
		}
		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					lastPosition, err = file.Seek(0, io.SeekCurrent)
					saveLastPosition(lastPosition)
					if err != nil {
						logger.Error("error seeking file", err)
						sendNotification("Houve um erro ao tentar buscar no arquivo de log")
						file.Close()
						saveLastPosition(0)
						continue
					}
					break
				}
				file.Close()
				logger.Error("error reading string", err)
				sendNotification("Houve um erro ao tentar ler a string do arquivo de log")
				break
			}
			if strings.Contains(line, "| 500 |") {
				sendNotification(fmt.Sprintf("Um erro 500 foi encontrado, log: %s", line))
			}
		}
		file.Close()
		rotateLog()
		time.Sleep(time.Second * 60)
	}
}

func rotateLog() {
	info, err := os.Stat(logPath)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		logger.Error("error getting file stat", err)
		sendNotification("Houve um erro ao tentar obter as informações do arquivo de log")
		return
	}
	if info.Size() >= int64(maxLogSizeMB*1024*1024) {
		clear()
	}
}

func clear() {
	err := os.Remove(logPath)
	if err != nil {
		logger.Error("error trying to remove the log file", err)
		sendNotification("Houve um erro ao tentar remover o arquivo de log")
		return
	}
	saveLastPosition(0)
	create()
}

func create() {
	newFile, err := os.Create(logPath)
	if err != nil {
		logger.Error("error trying to create a new log file", err)
		sendNotification("Houve um erro ao tentar criar um novo arquivo de log")
		return
	}
	newFile.Close()
}

func loadLastPosition() int64 {
	file, err := os.OpenFile(lastPositionPath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return 0
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		pos, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err == nil {
			return pos
		}
	}
	return 0
}

func saveLastPosition(pos int64) {
	file, err := os.OpenFile(lastPositionPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		logger.Error("error trying to open last_position.txt", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(strconv.FormatInt(pos, 10))
	if err != nil {
		logger.Error("error trying to write last position", err)
	}
	lastPosition = pos
}

func sendNotification(message string) {
	if(lastNoti == nil || time.Now().After(lastNoti.Add(time.Minute * 10))) {
		err := email.NewEmail().NewEmailConnection(to, "um erro foi encontrado", []byte(server.GetHtml(message)))
		if err != nil {
			logger.Error("error trying send email", err, zap.String("journey", "sendNotification"))
			return
		}
		if(lastNoti == nil) {
			lastNoti = &time.Time{}
		}
		*lastNoti = time.Now()
		logger.Info("notification sended with success", zap.String("journey", "sendNotification"))
	}
}