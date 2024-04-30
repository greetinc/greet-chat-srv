package dto

import "greet-chat-srv/entity"

type GetAllRequest struct {
	Page     int      `query:"page"`
	Limit    int      `query:"limit"`
	Searchs  []Search `query:"searchs" json:"searchs"`
	SortBy   string   `query:"sort_by"`
	SortDesc bool     `query:"sort_desc"`
	UserID   string   `json:"user_id"`
}

type PaginationResponse struct {
	TotalData    int             `json:"total_data"`
	TotalRows    int             `json:"total_rows"`
	Limit        int             `json:"limit"`
	PreviousPage int             `json:"previous_page"`
	NextPage     int             `json:"next_page"`
	NextPageData int             `json:"next_page_data"`
	Data         []entity.Friend `json:"data"`
	Searchs      []Search        `json:"searchs"`
}
