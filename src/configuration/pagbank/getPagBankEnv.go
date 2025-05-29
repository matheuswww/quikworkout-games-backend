package pagbank

import (
	"errors"
	"os"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"go.uber.org/zap"
)

var (
	apiPathTest    = "https://sandbox.api.pagseguro.com"
	apiPathSdkTest = "https://sandbox.sdk.pagseguro.com"
)

func GetPagbankEnv() (string, string, string) {
	var apiPath, apiPathSdk, token string
	envMode := os.Getenv("ENV_MODE")
	token = os.Getenv("PAGBANK_TOKEN")
	if envMode == "DEV" {
		apiPath = apiPathTest
		apiPathSdk = apiPathSdkTest
	} else if envMode == "PROD" {
		apiPath = "https://api.pagseguro.com"
		apiPathSdk = "https://sdk.pagseguro.com"
	}
	if envMode == "" || apiPath == "" || apiPathSdk == "" || token == ""  {
		logger.Error("Error trying get env", errors.New("was not possible get env"), zap.String("journey", "init pagbank_payment"))
		return "", "", ""
	}
	return token, apiPath, apiPathSdk
}
