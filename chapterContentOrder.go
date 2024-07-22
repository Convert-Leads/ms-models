package models

type ChapterContentOrdersRequest struct {
	Orders []ChapterContentOrder `json:"orders"`
}

type ChapterContentOrder struct {
	ID    uint `json:"id"`
	Order int  `json:"orderInChapter"`
}
