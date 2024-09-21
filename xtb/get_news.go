package xtb

import "time"

type getNewsRequest struct {
	End   Time `json:"end"`   // Time, 0 means current time for simplicity.
	Start Time `json:"start"` // Time.
}

type NewsTopicRecord struct {
	Body       string `json:"body"`       // Body.
	BodyLen    int    `json:"bodylen"`    // Body lenght.
	Key        string `json:"key"`        // News key.
	Time       Time   `json:"time"`       // Time.
	TimeString string `json:"timeString"` // Time string.
	Title      string `json:"title"`      // News title.
}

type StreamingNewsRecord struct {
	Body  string `json:"body"`  // Body.
	Key   string `json:"key"`   // News key.
	Time  Time   `json:"time"`  // Time.
	Title string `json:"title"` // News title.
}

type GetNewsResponse struct {
	Response
	ReturnData []NewsTopicRecord `json:"returnData"`
}

type StreamGetNewsResponse struct {
	StreamResponse
	Data StreamingNewsRecord `json:"data"`
}

func NewGetNewsCommand(end, start time.Time) *Command {
	return &Command{
		Command: "getNews",
		Arguments: &getNewsRequest{
			End:   Time(end),
			Start: Time(end),
		},
	}
}

func NewSubscribeGetNewsCommand(streamSessionID string) *StreamCommand {
	return &StreamCommand{
		Command:         "getNews",
		StreamSessionID: streamSessionID,
	}
}
