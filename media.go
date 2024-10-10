package models

import "gorm.io/gorm"

type Media struct {
	gorm.Model
	OrganisationId uint   `json:"o"`
	Uri            string `json:"uri"`
	ParentType     string `json:"parentType"`
	ParentID       uint   `json:"parentId"`
	MimeType       string `json:"mimeType"`
	Transcoded     bool   `json:"transcoded"`
}
