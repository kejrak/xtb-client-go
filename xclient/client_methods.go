package xclient

import (
	"time"

	"github.com/kejrak/xtb-client-go/xtb"
)

// GetVersion performs getVersion command described here: http://developers.xstore.pro/documentation/#getVersion
func (c *Client) GetVersion() (*xtb.GetVersionResponse, error) {
	res := &xtb.GetVersionResponse{}
	err := c.executeCommand(xtb.NewGetVersionCommand(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Ping performs ping command described here: http://developers.xstore.pro/documentation/#ping
func (c *Client) Ping() (*xtb.PingResponse, error) {
	res := &xtb.PingResponse{}
	err := c.executeCommand(xtb.NewPingCommand(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AUTH METHODS //

// Login performs login command described here: http://developers.xstore.pro/documentation/#login
func (c *Client) Login(user, password string) (*xtb.LoginResponse, error) {
	res := &xtb.LoginResponse{}
	err := c.executeCommand(xtb.NewLoginCommand(user, password), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Login performs login command described here: http://developers.xstore.pro/documentation/#login
func (c *Client) Logout() (*xtb.Response, error) {
	res := &xtb.Response{}
	err := c.executeCommand(xtb.NewLogoutCommand(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SYMBOLS METHODS //

// GetAllSymbols performs getAllSymbols command described here: http://developers.xstore.pro/documentation/#getAllSymbols
func (c *Client) GetAllSymbols() (*xtb.GetAllSymbolsResponse, error) {
	res := &xtb.GetAllSymbolsResponse{}
	err := c.executeCommand(xtb.NewGetAllSymbolsCommand(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetAllSymbols performs getAllSymbols command described here: http://developers.xstore.pro/documentation/#getSymbol
func (c *Client) GetSymbol(symbol string) (*xtb.GetSymbolResponse, error) {
	res := &xtb.GetSymbolResponse{}
	err := c.executeCommand(xtb.NewGetSymbolCommand(symbol), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CALENDAR METHODS //

// GetCalendar performs getCalendar command described here: http://developers.xstore.pro/documentation/#getCalendar
func (c *Client) GetCalendar() (*xtb.GetCalendarResponse, error) {
	res := &xtb.GetCalendarResponse{}
	err := c.executeCommand(xtb.NewGetCalendarCommand(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CHART METHODS //

// GetChartLastRequest performs getChartLastRequest command described here: http://developers.xstore.pro/documentation/#getChartLastRequest
func (c *Client) GetChartLastRequest(period xtb.Period, start time.Time, symbol string) (*xtb.GetChartLastRequestResponse, error) {
	res := &xtb.GetChartLastRequestResponse{}
	err := c.executeCommand(xtb.NewGetChartLastRequestCommand(period, start, symbol), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetChartRangeRequest performs getChartRangeRequest command described here: http://developers.xstore.pro/documentation/#getChartRangeRequest
func (c *Client) GetChartRangeRequest(end, start time.Time, period xtb.Period, symbol string, ticks *int) (*xtb.GetChartRangeRequestResponse, error) {
	res := &xtb.GetChartRangeRequestResponse{}
	err := c.executeCommand(xtb.NewGetChartRangeRequestCommand(end, start, period, symbol, ticks), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// COMMISSIONS METHODS //

// GetCommissionDef performs getCommissionDef command described here: http://developers.xstore.pro/documentation/#getCommissionDef
func (c *Client) GetCommissionDef(symbol string, volume float64) (*xtb.GetCommissionDefResponse, error) {
	res := &xtb.GetCommissionDefResponse{}
	err := c.executeCommand(xtb.NewGetCommissionDefCommand(symbol, volume), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// USER METHODS //

// GetCurrentUserData performs getCurrentUserData command described here: http://developers.xstore.pro/documentation/#getCurrentUserData
func (c *Client) GetCurrentUserData() (*xtb.GetCurrentUserDataResponse, error) {
	res := &xtb.GetCurrentUserDataResponse{}
	err := c.executeCommand(xtb.NewGetCurrentUserDataCommand(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// IBS METHODS //

// GetIbsHistory performs getIbsHistory command described here: http://developers.xstore.pro/documentation/#getIbsHistory
func (c *Client) GetIbsHistory(end, start time.Time) (*xtb.GetIbsHistoryResponse, error) {
	res := &xtb.GetIbsHistoryResponse{}
	err := c.executeCommand(xtb.NewGetIbsHistoryCommand(end, start), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MARGIN METHODS //

// GetMarginLevel performs getMarginLevel command described here: http://developers.xstore.pro/documentation/#getMarginLevel
func (c *Client) GetMarginLevel() (*xtb.GetMarginLevelResponse, error) {
	res := &xtb.GetMarginLevelResponse{}
	err := c.executeCommand(xtb.NewGetMarginLevelCommand(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetMarginTrade performs getMarginTrade command described here: http://developers.xstore.pro/documentation/#getMarginTrade
func (c *Client) GetMarginTrade(symbol string, volume float64) (*xtb.GetMarginTradeResponse, error) {
	res := &xtb.GetMarginTradeResponse{}
	err := c.executeCommand(xtb.NewGetMarginTradeCommand(symbol, volume), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// NewS METHODS //

// GetNews performs getNews command described here: http://developers.xstore.pro/documentation/#getNews
func (c *Client) GetNews(end, start time.Time) (*xtb.GetNewsResponse, error) {
	res := &xtb.GetNewsResponse{}
	err := c.executeCommand(xtb.NewGetNewsCommand(end, start), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PROFIT METHODS //

// GetProfitCalculation performs getProfitCalculation command described here: http://developers.xstore.pro/documentation/#getProfitCalculation
func (c *Client) GetProfitCalculation(closePrice, openPrice, volume float32, cmd xtb.TradeCommand, symbol string) (*xtb.GetProfitCalculationResponse, error) {
	res := &xtb.GetProfitCalculationResponse{}
	err := c.executeCommand(xtb.NewGetProfitCalculationCommand(closePrice, openPrice, volume, cmd, symbol), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SERVER METHODS //

// GetServerTime performs getServerTime command described here: http://developers.xstore.pro/documentation/#getServerTime
func (c *Client) GetServerTime() (*xtb.GetServerTimeResponse, error) {
	res := &xtb.GetServerTimeResponse{}
	err := c.executeCommand(xtb.NewGetServerTimeCommand(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// STEPS METHODS //

// GetStepRules performs getStepRules command described here: http://developers.xstore.pro/documentation/#getStepRules
func (c *Client) GetStepRules() (*xtb.GetStepRulesResponse, error) {
	res := &xtb.GetStepRulesResponse{}
	err := c.executeCommand(xtb.NewGetStepRulesCommand(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TICK METHODS //

// GetTickPrices performs getTickPrices command described here: http://developers.xstore.pro/documentation/#getTickPrices
func (c *Client) GetTickPrices(level int, symbols []string, timestamp time.Time) (*xtb.GetTickPricesResponse, error) {
	res := &xtb.GetTickPricesResponse{}
	err := c.executeCommand(xtb.NewGetTickPricesCommand(level, symbols, timestamp), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TRADE METHODS //

// GetTradeRecords performs getTradeRecords command described here: http://developers.xstore.pro/documentation/#getTradeRecords
func (c *Client) GetTradeRecords(orders []int) (*xtb.GetTradeRecordsResponse, error) {
	res := &xtb.GetTradeRecordsResponse{}
	err := c.executeCommand(xtb.NewGetTradeRecordsCommand(orders), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetTrades performs getTrades command described here: http://developers.xstore.pro/documentation/#getTradeRecords
func (c *Client) GetTrades(openedOnly bool) (*xtb.GetTradesResponse, error) {
	res := &xtb.GetTradesResponse{}
	err := c.executeCommand(xtb.NewGetTradesCommand(openedOnly), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetTradesHistory performs getTradesHistory command described here: http://developers.xstore.pro/documentation/#getTradesHistory
func (c *Client) GetTradesHistory(end, start time.Time) (*xtb.GetTradesHistoryResponse, error) {
	res := &xtb.GetTradesHistoryResponse{}
	err := c.executeCommand(xtb.NewGetTradesHistoryCommand(end, start), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetTradingHours performs getTradingHours command described here: http://developers.xstore.pro/documentation/#getTradingHours
func (c *Client) GetTradingHours(symbols []string) (*xtb.GetTradingHoursResponse, error) {
	res := &xtb.GetTradingHoursResponse{}
	err := c.executeCommand(xtb.NewGetTradingHoursCommand(symbols), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TradeTransaction performs tradeTransaction command described here: http://developers.xstore.pro/documentation/#tradeTransaction
func (c *Client) TradeTransaction(cmd xtb.TradeCommand, customComment string, expiration int64, offset int, order uint, price float64, sl float64, symbol string, tp float64, tradeType xtb.TradeType, volume float64) (*xtb.TradeTransactionResponse, error) {
	res := &xtb.TradeTransactionResponse{}
	err := c.executeCommand(xtb.NewTradeTransactionCommand(cmd, customComment, expiration, offset, order, price, sl, symbol, tp, tradeType, volume), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TradeTransactionStatus performs tradeTransactionStatus command described here: http://developers.xstore.pro/documentation/#tradeTransactionStatus
func (c *Client) TradeTransactionStatus(order uint) (*xtb.TradeTransactionStatusResponse, error) {
	res := &xtb.TradeTransactionStatusResponse{}
	err := c.executeCommand(xtb.NewTradeTransactionStatusCommand(order), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
