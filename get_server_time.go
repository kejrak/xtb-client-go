package xtb

type GetServerTimeReturnData struct {
	Time       Time   `json:"time"`       // Time.
	TimeString string `json:"timeString"` // Time described in form set on server (local time of server).
}

type GetServerTimeResponse struct {
	Response
	ReturnData GetServerTimeReturnData `json:"returnData"`
}

func newGetServerTimeCommand() *Command {
	return &Command{
		Command: "getServerTime",
	}
}
