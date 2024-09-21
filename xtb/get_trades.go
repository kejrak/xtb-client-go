package xtb

type getTradesRequest struct {
	OpenedOnly bool `json:"openedOnly"` // If true then only opened trades will be returned.
}

type GetTradesResponse struct {
	Response
	ReturnData []TradeRecord `json:"returnData"`
}

type StreamGetTradesResponse struct {
	StreamResponse
	Data StreamingTradeRecord `json:"data"`
}

func NewGetTradesCommand(openedOnly bool) *Command {
	return &Command{
		Command: "getTrades",
		Arguments: &getTradesRequest{
			OpenedOnly: openedOnly,
		},
	}
}

func NewSubscribeGetTradesCommand(streamSessionID string) *StreamCommand {
	return &StreamCommand{
		Command:         "getTrades",
		StreamSessionID: streamSessionID,
	}
}
