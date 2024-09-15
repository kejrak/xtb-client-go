package xtb

type getMarginTradeRequest struct {
	Symbol string  `json:"symbol"` // Symbol.
	Volume float64 `json:"volume"` // Volume.
}

type GetMarginTradeReturnData struct {
	Margin float64 `json:"margin"` // Calculated margin in account currency.
}

type GetMarginTradeResponse struct {
	Response
	ReturnData GetMarginTradeReturnData `json:"returnData"`
}

func newGetMarginTradeCommand(symbol string, volume float64) *Command {
	return &Command{
		Command: "getMarginTrade",
		Arguments: &getMarginTradeRequest{
			Symbol: symbol,
			Volume: volume,
		},
	}
}
