package xtb

import "time"

type ChartRangeInfoRecord struct {
	End    Time   `json:"end"`    // End of chart block (rounded down to the nearest interval and excluding).
	Period Period `json:"period"` // Period code.
	Start  Time   `json:"start"`  // Symbol.
	Symbol string `json:"symbol"` // Highest value in the given period (shift from open price).
	Ticks  *int   `json:"ticks"`  // Number of ticks needed, this field is optional, please read the description above.
}

type getChartRangeRequest struct {
	Info ChartRangeInfoRecord `json:"info"`
}

type GetChartRangeRequestReturnData struct {
	Digits    int
	RateInfos []RateInfoRecord
}

type GetChartRangeRequestResponse struct {
	Response
	ReturnData GetChartRangeRequestReturnData
}

func NewGetChartRangeRequestCommand(end, start time.Time, period Period, symbol string, ticks *int) *Command {
	return &Command{
		Command: "getChartRangeRequest",
		Arguments: &getChartRangeRequest{
			Info: ChartRangeInfoRecord{
				End:    Time(end),
				Period: period,
				Start:  Time(start),
				Symbol: symbol,
				Ticks:  ticks,
			},
		},
	}
}
