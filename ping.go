package xtb

type PingResponse struct {
	Response
}

func NewPingCommand() *Command {
	return &Command{
		Command: "ping",
	}
}
