package xtb

import "time"

type ChartLastInfoRecord struct {
	Period Period `json:"period"` // Period code.
	Start  Time   `json:"start"`  // Start of chart block (rounded down to the nearest interval and excluding).
	Symbol string `json:"symbol"` // Symbol.
}

type getChartLastRequest struct {
	Info ChartLastInfoRecord `json:"info"`
}

type GetChartLastRequestReturnData struct {
	Digits    int
	RateInfos []RateInfoRecord
}

type GetChartLastRequestResponse struct {
	Response
	ReturnData GetChartLastRequestReturnData
}

func newGetChartLastRequestCommand(period Period, start time.Time, symbol string) *Command {
	return &Command{
		Command: "getChartLastRequest",
		Arguments: &getChartLastRequest{
			Info: ChartLastInfoRecord{
				Period: period,
				Start:  Time(start),
				Symbol: symbol,
			},
		},
	}
}
