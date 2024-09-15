package xtb

type getSymbolRequest struct {
	Symbol string `json:"symbol"` // Symbol.
}

type GetSymbolResponse struct {
	Response
	ReturnData *SymbolRecord `json:"returnData"`
}

func newGetSymbolCommand(symbol string) *Command {
	return &Command{
		Command: "getSymbol",
		Arguments: &getSymbolRequest{
			Symbol: symbol,
		},
	}
}
