package models

type CallToAction struct {
	gorm.Model
	Text 			 string `json:"text"`
	Type       string `json:"type"`
	ContentID uint   `json:"content_id"`
	Content   *Content `json:"content"`
	URL        string `json:"url"`
}