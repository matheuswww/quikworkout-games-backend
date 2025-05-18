package logger

import (
	"fmt"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zapLogger *zap.Logger
)

func LoadLogger() {
	absolutePath, err := os.Getwd()
	if err != nil {
		log.Fatalf("error trying get abs path, err: %v", err)
	}
	path := fmt.Sprintf("%s/log/log.txt", absolutePath)
	paths := []string{"stdout"}
	test := os.Getenv("TEST")
	if test != "TRUE" {
		paths = append(paths, path)
	}
	logConfig := zap.Config{
		OutputPaths: paths,
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	zapLogger, _ = logConfig.Build()
}

func Info(message string, tags ...zap.Field) {
	zapLogger.Info(message, tags...)
	zapLogger.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	zapLogger.Error(message, tags...)
	zapLogger.Sync()
}