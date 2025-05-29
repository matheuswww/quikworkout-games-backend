package pagbank_payment_util

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"go.uber.org/zap"
)

func LogResponse(bodyBuffer *bytes.Buffer, journey string) {
	bodyRes, err := ioutil.ReadAll(bodyBuffer)
	if err != nil {
		logger.Error("Error trying read body response", err, zap.String("journey", "GetOrder"))
	} else {
		logger.Error("Error trying decode response", errors.New(fmt.Sprintf("body: %s", string(bodyRes))))
	}
}
