package models

// PaginationQuery represents pagination and filtering parameters
type PaginationQuery struct {
	Page     int         `form:"page,default=1" json:"page"`
	PageSize int         `form:"page_size,default=10" json:"page_size"`
	Status   *TaskStatus `form:"status" json:"status,omitempty"`
}

// PaginatedResponse represents a paginated response
type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Page       int        `json:"page"`
	PageSize   int        `json:"page_size"`
	TotalItems int64      `json:"total_items"`
	TotalPages int        `json:"total_pages"`
}

// NewPaginatedResponse creates a new paginated response
func NewPaginatedResponse(data interface{}, query PaginationQuery, totalItems int64) PaginatedResponse {
	totalPages := int(totalItems) / query.PageSize
	if int(totalItems)%query.PageSize != 0 {
		totalPages++
	}

	return PaginatedResponse{
		Data:       data,
		Page:       query.Page,
		PageSize:   query.PageSize,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}
} 