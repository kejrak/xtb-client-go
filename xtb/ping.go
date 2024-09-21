package xtb

type PingResponse struct {
	Response
}

func NewPingCommand() *Command {
	return &Command{
		Command: "ping",
	}
}

func NewSubscribePingCommand(streamSessionID string) *StreamCommand {
	return &StreamCommand{
		Command:         "ping",
		StreamSessionID: streamSessionID,
	}
}
