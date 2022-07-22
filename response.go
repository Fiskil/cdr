package cdr

// BaseResponse is the base values that all endpoints return.
type BaseResponse struct {
	Meta struct {
		TotalRecords int `json:"totalRecords"`
		Totalpages   int `json:"totalPages"`
	} `json:"meta"`
	Links struct {
		Self     string `json:"self"`
		First    string `json:"first"`
		Previous string `json:"prev"`
		Next     string `json:"next"`
		Last     string `json:"last"`
	} `json:"links"`
}
