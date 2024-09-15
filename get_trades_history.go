package xtb

import "time"

type getTradesHistoryRequest struct {
	End   Time `json:"end"`   // Time, 0 means current time for simplicity
	Start Time `json:"start"` // Time, 0 means current time for simplicity
}

type GetTradesHistoryResponse struct {
	Response
	ReturnData []TradeRecord `json:"returnData"`
}

func newGetTradesHistoryCommand(end, start time.Time) *Command {
	return &Command{
		Command: "getTradesHistory",
		Arguments: &getTradesHistoryRequest{
			End:   Time(end),
			Start: Time(start),
		},
	}
}
