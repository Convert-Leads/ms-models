package models


type Media struct {
    QModel
    OrganisationId uint   `json:"-"`
    Uri            string `json:"uri"`
    ParentType          string              `json:"-"`
    ParentID            uint                `json:"-"`
}