package xtb

type tradeTransactionStatusRequest struct {
	Order uint `json:"order"` // Order.
}

type TradeTransactionStatusReturnData struct {
	Ask           float64 `json:"ask"`           // Price in base currency.
	Bid           float64 `json:"bid"`           // Price in base currency.
	CustomComment string  `json:"customComment"` // The value the customer may provide in order to retrieve it later.
	Message       *string `json:"message"`       // Can be null.
	Order         uint    `json:"order"`         // Unique order number.
	RequestStatus uint8   `json:"requestStatus"` // Request status code, described below
}

type TradeTransactionStatusResponse struct {
	Response
	ReturnData TradeTransactionStatusReturnData `json:"returnData"`
}

func NewTradeTransactionStatusCommand(order uint) *Command {
	return &Command{
		Command: "tradeTransactionStatus",
		Arguments: &tradeTransactionStatusRequest{
			Order: order,
		},
	}
}
