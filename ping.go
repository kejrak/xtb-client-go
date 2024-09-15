package xtb

type PingResponse struct {
	Response
}

func newPingCommand() *Command {
	return &Command{
		Command: "ping",
	}
}
