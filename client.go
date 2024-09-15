package xtb

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Period int
type TradeCommand uint8
type TradeType uint8

const (
	maxViolationCount = 5 // User should send requests in 200 ms intervals. This rule can be broken, but if it happens 6 times in a row the connection is dropped.
)

type Client struct {
	url               *url.URL
	con               *websocket.Conn
	logger            *zap.Logger
	lastRequestAt     time.Time
	violationCount    int
	maxViolationCount int
}

func NewClient(serverURL string, logger *zap.Logger) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, fmt.Errorf("invalid server URL: %w", err)
	}

	dialer := websocket.DefaultDialer
	dialer.HandshakeTimeout = 10 * time.Second

	con, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to WebSocket: %w", err)
	}

	if logger == nil {
		logger, _ = zap.NewProduction()
	}

	return &Client{url: u, con: con, logger: logger, violationCount: 0, maxViolationCount: maxViolationCount}, nil
}

func (c *Client) executeCommand(command *Command, response interface{}) error {
	now := time.Now()

	if now.Sub(c.lastRequestAt) < 200*time.Millisecond {
		c.violationCount++

		if c.violationCount > c.maxViolationCount {
			c.logger.Error("Exceeded maximum request frequency violations.")
			sleepDuration := 200*time.Millisecond - now.Sub(c.lastRequestAt)
			time.Sleep(sleepDuration)
		}
	} else {
		c.violationCount = 0
	}

	c.lastRequestAt = time.Now()

	// Convert command to JSON.
	commandJSON, err := json.Marshal(command)
	if err != nil {
		c.logger.Error("Failed to marshal command", zap.Error(err))
		return fmt.Errorf("failed to marshal command: %w", err)
	}

	// Check command size limit.
	if len(commandJSON) > 1024 {
		c.logger.Error("Command size exceeds 1KB limit")
		return fmt.Errorf("command size exceeds 1KB limit")
	}

	// Log the command.
	// c.logger.Info("Sending command", zap.Any("command", command))

	// Set a write deadline.
	c.con.SetWriteDeadline(now.Add(time.Second))
	if err := c.con.WriteJSON(command); err != nil {
		c.logger.Error("Failed to send command", zap.Error(err))
		return fmt.Errorf("failed to send command: %w", err)
	}

	// Read response with a read deadline.
	c.con.SetReadDeadline(now.Add(10 * time.Second))
	if err := c.con.ReadJSON(response); err != nil {
		c.logger.Error("Failed to read response", zap.Error(err))
		return fmt.Errorf("failed to read response: %w", err)
	}

	// Check for error in response.
	if result, ok := response.(*Response); ok && !result.Status {
		c.logger.Error("Received error response", zap.Any("response", response))
		return fmt.Errorf("error in response: %s", result.ErrorDescr)
	}

	// Log the response
	// c.logger.Info("Received response", zap.Any("response", response))
	return nil

}
