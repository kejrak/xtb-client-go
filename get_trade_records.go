package xtb

type TradeRecord struct {
	ClosePrice       float64  `json:"close_price"`      // Close price in base currency.
	CloseTime        *Time    `json:"close_time"`       // Null if order is not closed.
	CloseTimeString  *string  `json:"close_timeString"` // Null if order is not closed.
	Closed           bool     `json:"closed"`           // Closed.
	Cmd              uint8    `json:"cmd"`              // Operation code.
	Comment          string   `json:"comment"`          // Comment.
	Commission       *float64 `json:"commission"`       // Commission in account currency, null if not applicable.
	CustomComment    string   `json:"customComment"`    // The value the customer may provide in order to retrieve it later.
	Digits           int      `json:"digits"`           // Number of decimal places.
	Expiration       *Time    `json:"expiration"`       // Null if order is not closed.
	ExpirationString *string  `json:"expirationString"` // Null if order is not closed.
	MarginRate       float64  `json:"margin_rate"`      // Margin Rate.
	Offset           int      `json:"offset"`           // Trailing offset.
	OpenPrice        float64  `json:"open_price"`       // Open price in base currency.
	OpenTime         Time     `json:"open_time"`        // Open time.
	OpenTimeString   string   `json:"open_timeString"`  // Open time string.
	Order            uint     `json:"order"`            // Order number for opened transaction.
	Order2           uint     `json:"order2"`           // Order number for closed transaction.
	Position         int      `json:"position"`         // Order number common both for opened and closed transaction.
	Profit           float64  `json:"profit"`           // Profit in account currency.
	Sl               float64  `json:"sl"`               // Zero if stop loss is not set (in base currency).
	Storage          float64  `json:"storage"`          // Order swaps in account currency.
	Symbol           string   `json:"symbol"`           // Symbol name or null for deposit/withdrawal operations.
	Timestamp        Time     `json:"timestamp"`        // Timestamp
	Tp               float64  `json:"tp"`               // Zero if take profit is not set (in base currency).
	Volume           float64  `json:"volume"`           // Volume in lots
}

type getTradeRecordsRequest struct {
	Orders []int // Array of orders (pistion numbers),
}

type GetTradeRecordsResponse struct {
	Response
	ReturnData []TradeRecord `json:"returnData"`
}

func newGetTradeRecordsCommand(orders []int) *Command {
	return &Command{
		Command: "getTradeRecords",
		Arguments: &getTradeRecordsRequest{
			Orders: orders,
		},
	}
}
