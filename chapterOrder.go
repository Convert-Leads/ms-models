package models

// ChapterOrdersRequest represents the request body containing a list of chapter orders
type ChapterOrdersRequest struct {
	Orders []ChapterOrderUpdate `json:"orders"`
}

// ChapterOrderUpdate represents the data needed to update a chapter's order
type ChapterOrderUpdate struct {
	ID    uint `json:"id"`
	Order int  `json:"orderInCollection"`
}
