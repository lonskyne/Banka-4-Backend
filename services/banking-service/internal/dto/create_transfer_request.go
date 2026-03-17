package dto

type CreateTransferRequest struct {
	SourceAccountNum string  `json:"source_account_num" binding:"required,len=18"`
	DestAccountNum   string  `json:"dest_account_num" binding:"required,len=18"`
	Amount           float64 `json:"amount" binding:"required,gt=0"`
	Description      string  `json:"description"`
}
