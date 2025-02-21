package types

// type struct for pagination response generic type page, total_page, data
type PaginationResponse struct {
	Page      int         `json:"page"`
	TotalPage int         `json:"total_page"`
	Data      interface{} `json:"data"`
}
