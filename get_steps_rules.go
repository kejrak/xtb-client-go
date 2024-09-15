package xtb

type StepRuleRecord struct {
	ID    int          `json:"id"`    // Step rule ID.
	Name  string       `json:"name"`  // Step rule name.
	Steps []StepRecord `json:"steps"` // Array of StepRecord.
}

type StepRecord struct {
	FromValue float32 `json:"fromValue"` // Lower border of the volume range.
	Step      float64 `json:"step"`      // lotStep value in the given volume range.
}

type GetStepRulesResponse struct {
	Response
	ReturnData []StepRuleRecord `json:"returnData"`
}

func newGetStepRulesCommand() *Command {
	return &Command{
		Command: "getStepRules",
	}
}
