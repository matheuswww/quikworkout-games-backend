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

func (er *participantRepository) HasTicket(user_id string) ([]PaymentInfos, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var clothingId, editionID string
	query := "SELECT edition_id, clothing_id FROM edition ORDER BY created_at DESC LIMIT 1"
	err := er.mysql.QueryRowContext(ctx, query).Scan(&editionID, &clothingId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, rest_err.NewBadRequestError("no edition found")
		}
		logger.Error("Error trying get edition", err, zap.String("journey", "HasTicket Repository"))
		return nil, rest_err.NewInternalServerError("server error")
	}
	
	query = "SELECT ci.order_id_api, co.payment_method FROM clothing_order_info AS ci JOIN client_order AS co ON ci.order_id_api = co.order_id_api WHERE ci.user_id = ? AND ci.clothing_id = ?"
	rows, err := er.mysql.QueryContext(ctx, query, user_id, clothingId)
	if err != nil {
		logger.Error("Error trying get clothing_order_info", err, zap.String("journey", "HasTicket Repository"))
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
		query = "SELECT 1 FROM direct_ticket WHERE user_id = ? AND edition_id = ? LIMIT 1"
		var exists bool
		err = er.mysql.QueryRowContext(ctx, query, user_id, editionID).Scan(&exists)
		if err != nil && err != sql.ErrNoRows {
			logger.Error("Error trying get direcit_ticket", err, zap.String("journey", "HasTicket Repository"))
			return nil, rest_err.NewInternalServerError("server error")
		}
		if exists {
			return nil, nil
		}
		return nil, rest_err.NewNotFoundError("no ticket found")
	}

	return paymentInfos, nil
}