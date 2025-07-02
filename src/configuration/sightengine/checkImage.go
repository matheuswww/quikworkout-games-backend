package sightengine

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	"go.uber.org/zap"
)

type Response struct {
	Status  string  `json:"status"`
	Summary Summary `json:"summary"`
}

type Summary struct {
	Action       string         `json:"action"`
	RejectReason []RejectReason `json:"reject_reason"`
}

type RejectReason struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

var (
	user = "556621385"
)

func CheckImage(fileHeader *multipart.FileHeader, fileName string) *rest_err.RestErr {
	if os.Getenv("TEST") == "TRUE" {
		return nil
	}
	secret := os.Getenv("SIGHTENGINE_TOKEN")
	workflow := os.Getenv("SIGHTENGINE_WORKFLOW")
	if secret == "" || workflow == "" {
		logger.Error("Error trying get env", errors.New("env not found"), zap.String("journey", "CheckImage"))
		return rest_err.NewInternalServerError("server error")
	}
	file, err := fileHeader.Open()
	if err != nil {
		logger.Error("Error trying open file", err, zap.String("journey", "CheckImage"))
		return rest_err.NewInternalServerError("server error")
	}
	defer file.Close()
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	part, err := writer.CreateFormFile("media", fileName)
	if err != nil {
		logger.Error("Error trying write field", err, zap.String("journey", "CheckImage"))
		return rest_err.NewInternalServerError("server error")
	}
	_, err = io.Copy(part, file)
	if err != nil {
		logger.Error("Error trying copy file", err, zap.String("journey", "CheckImage"))
		return rest_err.NewInternalServerError("server error")
	}
	err = writer.WriteField("workflow", workflow)
	if err != nil {
		logger.Error("Error trying write field", err, zap.String("journey", "CheckImage"))
		return rest_err.NewInternalServerError("server error")
	}
	err = writer.WriteField("api_user", user)
	if err != nil {
		logger.Error("Error trying write field", err, zap.String("journey", "CheckImage"))
		return rest_err.NewInternalServerError("server error")
	}
	err = writer.WriteField("api_secret", secret)
	if err != nil {
		logger.Error("Error trying write field", err, zap.String("journey", "CheckImage"))
		return rest_err.NewInternalServerError("server error")
	}
	err = writer.Close()
	if err != nil {
		logger.Error("Error trying close writer", err, zap.String("journey", "CheckImage"))
		return rest_err.NewInternalServerError("server error")
	}

	req, err := http.NewRequest("POST", "https://api.sightengine.com/1.0/check-workflow.json", &requestBody)
	if err != nil {
		logger.Error("Error trying make request", err, zap.String("journey", "CheckImage"))
		return rest_err.NewInternalServerError("server error")
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error("Error trying do request", err, zap.String("journey", "CheckImage"))
		return rest_err.NewInternalServerError("server error")
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error trying read response body", err, zap.String("journey", "CheckImage"))
		return rest_err.NewInternalServerError("server error")
	}

	var respJson Response
	if err := json.Unmarshal(respBody, &respJson); err != nil {
		logger.Error("Error trying unmarshal response", err, zap.String("journey", "CheckImage"))
		return rest_err.NewInternalServerError("server error")
	}
	
	if respJson.Summary.Action == "reject" {
		var reasons string
		if respJson.Summary.RejectReason != nil {
			for _, reason := range respJson.Summary.RejectReason {
				reasons += fmt.Sprintf("%s: %s | ", reason.Id, reason.Text)
			}
		}
		logger.Error("Image rejected by Sightengine", errors.New(reasons), zap.String("journey", "CheckImage"))
		return rest_err.NewBadRequestError("innappropriate image content detected")
	}

	if respJson.Status != "success" {
		logger.Error("Error in response from Sightengine", errors.New(respJson.Status), zap.String("journey", "CheckImage"))
		return rest_err.NewInternalServerError("server error")
	}
	fmt.Println(respJson)
	return nil
}
