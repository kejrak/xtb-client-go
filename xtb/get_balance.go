package xtb

type StreamingBalanceRecord struct {
	Balance     float64 `json:"balance"`      // Balance in account currency.
	Credit      float64 `json:"credit"`       // Credit.
	Currency    string  `json:"currency"`     // User currency.
	Equity      float64 `json:"equity"`       // Sum of balance and all profits in account currency.
	Margin      float64 `json:"margin"`       // Margin requirements in account currency.
	MarginFree  float64 `json:"margin_free"`  // Free margin in account currency.
	MarginLevel float64 `json:"margin_level"` // Margin level percentage.
}

type StreamGetBalanceResponse struct {
	StreamResponse
	Data StreamingBalanceRecord `json:"data"`
}

func NewSubscribeGetBalanceCommand(streamSessionID string) *StreamCommand {
	return &StreamCommand{
		Command:         "getBalance",
		StreamSessionID: streamSessionID,
	}
}
