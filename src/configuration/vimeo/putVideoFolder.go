package vimeo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"go.uber.org/zap"
)

type PutVideoFolderRequest struct {
	Items []Item `json:"items"`
}

type Item struct {
	Uri string `json:"uri"`
}

func putVideoFolder(uri string) error {
	jsonReq := PutVideoFolderRequest{
		[]Item{
			{
				Uri: uri,
			},
		},
	}

	jsonBytes, err := json.Marshal(jsonReq)
	if err != nil {
		logger.Error("Error trying Marshal json", err, zap.String("journey", "UploadVideo"))
		return nil
	}
	
	jsonReader := bytes.NewReader(jsonBytes)
	url := "https://api.vimeo.com/users/"+os.Getenv("VIMEO_USER_ID")+"/projects/"+os.Getenv("EDITION")+"/items"
	req, err := http.NewRequest("POST", url, jsonReader)
	if err != nil {
		logger.Error("Error trying NewRequest", err, zap.String("journey", "UploadVideo"))
		return nil
	}
	
	token := os.Getenv("VIMEO_TOKEN")
	req.Header.Set("Authorization", "Bearer " + token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.vimeo.*+json;version=3.4")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error trying ReadAll", err, zap.String("journey", "PutVideoFolder"))
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		logger.Error("Error trying PutVideoFolder", errors.New(string(body)), zap.String("journey", "PutVideoFolder"))
		return errors.New("was not possible upload video")
	}

	return nil
}