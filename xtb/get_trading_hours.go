package xtb

type TradingRecord struct {
	Day   uint8 `json:"day"`   // Day of week.
	FromT Time  `json:"fromT"` // Start time in ms from 00:00 CET / CEST time zone (see Daylight Saving Time, DST).
	ToT   Time  `json:"toT"`   // End time in ms from 00:00 CET / CEST time zone (see Daylight Saving Time, DST).
}

type QuotesRecord struct {
	Day   uint8 `json:"day"`   // Day of week.
	FromT Time  `json:"fromT"` // Start time in ms from 00:00 CET / CEST time zone (see Daylight Saving Time, DST).
	ToT   Time  `json:"toT"`   // End time in ms from 00:00 CET / CEST time zone (see Daylight Saving Time, DST).
}

type TradingHoursRecord struct {
	Quotes  []QuotesRecord  `json:"quotes"`  // Array of Quotes Record.
	Symbol  string          `json:"symbol"`  // Symbol
	Trading []TradingRecord `json:"Trading"` // Array of Trading Record.
}

type getTradingHoursRequest struct {
	Symbols []string `json:"symbols"` // Array of symbol names (Strings).
}

type GetTradingHoursResponse struct {
	Response
	ReturnData []TradingHoursRecord `json:"returnData"`
}

func NewGetTradingHoursCommand(symbols []string) *Command {
	return &Command{
		Command: "getTradingHours",
		Arguments: &getTradingHoursRequest{
			Symbols: symbols,
		},
	}
}
