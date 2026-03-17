package dto

type ListTransfersQuery struct {
	AccountNum string `form:"account_num"`
	Status     string `form:"status"`
	StartDate  string `form:"start_date"` // RFC3339 format
	EndDate    string `form:"end_date"`   // RFC3339 format
	Page       int    `form:"page"`
	PageSize   int    `form:"page_size"`
}
