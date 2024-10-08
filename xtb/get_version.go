package xtb

type GetVersionReturnData struct {
	Version string
}

type GetVersionResponse struct {
	Response
	ReturnData GetVersionReturnData
}

func NewGetVersionCommand() *Command {
	return &Command{
		Command: "getVersion",
	}
}
