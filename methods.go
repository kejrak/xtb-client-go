package xtb

import (
	"time"
)

// Close is closing connection with server.
func (c *Client) Close() error {
	return c.con.Close()
}

// GetVersion performs getVersion command described here: http://developers.xstore.pro/documentation/#getVersion
func (c *Client) GetVersion() (*GetVersionResponse, error) {
	res := &GetVersionResponse{}
	err := c.executeCommand(NewGetVersionCommand(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Ping performs ping command described here: http://developers.xstore.pro/documentation/#ping
func (c *Client) Ping() (*PingResponse, error) {
	res := &PingResponse{}
	err := c.executeCommand(NewPingCommand(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AUTH METHODS //

// Login performs login command described here: http://developers.xstore.pro/documentation/#login
func (c *Client) Login(user, password string) (*LoginResponse, error) {
	res := &LoginResponse{}
	err := c.executeCommand(NewLoginCommand(user, password), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Login performs login command described here: http://developers.xstore.pro/documentation/#login
func (c *Client) Logout() (*Response, error) {
	res := &Response{}
	err := c.executeCommand(NewLogoutCommand(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SYMBOLS METHODS //

// GetAllSymbols performs getAllSymbols command described here: http://developers.xstore.pro/documentation/#getAllSymbols
func (c *Client) GetAllSymbols() (*GetAllSymbolsResponse, error) {
	res := &GetAllSymbolsResponse{}
	err := c.executeCommand(NewGetAllSymbolsCommand(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetAllSymbols performs getAllSymbols command described here: http://developers.xstore.pro/documentation/#getSymbol
func (c *Client) GetSymbol(symbol string) (*GetSymbolResponse, error) {
	res := &GetSymbolResponse{}
	err := c.executeCommand(NewGetSymbolCommand(symbol), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CALENDAR METHODS //

// GetCalendar performs getCalendar command described here: http://developers.xstore.pro/documentation/#getCalendar
func (c *Client) GetCalendar() (*GetCalendarResponse, error) {
	res := &GetCalendarResponse{}
	err := c.executeCommand(NewGetCalendarCommand(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CHART METHODS //

// GetChartLastRequest performs getChartLastRequest command described here: http://developers.xstore.pro/documentation/#getChartLastRequest
func (c *Client) GetChartLastRequest(period Period, start time.Time, symbol string) (*GetChartLastRequestResponse, error) {
	res := &GetChartLastRequestResponse{}
	err := c.executeCommand(NewGetChartLastRequestCommand(period, start, symbol), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetChartRangeRequest performs getChartRangeRequest command described here: http://developers.xstore.pro/documentation/#getChartRangeRequest
func (c *Client) GetChartRangeRequest(end, start time.Time, period Period, symbol string, ticks *int) (*GetChartRangeRequestResponse, error) {
	res := &GetChartRangeRequestResponse{}
	err := c.executeCommand(NewGetChartRangeRequestCommand(end, start, period, symbol, ticks), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// COMMISSIONS METHODS //

// GetCommissionDef performs getCommissionDef command described here: http://developers.xstore.pro/documentation/#getCommissionDef
func (c *Client) GetCommissionDef(symbol string, volume float64) (*GetCommissionDefResponse, error) {
	res := &GetCommissionDefResponse{}
	err := c.executeCommand(NewGetCommissionDefCommand(symbol, volume), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// USER METHODS //

// GetCurrentUserData performs getCurrentUserData command described here: http://developers.xstore.pro/documentation/#getCurrentUserData
func (c *Client) GetCurrentUserData() (*GetCurrentUserDataResponse, error) {
	res := &GetCurrentUserDataResponse{}
	err := c.executeCommand(NewGetCurrentUserDataCommand(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// IBS METHODS //

// GetIbsHistory performs getIbsHistory command described here: http://developers.xstore.pro/documentation/#getIbsHistory
func (c *Client) GetIbsHistory(end, start time.Time) (*GetIbsHistoryResponse, error) {
	res := &GetIbsHistoryResponse{}
	err := c.executeCommand(NewGetIbsHistoryCommand(end, start), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MARGIN METHODS //

// GetMarginLevel performs getMarginLevel command described here: http://developers.xstore.pro/documentation/#getMarginLevel
func (c *Client) GetMarginLevel() (*GetMarginLevelResponse, error) {
	res := &GetMarginLevelResponse{}
	err := c.executeCommand(NewGetMarginLevelCommand(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetMarginTrade performs getMarginTrade command described here: http://developers.xstore.pro/documentation/#getMarginTrade
func (c *Client) GetMarginTrade(symbol string, volume float64) (*GetMarginTradeResponse, error) {
	res := &GetMarginTradeResponse{}
	err := c.executeCommand(NewGetMarginTradeCommand(symbol, volume), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// NewS METHODS //

// GetNews performs getNews command described here: http://developers.xstore.pro/documentation/#getNews
func (c *Client) GetNews(end, start time.Time) (*GetNewsResponse, error) {
	res := &GetNewsResponse{}
	err := c.executeCommand(NewGetNewsCommand(end, start), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PROFIT METHODS //

// GetProfitCalculation performs getProfitCalculation command described here: http://developers.xstore.pro/documentation/#getProfitCalculation
func (c *Client) GetProfitCalculation(closePrice, openPrice, volume float32, cmd TradeCommand, symbol string) (*GetProfitCalculationResponse, error) {
	res := &GetProfitCalculationResponse{}
	err := c.executeCommand(NewGetProfitCalculationCommand(closePrice, openPrice, volume, cmd, symbol), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SERVER METHODS //

// GetServerTime performs getServerTime command described here: http://developers.xstore.pro/documentation/#getServerTime
func (c *Client) GetServerTime() (*GetServerTimeResponse, error) {
	res := &GetServerTimeResponse{}
	err := c.executeCommand(NewGetServerTimeCommand(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// STEPS METHODS //

// GetStepRules performs getStepRules command described here: http://developers.xstore.pro/documentation/#getStepRules
func (c *Client) GetStepRules() (*GetStepRulesResponse, error) {
	res := &GetStepRulesResponse{}
	err := c.executeCommand(NewGetStepRulesCommand(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TICK METHODS //

// GetTickPrices performs getTickPrices command described here: http://developers.xstore.pro/documentation/#getTickPrices
func (c *Client) GetTickPrices(level int, symbols []string, timestamp time.Time) (*GetTickPricesResponse, error) {
	res := &GetTickPricesResponse{}
	err := c.executeCommand(NewGetTickPricesCommand(level, symbols, timestamp), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TRADE METHODS //

// GetTradeRecords performs getTradeRecords command described here: http://developers.xstore.pro/documentation/#getTradeRecords
func (c *Client) GetTradeRecords(orders []int) (*GetTradeRecordsResponse, error) {
	res := &GetTradeRecordsResponse{}
	err := c.executeCommand(NewGetTradeRecordsCommand(orders), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetTrades performs getTrades command described here: http://developers.xstore.pro/documentation/#getTradeRecords
func (c *Client) GetTrades(openedOnly bool) (*GetTradesResponse, error) {
	res := &GetTradesResponse{}
	err := c.executeCommand(NewGetTradesCommand(openedOnly), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetTradesHistory performs getTradesHistory command described here: http://developers.xstore.pro/documentation/#getTradesHistory
func (c *Client) GetTradesHistory(end, start time.Time) (*GetTradesHistoryResponse, error) {
	res := &GetTradesHistoryResponse{}
	err := c.executeCommand(NewGetTradesHistoryCommand(end, start), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetTradingHours performs getTradingHours command described here: http://developers.xstore.pro/documentation/#getTradingHours
func (c *Client) GetTradingHours(symbols []string) (*GetTradingHoursResponse, error) {
	res := &GetTradingHoursResponse{}
	err := c.executeCommand(NewGetTradingHoursCommand(symbols), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TradeTransaction performs tradeTransaction command described here: http://developers.xstore.pro/documentation/#tradeTransaction
func (c *Client) TradeTransaction(cmd TradeCommand, customComment string, expiration int64, offset int, order uint, price float64, sl float64, symbol string, tp float64, tradeType TradeType, volume float64) (*TradeTransactionResponse, error) {
	res := &TradeTransactionResponse{}
	err := c.executeCommand(NewTradeTransactionCommand(cmd, customComment, expiration, offset, order, price, sl, symbol, tp, tradeType, volume), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TradeTransactionStatus performs tradeTransactionStatus command described here: http://developers.xstore.pro/documentation/#tradeTransactionStatus
func (c *Client) TradeTransactionStatus(order uint) (*TradeTransactionStatusResponse, error) {
	res := &TradeTransactionStatusResponse{}
	err := c.executeCommand(NewTradeTransactionStatusCommand(order), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
