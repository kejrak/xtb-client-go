package xtb

import "time"

type IBRecord struct {
	ClosePrice *float64 `json:"closePrice"` // IB close price or null if not allowed to view.
	Login      *string  `json:"login"`      // IB user login or null if not allowed to view.
	Nominal    *float64 `json:"nominal"`    // IB nominal or null if not allowed to view.
	OpenPrice  *float64 `json:"openPrice"`  // IB open price or null if not allowed to view.
	Side       *int     `json:"side"`       // Operation code or null if not allowed to view. (0 BUY, 1 SELL)
	Surname    *string  `json:"surname"`    // IB user surname or null if not allowed to view.
	Symbol     *string  `json:"symbol"`     // Symbol or null if not allowed to view.
	Timestamp  *Time    `json:"timestamp"`  // Time the record was created or null if not allowed to view
	Volume     *float64 `json:"volume"`     // Volume in lots or null if not allowed to view.
}

type getIbsHistoryRequest struct {
	End   Time `json:"end"`   // End of IB history block.
	Start Time `json:"start"` // Start of IB history block.
}

type GetIbsHistoryResponse struct {
	Response
	ReturnData []IBRecord `json:"returnData"`
}

func newGetIbsHistoryCommand(end, start time.Time) *Command {
	return &Command{
		Command: "getIbsHistory",
		Arguments: &getIbsHistoryRequest{
			End:   Time(end),
			Start: Time(start),
		},
	}
}
