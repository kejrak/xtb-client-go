package xtb

type TradeTransInfo struct {
	Cmd           TradeCommand `json:"cmd"`           // Operation code.
	CustomComment string       `json:"customComment"` // The value the customer may provide in order to retrieve it later.
	Expiration    int64        `json:"expiration"`    // Null if order is not closed.
	Offset        int          `json:"offset"`        // Trailing offset.
	Order         uint         `json:"order"`         // Order number for opened transaction.
	Price         float64      `json:"price"`         // Trade price.
	Sl            float64      `json:"sl"`            // Zero if stop loss is not set (in base currency).
	Symbol        string       `json:"symbol"`        // symbol name or null for deposit/withdrawal operations.
	Tp            float64      `json:"tp"`            // Zero if take profit is not set (in base currency).
	Type          TradeType    `json:"type"`          // Trade transaction type.
	Volume        float64      `json:"volume"`        // Volume in lots.
}

type tradeTransactionRequest struct {
	TradeTransInfo TradeTransInfo `json:"tradeTransInfo"` // if true then only opened trades will be returned.
}

type TradeTransactionReturnData struct {
	Order uint `json:"order"` // Order.
}

type TradeTransactionResponse struct {
	Response
	ReturnData TradeTransactionReturnData `json:"returnData"`
}

func newTradeTransactionCommand(cmd TradeCommand, customComment string, expiration int64, offset int, order uint, price float64, sl float64, symbol string, tp float64, tradeType TradeType, volume float64) *Command {
	return &Command{
		Command: "tradeTransaction",
		Arguments: &tradeTransactionRequest{
			TradeTransInfo: TradeTransInfo{
				Cmd:           cmd,
				CustomComment: customComment,
				Expiration:    expiration,
				Offset:        offset,
				Order:         order,
				Price:         price,
				Sl:            sl,
				Symbol:        symbol,
				Tp:            tp,
				Type:          tradeType,
				Volume:        volume,
			},
		},
	}
}
