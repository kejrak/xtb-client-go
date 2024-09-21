package xtb

type StreamingProfitRecord struct {
	Order    uint    // Order number.
	Order2   uint    // Transaction ID.
	Position uint    // Position number.
	Profit   float64 // Profit in account currency.
}

type StreamGetProfitsResponse struct {
	StreamResponse
	Data StreamingProfitRecord `json:"data"`
}

func NewSubscribeGetProfitsCommand(streamSessionID string) *StreamCommand {
	return &StreamCommand{
		Command:         "getProfits",
		StreamSessionID: streamSessionID,
	}
}
