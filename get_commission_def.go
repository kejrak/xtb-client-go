package xtb

type getCommissionDefRequest struct {
	Symbol string  `json:"symbol"` // Symbol
	Volume float64 `json:"volume"` // Volume
}

type GetCommissionDefReturnData struct {
	Comission      *float64 `json:"commission"`     // Calculated commission in account currency, could be null if not applicable
	RateOfExchange *float64 `json:"rateOfExchange"` // Rate of exchange between account currency and instrument base currency, could be null if not applicable
}

type GetCommissionDefResponse struct {
	Response
	ReturnData GetCommissionDefReturnData `json:"returnData"`
}

func newGetCommissionDefCommand(symbol string, volume float64) *Command {
	return &Command{
		Command: "getCommissionDef",
		Arguments: &getCommissionDefRequest{
			Symbol: symbol,
			Volume: volume,
		},
	}
}
