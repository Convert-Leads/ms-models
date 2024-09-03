package models

type Role struct {
	QModel
	OrganisationID uint   `json:"oid,omitempty"` // Foreign Key
	Name           string `json:"n"`
	Description    string `json:"d"`
}
