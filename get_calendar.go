package xtb

type CalendarRecord struct {
	Country  string `json:"country"`  // Two letter country code.
	Current  string `json:"current"`  // Market value (current), empty before time of release of this value (time from "time" record).
	Forecast string `json:"forecast"` // Forecasted value.
	Impact   string `json:"impact"`   // Impact on market.
	Period   string `json:"period"`   // Information period.
	Previous string `json:"previous"` // Value rom previous infromation release.
	Time     Time   `json:"time"`     // Time, when the information will be released (in this time empty "current" value should be changed with exact released value).
	Title    string `json:"title"`    // Name of the indicator for which values will be released.
}

type GetCalendarResponse struct {
	Response
	ReturnData []CalendarRecord `json:"returnData"`
}

func NewGetCalendarCommand() *Command {
	return &Command{
		Command: "getCalendar",
	}
}
