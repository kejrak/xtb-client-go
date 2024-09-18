package xtb

type GetCurrentUserDataReturnData struct {
	CompanyUnit        int     `json:"companyUnit"`        // Unit the account is assigned to.
	Currency           string  `json:"currency"`           // Account currency.
	Group              string  `json:"group"`              // Group.
	IBAccount          bool    `json:"ibAccount"`          // Indicates whether this account is an IB account.
	Leverage           int     `json:"leverage"`           // This field should not be used. It is inactive and its value is always 1.
	LeverageMultiplier float64 `json:"LeverageMultiplier"` // The factor used for margin calculations. The actual value of leverage can be calculated by dividing this value by 100.
	SpreadType         string  `json:"spreadType"`         // SpreadType, null if not applicable.
	TrailingStop       bool    `json:"trailingStop"`       // Indicates whether this account is enabled to use trailing stop.
}

type GetCurrentUserDataResponse struct {
	Response
	ReturnData GetCurrentUserDataReturnData `json:"returnData"`
}

func NewGetCurrentUserDataCommand() *Command {
	return &Command{
		Command: "getCurrentUserData",
	}
}
