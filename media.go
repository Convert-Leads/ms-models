package models


type Media struct {
    Qmodel
    OrganisationId uint   `json:"-"`
    Uri            string `json:"uri"`
    ParentType          string              `json:"-"`
    ParentID            uint                `json:"-"`
}