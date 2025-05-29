package participant_service

import (

	pagbank_payment "github.com/matheuswww/quikworkout-games-backend/src/configuration/pagbank/payment"
	pagbank_payment_util "github.com/matheuswww/quikworkout-games-backend/src/configuration/pagbank/payment/util"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
)

func (es *participantService) HasTicket(cookieId string) *rest_err.RestErr {
	paymentInfos, restErr := es.participantRepository.HasTicket(cookieId)
	if restErr != nil {
		return restErr
	}
	var msg string
	var found bool
	var error string
	for _, paymentInfo := range paymentInfos {
		status := pagbank_payment.GetOrder(paymentInfo.PaymentMethod, paymentInfo.OrderId)

		if status == pagbank_payment_util.Pago {
			found = true
			break
		}
		if status == pagbank_payment_util.EmAnalise || status == pagbank_payment_util.Autorizado {
			msg = "payment is still being processed, please wait a few minutes and try again"
		}
		if status == pagbank_payment_util.Desconhecido {
			error = "unknown payment status"
		}
	}
	if !found {
		if msg != "" {
			return rest_err.NewBadRequestError(msg)
		}
		if error != "" {
			return rest_err.NewInternalServerError(error)
		}
		return rest_err.NewNotFoundError("no ticket found")
	}
	return nil
}