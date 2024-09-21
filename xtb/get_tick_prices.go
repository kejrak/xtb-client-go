package xtb

import "time"

type TickRecord struct {
	Ask         float64 `json:"ask"`         // Ask price in base currency.
	AskVolume   int     `json:"askVolume"`   // Number of available lots to buy at given price or null if not applicable.
	Bid         float64 `json:"bid"`         // Bid price in base currency.
	BidVolume   int     `json:"bidVolume"`   // Number of available lots to buy at given price or null if not applicable.
	High        float64 `json:"high"`        // The highest price of the day in base currency.
	Level       int     `json:"level"`       // Price level.
	Low         float64 `json:"low"`         // The lowest price of the day in base currency.
	SpreadRaw   float64 `json:"spreadRaw"`   // The difference between raw ask and bid prices.
	SpreadTable float64 `json:"spreadTable"` // Spread representation.
	Symbol      string  `json:"symbol"`      // Symbol.
	Timestamp   Time    `json:"timestamp"`   // Timestamp.
}

type StreamingTickRecord struct {
	Ask         float64 `json:"ask"`         // Ask price in base currency.
	AskVolume   int     `json:"askVolume"`   // Number of available lots to buy at given price or null if not applicable.
	Bid         float64 `json:"bid"`         // Bid price in base currency.
	BidVolume   int     `json:"bidVolume"`   // Number of available lots to buy at given price or null if not applicable.
	High        float64 `json:"high"`        // The highest price of the day in base currency.
	Level       int     `json:"level"`       // Price level.
	Low         float64 `json:"low"`         // The lowest price of the day in base currency.
	QuoteID     int     `json:"quoteId"`     // Source of price.
	SpreadRaw   float64 `json:"spreadRaw"`   // The difference between raw ask and bid prices.
	SpreadTable float64 `json:"spreadTable"` // Spread representation.
	Symbol      string  `json:"symbol"`      // Symbol.
	Timestamp   Time    `json:"timestamp"`   // Timestamp.
}

type getTickPricesRequest struct {
	Level     int      `json:"level"`     // Price level.
	Symbols   []string `json:"symbols"`   // Array of symbol names (String)
	Timestamp Time     `json:"timestamp"` // The time from which the most recent tick should be looked for. Historical prices cannot be obtained using this parameter. It can only be used to verify whether a price has changed since the given time.
}

type GetTickPricesReturnData struct {
	Quotations []TickRecord `json:"quotations"` // Array of TickRecord
}

type GetTickPricesResponse struct {
	Response
	ReturnData GetTickPricesReturnData `json:"returnData"`
}

type StreamGetTickPricesResponse struct {
	StreamResponse
	Data StreamingTickRecord `json:"data"`
}

func NewGetTickPricesCommand(level int, symbols []string, timestamp time.Time) *Command {
	return &Command{
		Command: "getTickPrices",
		Arguments: &getTickPricesRequest{
			Level:     level,
			Symbols:   symbols,
			Timestamp: Time(timestamp),
		},
	}
}

func NewSubscribeGetTickPricesCommand(streamSessionID, symbol string, minArrivalTime *int, maxLevel *uint8) *StreamTickCommand {
	return &StreamTickCommand{
		Command:         "getTickPrices",
		StreamSessionID: streamSessionID,
		Symbol:          symbol,
		MinArrivalTime:  minArrivalTime,
		MaxLevel:        maxLevel,
	}
}
