package pagbank_payment_util

var (
	in_analysis = "IN_ANALYSIS"
	authorized  = "AUTHORIZED"
	declined    = "DECLINED"
	Canceled    = "CANCELED"
	waiting     = "WAITING"

	Cancelado    = "cancelado"
	Recusado     = "recusado"
	EmAnalise    = "em análise"
	Aguardando   = "aguardando"
	Autorizado   = "autorizado"
	Pago         = "pago"
	Desconhecido = "desconhecido"

	Paid   = "PAID"
	Card   = "CARD"
	Boleto = "BOLETO"
	Pix    = "PIX"
)

func HandlePaymentStatus(paymentType, status string) string {
	if paymentType == Card {
		if status == in_analysis {
			return EmAnalise
		}
		if status == authorized {
			return Autorizado
		}
		if status == declined {
			return Recusado
		}
	} else if paymentType == Boleto {
		if status == waiting {
			return Aguardando
		}
	} else if paymentType == Pix {
		if status == authorized {
			return Autorizado
		}
	}
	if status == Paid {
		return Pago
	}
	if status == Canceled {
		return Cancelado
	}
	return Desconhecido
}
