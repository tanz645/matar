package responses

type FailedResponse struct {
	Status  int         `json:"status"`
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type SuccessResponse struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ListingResponse struct {
	TotalCount uint64      `json:"total_count"`
	List       interface{} `json:"list"`
}
