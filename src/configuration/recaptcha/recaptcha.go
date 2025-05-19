package recaptcha

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	"go.uber.org/zap"
)

var (
	urlRecaptcha = "https://www.google.com/recaptcha/api/siteverify"
)

type Response struct {
	Success    bool  `json:"success"`
	ErrorCodes []any `json:"error-codes"`
}

func (rc *recaptcha) ValidateRecaptcha(token string) *rest_err.RestErr {
	logger.Info("Init ValidateRecaptcha", zap.String("journey", "ValidateRecaptcha"))
	if os.Getenv("TEST") == "TRUE" {
		return nil
	}
	key := os.Getenv("RECAPTCHA_KEY")
	if key == "" {
		logger.Error("Error trying Get env", errors.New("error trying recaptcha env"), zap.String("journey", "ValidateRecaptcha"))
		return rest_err.NewInternalServerError("server error")
	}
	data := url.Values{
		"secret":   {key},
		"response": {token},
	}
	req, err := http.NewRequest("POST", urlRecaptcha, bytes.NewBufferString(data.Encode()))
	if err != nil {
		logger.Error("Error trying to create request", err, zap.String("journey", "ValidateRecaptcha"))
		return rest_err.NewInternalServerError("server error")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error trying make request", err, zap.String("journey", "ValidateRecaptcha"))
		return rest_err.NewInternalServerError("server error")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error trying read response", err, zap.String("journey", "ValidateRecaptcha"))
		return rest_err.NewInternalServerError("server error")
	}
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		logger.Error("Error trying Unmarshal json", err, zap.String("journey", "ValidateRecaptcha"))
		return rest_err.NewInternalServerError("server error")
	}
	if resp.StatusCode != http.StatusOK {
		logger.Error("invalid recaptcha", fmt.Errorf("error-codes: %v", response.ErrorCodes), zap.String("journey", "ValidateRecaptcha"))
		return rest_err.NewBadRequestError("server error")
	}
	if !response.Success {
		logger.Error("invalid recaptcha", fmt.Errorf("error-codes: %v", response.ErrorCodes), zap.String("journey", "ValidateRecaptcha"))
		return rest_err.NewBadRequestError("recaptcha inválido")
	}
	return nil
}
