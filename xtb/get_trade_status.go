package xtb

type StreamingTradeStatusRecord struct {
	CustomComment string  `json:"customComment"` // The value the customer may provide in order to retrieve it later.
	Message       *string `json:"message"`       // Can be null.
	Order         int     `json:"order"`         // Unique order number.
	Price         float64 `json:"price"`         // Price in base currency.
	RequestStatus int     `json:"requestStatus"` // Request status code.
}

type StreamGetTradeStatus struct {
	StreamResponse
	Data StreamingTradeStatusRecord `json:"data"`
}

func NewSubscribeGetTradeStatusCommand(streamSessionID string) *StreamCommand {
	return &StreamCommand{
		Command:         "tradeStatus",
		StreamSessionID: streamSessionID,
	}
}
