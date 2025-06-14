package user_service_util

import (
	"os"
	"path/filepath"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
)

func GetUserImage(user string) (string, *rest_err.RestErr) {
	absPath, err := filepath.Abs("images")
	if err != nil {
		return "", rest_err.NewInternalServerError("servere error")
	}
	files, err := os.ReadDir(absPath)
	if err != nil {
		return "", rest_err.NewInternalServerError("server error")
	}
	for _, file := range files {
		if !file.IsDir() {
			fileName := file.Name()
			nameWithoutExt := fileName[:len(fileName)-len(filepath.Ext(fileName))]
			if nameWithoutExt == user {
				url := os.Getenv("URL")
				return url+"/images/"+fileName, nil
			}
		}
	}
	return "", nil
}