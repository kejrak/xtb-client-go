package xtb

type getTradesRequest struct {
	OpenedOnly bool `json:"openedOnly"` // If true then only opened trades will be returned.
}

type GetTradesResponse struct {
	Response
	ReturnData []TradeRecord
}

func NewGetTradesCommand(openedOnly bool) *Command {
	return &Command{
		Command: "getTrades",
		Arguments: &getTradesRequest{
			OpenedOnly: openedOnly,
		},
	}
}
