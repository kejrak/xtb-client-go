package xtb

type getProfitCalculationRequest struct {
	ClosePrice float32      `json:"closePrice"` // Theoretical close price of order.
	Cmd        TradeCommand `json:"cmd"`        // Operation code.
	OpenPrice  float32      `json:"openPrice"`  // Theoretical open price of order.
	Symbol     string       `json:"symbol"`     // Symbol.
	Volume     float32      `json:"volume"`     // Volume.
}

type GetProfitCalculationReturnData struct {
	Profit float64 `json:"profit"` // Profit in account currency.
}

type GetProfitCalculationResponse struct {
	Response
	ReturnData GetProfitCalculationReturnData `json:"returnData"`
}

func NewGetProfitCalculationCommand(closePrice, openPrice, volume float32, cmd TradeCommand, symbol string) *Command {
	return &Command{
		Command: "getProfitCalculation",
		Arguments: &getProfitCalculationRequest{
			ClosePrice: closePrice,
			Cmd:        cmd,
			OpenPrice:  openPrice,
			Symbol:     symbol,
			Volume:     volume,
		},
	}
}
