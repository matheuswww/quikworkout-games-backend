package participant_repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	"go.uber.org/zap"
)

type PaymentInfos struct {
	OrderId 			string
	PaymentMethod string
}

func (er *participantRepository) HasTicket(cookieId string) ([]PaymentInfos, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var clothingId string
	query := "SELECT clothing_id FROM edition ORDER BY created_at DESC LIMIT 1"
	err := er.mysql.QueryRowContext(ctx, query).Scan(&clothingId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, rest_err.NewBadRequestError("no edition found")
		}
		logger.Error("Error trying QueryRowContext", err, zap.String("journey", "HasTicket Repository"))
		return nil, rest_err.NewInternalServerError("server error")
	}
	
	query = "SELECT ci.order_id_api, co.payment_method FROM clothing_order_info AS ci JOIN client_order AS co ON ci.order_id_api = co.order_id_api WHERE ci.user_id = ? AND ci.clothing_id = ?"
	rows, err := er.mysql.QueryContext(ctx, query, cookieId, clothingId)
	if err != nil {
		logger.Error("Error trying QueryRowContext", err, zap.String("journey", "HasTicket Repository"))
		return nil, rest_err.NewInternalServerError("server error")
	}
	defer rows.Close()

	var paymentInfos []PaymentInfos
	for rows.Next() {
		var orderId, paymentMethod string
		if err := rows.Scan(&orderId, &paymentMethod); err != nil {
			logger.Error("Error trying scan row", err, zap.String("journey", "HasTicket Repository"))
			return nil, rest_err.NewInternalServerError("server error")
		}
		paymentInfos = append(paymentInfos, PaymentInfos{
			OrderId:        orderId,
			PaymentMethod: paymentMethod,
		})
	}

	if len(paymentInfos) == 0 {
		return nil, rest_err.NewNotFoundError("no ticket found")
	}

	return paymentInfos, nil
}