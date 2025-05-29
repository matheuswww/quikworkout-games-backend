package pagbank_payment

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/pagbank"
	pagbank_struct_payment "github.com/matheuswww/quikworkout-games-backend/src/configuration/pagbank/payment/struct"
	pagbank_payment_util "github.com/matheuswww/quikworkout-games-backend/src/configuration/pagbank/payment/util"
	"go.uber.org/zap"
)

var (
	getOrderPath       = "orders"
	unknow             = "desconhecido"
	UnsolicitedPayment = "pagamento não solicitado"
)

func GetOrder(payment_method, order_id string) string {
	journey := "GetOrder"
	logger.Info("Init GetOrder", zap.String("journey", journey))
	token, apiPath, _ := pagbank.GetPagbankEnv()
	if token == "" || apiPath == "" {
		return unknow
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", apiPath, getOrderPath, order_id), nil)
	if err != nil {
		logger.Error("Error trying make request", err, zap.String("journey", journey))
		return unknow
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error("Error trying make request", err, zap.String("journey", journey))
		return unknow
	}
	bodyBuffer := new(bytes.Buffer)
	_, err = bodyBuffer.ReadFrom(res.Body)
	if err != nil {
		logger.Error("Error trying read body", err, zap.String("journey", journey))
		return unknow
	}
	bodyReader := bytes.NewReader(bodyBuffer.Bytes())

	var getOrderResponse pagbank_struct_payment.GetOrderResponseStatus
	if err := json.NewDecoder(bodyReader).Decode(&getOrderResponse); err != nil || getOrderResponse.Id == "" {
		logger.Error("Error trying decode response", err, zap.String("journey", "GetCharge"))
		pagbank_payment_util.LogResponse(bodyBuffer, journey)
		return unknow
	}
	lastCharge := len(getOrderResponse.ChargesResponse) - 1
	if payment_method == "PIX" && len(getOrderResponse.ChargesResponse) == 0 && getOrderResponse.Id != "" {
		return UnsolicitedPayment
	}
	if len(getOrderResponse.ChargesResponse) == 0 {
		if getOrderResponse.Id != "" {
			return UnsolicitedPayment
		}
		pagbank_payment_util.LogResponse(bodyBuffer, journey)
		logger.Error("Error trying get charge response", errors.New("error getting charge response"), zap.String("journey", "GetCharge"))
		return unknow
	}
	if getOrderResponse.ChargesResponse[lastCharge].Status == "" {
		pagbank_payment_util.LogResponse(bodyBuffer, journey)
		logger.Error("Error trying get charge response", errors.New("error getting charge response"), zap.String("journey", "GetCharge"))
		return unknow
	}
	if payment_method == "CREDIT_CARD" || payment_method == "DEBIT_CARD" {
		payment_method = pagbank_payment_util.Card
	} else if payment_method == "BOLETO" {
		payment_method = pagbank_payment_util.Boleto
	} else if payment_method == "PIX" {
		payment_method = pagbank_payment_util.Pix
	}
	status := pagbank_payment_util.HandlePaymentStatus(payment_method, getOrderResponse.ChargesResponse[lastCharge].Status)
	return status
}
