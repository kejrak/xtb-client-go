package xtb

type GetMarginLevelReturnData struct {
	Balance     float64 `json:"balance"`      // Balance in account currency.
	Credit      float64 `json:"credit"`       // Credit.
	Currency    string  `json:"currency"`     // User currency.
	Equity      float64 `json:"equity"`       // Sum of balance and all profits in account currency.
	Margin      float64 `json:"margin"`       // Margin requirements in account currency.
	MarginFree  float64 `json:"margin_free"`  // Free margin in account currency.
	MarginLevel float64 `json:"margin_level"` // Margin level percentage.
}

type GetMarginLevelResponse struct {
	Response
	ReturnData GetMarginLevelReturnData `json:"returnData"`
}

func NewGetMarginLevelCommand() *Command {
	return &Command{
		Command: "getMarginLevel",
	}
}
