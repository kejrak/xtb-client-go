package xtb

import (
	"fmt"
	"strconv"
	"time"
)

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
	PENDING
	CLOSE
	MODIFY
	DELETE
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

type RateInfoRecord struct {
	Close     float64 `json:"close"`     // Value of close price (shift from open price)
	Ctm       Time    `json:"ctm"`       // Candle start time in CET / CEST time zone (see Daylight Saving Time, DST)
	CtmString string  `json:"ctmString"` // String representation of the 'ctm' field
	High      float64 `json:"high"`      // Highest value in the given period (shift from open price)
	Low       float64 `json:"low"`       // Lowest value in the given period (shift from open price)
	Open      float64 `json:"open"`      // Open price (in base currency * 10 to the power of digits)
	Vol       float64 `json:"vol"`       // Volume in lots
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
