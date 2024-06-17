package models

type CallToAction struct {
	gorm.Model
	Text 			 string `json:"text"`
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
	URL        string `json:"url"`
}