package xtb

type GetAllSymbolsResponse struct {
	Response
	ReturnData []*SymbolRecord `json:"returnData"`
}

func NewGetAllSymbolsCommand() *Command {
	return &Command{
		Command: "getAllSymbols",
	}
}
