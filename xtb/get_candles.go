package xtb

type StreamingCandleRecord struct {
	Close     float64 `json:"close"`     // Value of close price
	Ctm       Time    `json:"ctm"`       // Candle start time in CET / CEST time zone (see Daylight Saving Time, DST)
	CtmString string  `json:"ctmString"` // String representation of the 'ctm' field
	High      float64 `json:"high"`      // Highest value in the given period
	Low       float64 `json:"low"`       // Lowest value in the given period
	Open      float64 `json:"open"`      // Open price
	QuoteID   int     `json:"quoteId"`   // Source of price.
	Symbol    string  `json:"symbol"`    // Symbol.
	Vol       float64 `json:"vol"`       // Volume in lots
}

type StreamGetCandlesResponse struct {
	StreamResponse
	Data StreamingCandleRecord `json:"data"`
}

func NewSubscribeGetCandlesCommand(streamSessionID, symbol string) *StreamCandleCommand {
	return &StreamCandleCommand{
		Command:         "getCandles",
		StreamSessionID: streamSessionID,
		Symbol:          symbol,
	}
}
