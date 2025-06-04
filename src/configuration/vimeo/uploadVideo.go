package vimeo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"go.uber.org/zap"
)

type UploadRequest struct {
	Upload 	upload  `json:"upload"`
	Name   	string  `json:"name"`
	Privacy privacy `json:"privacy"`
}

type upload struct {
	Approach string    `json:"approach"`
	Size     int64     `json:"size"`
	RedirectUrl string `json:"redirect_url"`
}

type privacy struct {
	Embed    string `json:"embed"`
	View     string `json:"view"`
}

type UploadResponse struct {
	Uri    string         `json:"uri"`
	Upload uploadResponse `json:"upload"`
}

type uploadResponse struct {
	Form string `json:"form"`
}

func UploadVideo(name string, size int64) (string, string, error) {
	jsonReq := UploadRequest{
		Name: name,
		Upload: upload{
			Approach: "post",
			Size:     size,
			RedirectUrl: os.Getenv("CORSORIGIN_1")+"/conta/minha-conta?sent=true",
		},
		Privacy: privacy{
			Embed: "public",
			View: "anybody",
		},
	}
	jsonBytes, err := json.Marshal(jsonReq)
	if err != nil {
		logger.Error("Error trying Marshal json", err, zap.String("journey", "UploadVideo"))
		return "", "", err
	}

	jsonReader := bytes.NewReader(jsonBytes)
	url := "https://api.vimeo.com/me/videos"

	req, err := http.NewRequest("POST", url, jsonReader)
	if err != nil {
		logger.Error("Error trying NewRequest", err, zap.String("journey", "UploadVideo"))
		return "", "", err
	}

	token := os.Getenv("VIMEO_TOKEN")
	req.Header.Set("Authorization", "Bearer " + token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.vimeo.*+json;version=3.4")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error("Error trying ", err, zap.String("journey", "PutVideoFolder"))
		return "", "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error trying ReadAll", err, zap.String("journey", "UploadVideo"))
		return "", "", err
	}

	if resp.StatusCode != http.StatusCreated {
		logger.Error("Error trying UploadVideo", errors.New(string(body)), zap.String("journey", "UploadVideo"))
		return "", "", errors.New("was not possible upload video")
	}
	var jsonResp UploadResponse
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		logger.Error("Error trying NewDecoder", err, zap.String("journey", "UploadVideo"))
		return "", "", err
	}
	parts := strings.Split(jsonResp.Uri, "/")
	id := parts[2]
	err = putVideoFolder(jsonResp.Uri)
	if err != nil {
		return "", "", err
	}
	return jsonResp.Upload.Form, id, nil
}
