package xclient

import (
	"github.com/kejrak/xtb-client-go/xtb"
)

// SubscribeGetBalance subscribes to getBalance command described here: http://developers.xstore.pro/documentation/#streamgetBalance
func (c *client) SubscribeGetBalance() error {

	err := c.executeStreamCommand(xtb.NewSubscribeGetBalanceCommand(c.streamSessionID))
	if err != nil {
		return err

	}

	return nil
}

// SubscribeGetCandles subscribes to getCandles command described here: http://developers.xstore.pro/documentation/#streamgetCandles
func (c *client) SubscribeGetCandles(symbol string) error {

	err := c.executeStreamCommand(xtb.NewSubscribeGetCandlesCommand(c.streamSessionID, symbol))
	if err != nil {
		return err

	}
	return nil
}

// SubscribeGetKeepAlive subscribes to getKeepAlive command described here: http://developers.xstore.pro/documentation/#streamgetKeepAlive
func (c *client) SubscribeGetKeepAlive() error {

	err := c.executeStreamCommand(xtb.NewSubscribeGetKeepAliveCommand(c.streamSessionID))
	if err != nil {
		return err

	}
	return nil
}

// SubscribeGetNews subscribes to getNews command described here: http://developers.xstore.pro/documentation/#streamgetNews
func (c *client) SubscribeGetNews() error {

	err := c.executeStreamCommand(xtb.NewSubscribeGetNewsCommand(c.streamSessionID))
	if err != nil {
		return err

	}
	return nil
}

// SubscribeGetProfits subscribes to getProfits command described here: http://developers.xstore.pro/documentation/#streamgetProfits
func (c *client) SubscribeGetProfits() error {

	err := c.executeStreamCommand(xtb.NewSubscribeGetProfitsCommand(c.streamSessionID))
	if err != nil {
		return err

	}
	return nil
}

// SubscribeGetTickPrices subscribes to getTickPrices command described here: http://developers.xstore.pro/documentation/#streamgetTickPrices
func (c *client) SubscribeGetTickPrices(symbol string, minArrivalTime *int, maxLevel *uint8) error {

	err := c.executeStreamCommand(xtb.NewSubscribeGetTickPricesCommand(c.streamSessionID, symbol, minArrivalTime, maxLevel))
	if err != nil {
		return err

	}
	return nil
}

// SubscribeGetTrades subscribes to getTrades command described here: http://developers.xstore.pro/documentation/#streamgetTrades
func (c *client) SubscribeGetTrades() error {

	err := c.executeStreamCommand(xtb.NewSubscribeGetTradesCommand(c.streamSessionID))
	if err != nil {
		return err

	}
	return nil
}

// SubscribeGetTradeStatus subscribes to getTradeStatus command described here: http://developers.xstore.pro/documentation/#streamgetTradeStatus
func (c *client) SubscribeGetTradeStatus() error {

	err := c.executeStreamCommand(xtb.NewSubscribeGetTradeStatusCommand(c.streamSessionID))
	if err != nil {
		return err

	}
	return nil
}

// SubscribePing subscribes to ping command described here: http://developers.xstore.pro/documentation/#streamping
func (c *client) SubscribePing() error {

	err := c.executeStreamCommand(xtb.NewSubscribePingCommand(c.streamSessionID))
	if err != nil {
		return err

	}
	return nil
}
