package xclient

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/kejrak/xtb-client-go/xtb"
)

type StreamMessage struct {
	Command string          `json:"command"`
	Data    json.RawMessage `json:"data"`
}

const (
	websocketBaseURL  = "wss://ws.xtb.com"
	maxViolationCount = 5 // User should send requests in 200 ms intervals. This rule can be broken, but if it happens 6 times in a row the connection is dropped.
)

func (c *Client) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
	if c.streamConn != nil {
		c.streamConn.Close()
	}

	close(c.StreamChannel)
}

type Client struct {
	websocketURL       string
	websocketStreamURL string
	conn               *websocket.Conn
	streamConn         *websocket.Conn
	streamSessionID    string
	lastRequestAt      time.Time
	violationCount     int
	maxViolationCount  int
	StreamChannel      chan StreamMessage
}

func NewClient(connectionType, userID, password string) (*Client, error) {

	var websocketURL string
	var websocketStreamURL string

	switch connectionType {
	case "demo":
		websocketURL = websocketBaseURL + "/demo"
		websocketStreamURL = websocketBaseURL + "/demoStream"
	case "real":
		websocketURL = websocketBaseURL + "/real"
		websocketStreamURL = websocketBaseURL + "/realStream"
	}

	c := &Client{
		websocketURL:       websocketURL,
		websocketStreamURL: websocketStreamURL,
		violationCount:     0,
		maxViolationCount:  maxViolationCount,
	}

	c.StreamChannel = make(chan StreamMessage)

	dialer := websocket.DefaultDialer
	dialer.HandshakeTimeout = 10 * time.Second

	conn, err := c.dialWebSocket(websocketURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to WebSocket: %w", err)
	}

	c.conn = conn

	loginResponse, err := c.Login(userID, password)
	if err != nil || !loginResponse.Status {
		log.Fatalf("Login failed: %v", err)
	}

	c.streamSessionID = loginResponse.StreamSessionID

	streamConn, err := c.dialWebSocket(websocketStreamURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to WebSocket: %w", err)
	}

	c.streamConn = streamConn
	return c, nil
}

func (c *Client) ReadStream() {
	for {
		c.streamConn.SetReadDeadline(time.Now().Add(5 * time.Second))
		_, message, err := c.streamConn.ReadMessage()
		if err != nil {
			log.Printf("Failed to read stream message: %v", err)
			return
		}

		var streamMessage StreamMessage
		err = json.Unmarshal(message, &streamMessage)
		if err != nil {
			log.Printf("Failed to unmarshal stream message: %v", err)
			continue
		}

		c.StreamChannel <- streamMessage
	}
}

func (c *Client) dialWebSocket(urlStr string) (*websocket.Conn, error) {
	dialer := websocket.DefaultDialer
	dialer.HandshakeTimeout = 10 * time.Second

	conn, _, err := dialer.Dial(urlStr, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to WebSocket at %s: %w", urlStr, err)
	}

	return conn, nil
}

func (c *Client) executeCommand(command interface{}, response interface{}) error {
	return c.executeWebSocketCommand(c.conn, command, response)
}

func (c *Client) executeStreamCommand(command interface{}) error {
	return c.executeWebSocketCommand(c.streamConn, command, nil)
}

func (c *Client) executeWebSocketCommand(conn *websocket.Conn, command interface{}, response interface{}) error {
	if err := c.checkRateLimit(); err != nil {
		return err
	}

	commandJSON, err := json.Marshal(command)
	if err != nil {
		log.Print("Failed to marshal command")
		return fmt.Errorf("failed to marshal command: %w", err)
	}

	if len(commandJSON) > 1024 {
		log.Print("Command size exceeds 1KB limit")
		return fmt.Errorf("command size exceeds 1KB limit")
	}

	return c.sendAndReceive(conn, command, response)
}

func (c *Client) checkRateLimit() error {
	now := time.Now()

	if now.Sub(c.lastRequestAt) < 200*time.Millisecond {
		c.violationCount++

		if c.violationCount > c.maxViolationCount {
			log.Print("Exceeded maximum request frequency violations.")
			sleepDuration := 200*time.Millisecond - now.Sub(c.lastRequestAt)
			time.Sleep(sleepDuration)
		}
	} else {
		c.violationCount = 0
	}

	c.lastRequestAt = now
	return nil
}

func (c *Client) sendAndReceive(conn *websocket.Conn, command interface{}, response interface{}) error {
	now := time.Now()

	conn.SetWriteDeadline(now.Add(time.Second))
	if err := conn.WriteJSON(command); err != nil {
		log.Print("Failed to send command")
		return fmt.Errorf("failed to send command: %w", err)
	}

	conn.SetReadDeadline(now.Add(10 * time.Second))
	if err := conn.ReadJSON(response); err != nil {
		log.Print("Failed to read socket response")
		return fmt.Errorf("failed to read response: %w", err)
	}

	if result, ok := response.(*xtb.Response); ok && !result.Status {
		log.Print("Received error response")
		return fmt.Errorf("error in response: %s", result.ErrorDescr)
	}

	return nil
}
