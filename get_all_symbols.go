package xtb

type SymbolRecord struct {
	Ask                float64 `json:"ask"`                // Ask price in base currency.
	Bid                float64 `json:"bid"`                // Bid price in base currency.
	CategoryName       string  `json:"categoryName"`       // Category name.
	ContractSize       int     `json:"contractSize"`       // Size of 1 lot.
	Currency           string  `json:"currency"`           // Currency.
	CurrencyPair       bool    `json:"currencyPair"`       // Indicates whether the symbol represents a currency pair.
	CurrencyProfit     string  `json:"currencyProfit"`     // The currency of calculated profit.
	Description        string  `json:"description"`        // Description.
	Expiration         *Time   `json:"expiration"`         // Null if not applicable.
	GroupName          string  `json:"groupName"`          // Symbol group name.
	High               float64 `json:"high"`               // The highest price of the day in base currency.
	InitialMargin      int     `json:"initialMargin"`      // nitial margin for 1 lot order, used for profit/margin calculation.
	InstantMaxVolume   int     `json:"instantMaxVolume"`   // Maximum instant volume multiplied by 100 (in lots).
	Leverage           float64 `json:"leverage"`           // Symbol leverage.
	LongOnly           bool    `json:"longOnly"`           // Long only.
	LotMax             float64 `json:"lotMax"`             // Maximum size of trade.
	LotMin             float64 `json:"lotMin"`             // Minimum size of trade.
	LotStep            float64 `json:"lotStep"`            // A value of minimum step by which the size of trade can be changed (within lotMin - lotMax range).
	Low                float64 `json:"low"`                // The lowest price of the day in base currency.
	MarginHedged       int     `json:"marginHedged"`       // Used for profit calculation.
	MarginHedgedStrong bool    `json:"marginHedgedStrong"` // For margin calculation.
	MarginMaintenance  *int    `json:"marginMaintenance"`  // For margin calculation, null if not applicable.
	MarginMode         int     `json:"marginMode"`         // For margin calculation
	Percentage         float64 `json:"percentage"`         // Percentage.
	PipsPrecision      int     `json:"pipsPrecision"`      // int of symbol's pip decimal places.
	Precision          int     `json:"precision"`          // Number of symbol's price decimal places.
	ProfitMode         int     `json:"profitMode"`         // For profit calculation.
	QuoteID            int     `json:"quoteId"`            // Source of price.
	ShortSelling       bool    `json:"shortSelling"`       // Indicates whether short selling is allowed on the instrument.
	SpreadRaw          float64 `json:"spreadRaw"`          // The difference between raw ask and bid prices.
	SpreadTable        float64 `json:"spreadTable"`        // Spread representation.
	Starting           *Time   `json:"starting"`           // Null if not applicable.
	StepRuleID         int     `json:"stepRuleId"`         // Appropriate step rule ID from getStepRules  command response.
	StopsLevel         int     `json:"stopsLevel"`         // Minimal distance (in pips) from the current price where the stopLoss/takeProfit can be set.
	SwapRollover3Days  int     `json:"swap_rollover3days"` // Time when additional swap is accounted for weekend.
	SwapEnable         bool    `json:"swapEnable"`         // Indicates whether swap value is added to position on end of day.
	SwapLong           float64 `json:"swapLong"`           // Swap value for long positions in pips.
	SwapShort          float64 `json:"swapShort"`          // Swap value for short positions in pips.
	SwapType           int     `json:"swapType"`           // Type of swap calculated..
	Symbol             string  `json:"symbol"`             // Symbol name.
	TickSize           float64 `json:"tickSize"`           // Smallest possible price change, used for profit/margin calculation, null if not applicable.
	TickValue          float64 `json:"tickValue"`          // Value of smallest possible price change (in base currency), used for profit/margin calculation, null if not applicable.
	Time               Time    `json:"time"`               // Ask & bid tick time.
	TimeString         string  `json:"timeString"`         // Time in String.
	TrailingEnabled    bool    `json:"trailingEnabled"`    // Indicates whether trailing stop (offset) is applicable to the instrument.
	Type               int     `json:"type"`               // Instrument class number.
}

type GetAllSymbolsResponse struct {
	Response
	ReturnData []*SymbolRecord `json:"returnData"`
}

func newGetAllSymbolsCommand() *Command {
	return &Command{
		Command: "getAllSymbols",
	}
}
