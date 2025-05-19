package base

type Response struct {
	Data       interface{} `json:"data"`
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Pagination Pagination  `json:"pagination,omitempty"`
}

type Pagination struct {
	ItemPerPage uint `json:"item_per_page"`
	CurrentPage uint `json:"current_page"`
	TotalCount  uint `json:"total_count"`
}
