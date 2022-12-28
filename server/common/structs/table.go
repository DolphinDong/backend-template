package structs

type TableQuery struct {
	Page     int    `json:"page,string" form:"page"`
	PageSize int    `json:"page_size,string" form:"page_size"`
	Search   string `json:"search" form:"search"`
}

type TableResponse struct {
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}
