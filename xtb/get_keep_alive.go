package xtb

type StreamingKeepAliveRecord struct {
	Timestamp Time `json:"timestamp"` // Timestamp.
}

type StreamGetKeepAliveResponse struct {
	StreamResponse
	Data StreamingKeepAliveRecord `json:"data"`
}

func NewSubscribeGetKeepAliveCommand(streamSessionId string) *StreamCommand {
	return &StreamCommand{
		Command:         "getKeepAlive",
		StreamSessionID: streamSessionId,
	}
}
