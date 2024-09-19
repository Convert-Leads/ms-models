package models

type Media struct {
	QModel
	OrganisationId uint   `json:"o"`
	Uri            string `json:"uri"`
	ParentType     string `json:"-"`
	ParentID       uint   `json:"-"`
	MimeType       string `json:"mimeType"`
}
