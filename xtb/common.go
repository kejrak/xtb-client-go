package xtb

import (
	"fmt"
	"strconv"
	"time"
)

type Period int
type TradeCommand uint8
type TradeType uint8
type RequestStatus uint8

const (
	PERIOD_M1  Period = 1     // 1 minute
	PERIOD_M5  Period = 5     // 5 minutes
	PERIOD_M15 Period = 15    // 15 minutes
	PERIOD_M30 Period = 30    // 30 minutes
	PERIOD_H1  Period = 60    // 60 minutes (1 hour)
	PERIOD_H4  Period = 240   // 240 minutes (4 hours)
	PERIOD_D1  Period = 1440  // 1440 minutes (1 day)
	PERIOD_W1  Period = 10080 // 10080 minutes (1 week)
	PERIOD_MN1 Period = 43200 // 43200 minutes (30 days / 1 month)
)

const (
	OPEN TradeType = iota
	PENDING_TRADE
	CLOSE
	MODIFY
	DELETE
)

const (
	ERROR RequestStatus = iota
	PENDING_REQUEST
	ACCEPTED
	REJECTED
)

const (
	BUY        TradeCommand = 0 // Buy
	SELL       TradeCommand = 1 // Sell
	BUY_LIMIT  TradeCommand = 2 // Buy limit
	SELL_LIMIT TradeCommand = 3 // Sell limit
	BUY_STOP   TradeCommand = 4 // Buy stop
	SELL_STOP  TradeCommand = 5 // Sell stop
	BALANCE    TradeCommand = 6 // Read only. Used in getTradesHistory (http://developers.xstore.pro/documentation/#getTradesHistory) for manager's deposit/withdrawal operations (profit>0 for deposit, profit<0 for withdrawal).
	CREDIT     TradeCommand = 7 // Read only
)

type SymbolRecord struct {
	Ask                float64    `json:"ask"`                // Ask price in base currency.
	Bid                float64    `json:"bid"`                // Bid price in base currency.
	CategoryName       string     `json:"categoryName"`       // Category name.
	ContractSize       int        `json:"contractSize"`       // Size of 1 lot.
	Currency           string     `json:"currency"`           // Currency.
	CurrencyPair       bool       `json:"currencyPair"`       // Indicates whether the symbol represents a currency pair.
	CurrencyProfit     string     `json:"currencyProfit"`     // The currency of calculated profit.
	Description        string     `json:"description"`        // Description.
	Expiration         *time.Time `json:"expiration"`         // Null if not applicable.
	GroupName          string     `json:"groupName"`          // Symbol group name.
	High               float64    `json:"high"`               // The highest price of the day in base currency.
	InitialMargin      int        `json:"initialMargin"`      // nitial margin for 1 lot order, used for profit/margin calculation.
	InstantMaxVolume   int        `json:"instantMaxVolume"`   // Maximum instant volume multiplied by 100 (in lots).
	Leverage           float64    `json:"leverage"`           // Symbol leverage.
	LongOnly           bool       `json:"longOnly"`           // Long only.
	LotMax             float64    `json:"lotMax"`             // Maximum size of trade.
	LotMin             float64    `json:"lotMin"`             // Minimum size of trade.
	LotStep            float64    `json:"lotStep"`            // A value of minimum step by which the size of trade can be changed (within lotMin - lotMax range).
	Low                float64    `json:"low"`                // The lowest price of the day in base currency.
	MarginHedged       int        `json:"marginHedged"`       // Used for profit calculation.
	MarginHedgedStrong bool       `json:"marginHedgedStrong"` // For margin calculation.
	MarginMaintenance  *int       `json:"marginMaintenance"`  // For margin calculation, null if not applicable.
	MarginMode         int        `json:"marginMode"`         // For margin calculation
	Percentage         float64    `json:"percentage"`         // Percentage.
	PipsPrecision      int        `json:"pipsPrecision"`      // int of symbol's pip decimal places.
	Precision          int        `json:"precision"`          // Number of symbol's price decimal places.
	ProfitMode         int        `json:"profitMode"`         // For profit calculation.
	QuoteID            int        `json:"quoteId"`            // Source of price.
	ShortSelling       bool       `json:"shortSelling"`       // Indicates whether short selling is allowed on the instrument.
	SpreadRaw          float64    `json:"spreadRaw"`          // The difference between raw ask and bid prices.
	SpreadTable        float64    `json:"spreadTable"`        // Spread representation.
	Starting           *time.Time `json:"starting"`           // Null if not applicable.
	StepRuleID         int        `json:"stepRuleId"`         // Appropriate step rule ID from getStepRules  command response.
	StopsLevel         int        `json:"stopsLevel"`         // Minimal distance (in pips) from the current price where the stopLoss/takeProfit can be set.
	SwapRollover3Days  int        `json:"swap_rollover3days"` // Time when additional swap is accounted for weekend.
	SwapEnable         bool       `json:"swapEnable"`         // Indicates whether swap value is added to position on end of day.
	SwapLong           float64    `json:"swapLong"`           // Swap value for long positions in pips.
	SwapShort          float64    `json:"swapShort"`          // Swap value for short positions in pips.
	SwapType           int        `json:"swapType"`           // Type of swap calculated..
	Symbol             string     `json:"symbol"`             // Symbol name.
	TickSize           float64    `json:"tickSize"`           // Smallest possible price change, used for profit/margin calculation, null if not applicable.
	TickValue          float64    `json:"tickValue"`          // Value of smallest possible price change (in base currency), used for profit/margin calculation, null if not applicable.
	Time               time.Time  `json:"time"`               // Ask & bid tick time.
	TimeString         string     `json:"timeString"`         // Time in String.
	TrailingEnabled    bool       `json:"trailingEnabled"`    // Indicates whether trailing stop (offset) is applicable to the instrument.
	Type               int        `json:"type"`               // Instrument class number.
}

type RateInfoRecord struct {
	Close     float64   `json:"close"`     // Value of close price (shift from open price)
	Ctm       time.Time `json:"ctm"`       // Candle start time in CET / CEST time zone (see Daylight Saving Time, DST)
	CtmString string    `json:"ctmString"` // String representation of the 'ctm' field
	High      float64   `json:"high"`      // Highest value in the given period (shift from open price)
	Low       float64   `json:"low"`       // Lowest value in the given period (shift from open price)
	Open      float64   `json:"open"`      // Open price (in base currency * 10 to the power of digits)
	Vol       float64   `json:"vol"`       // Volume in lots
}

// Command represents the structure of a command to be sent to the XTB API.
type Command struct {
	Command   string      `json:"command"`             // The type of command.
	Arguments interface{} `json:"arguments,omitempty"` // Optional arguments for the command.
}

// Response represents the structure of a response from the XTB API.
type Response struct {
	Status     bool   `json:"status"`               // Indicates if the command was successful.
	ErrorCode  string `json:"errorCode,omitempty"`  // Error code if the command failed.
	ErrorDescr string `json:"errorDescr,omitempty"` // Description of the error if applicable.
}

// Stream Command represents the structure of a command to be sent to the XTB API.
type StreamResponse struct {
	Command string `json:"command"` // The type of command.
}

// Stream Command represents the structure of a command to be sent to the XTB API.
type StreamCommand struct {
	Command         string `json:"command"`         // The type of command.
	StreamSessionID string `json:"streamSessionId"` // Stream Session ID.
}

// Stream Candle Command represents the structure of a candle command to be sent to the XTB API.
type StreamCandleCommand struct {
	Command         string `json:"command"`         // The type of command.
	StreamSessionID string `json:"streamSessionId"` // Stream Session ID.
	Symbol          string `json:"symbol"`          // Symbol for the command.
}

// Stream Candle Command represents the structure of a candle command to be sent to the XTB API.
type StreamTickCommand struct {
	Command         string `json:"command"`         // The type of command.
	StreamSessionID string `json:"streamSessionId"` // Stream Session ID.
	Symbol          string `json:"symbol"`          // Symbol for the command.
	MinArrivalTime  *int   `json:"minArrivalTime"`  // This field is optional and defines the minimal interval in milliseconds between any two consecutive updates. If this field is not present, or it is set to 0 (zero), ticks - if available - are sent to the client with interval equal to 200 milliseconds. In order to obtain ticks as frequently as server allows you, set it to 1 (one).
	MaxLevel        *uint8 `json:"maxLevel"`        // This field is optional and specifies the maximum level of the quote that the user is interested in. If this field is not specified, the subscription is active for all levels that are managed in the system.
}

// Time is a wrapper around time.Time to handle custom JSON serialization/deserialization.
type Time time.Time

// ToTime converts the custom Time type to the standard time.Time type.
func (t *Time) ToTime() time.Time {
	return time.Time(*t)
}

// UnmarshalJSON customizes the unmarshalling of Time from JSON.
// It converts a JSON timestamp (milliseconds since Unix epoch) to a Time object.
func (t *Time) UnmarshalJSON(text []byte) error {
	timestamp, err := strconv.ParseInt(string(text), 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse timestamp: %w", err)
	}

	*t = Time(time.UnixMilli(timestamp))
	return nil
}

// MarshalJSON customizes the marshalling of Time to JSON.
// It converts a Time object to its JSON representation (timestamp in milliseconds).
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).UnixMilli(), 10)), nil
}
