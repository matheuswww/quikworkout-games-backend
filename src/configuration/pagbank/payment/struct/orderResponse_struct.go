package pagbank_struct_payment

type GetOrderResponseStatus struct {
	Id              string                  `json:"id"`
	ChargesResponse []ChargesResponseStatus `json:"charges"`
}

type ChargesResponseStatus struct {
	Status string `json:"status"`
}
