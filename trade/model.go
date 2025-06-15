package trade

type Deal struct {
	InputToken  string `json:"inputToken" binding:"required"`
	InputAmount string `json:"inputAmount" binding:"required"`
	OutputToken string `json:"OutputToken" binding:"required"`
}
