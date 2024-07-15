package models


type Media struct {
    QModel
    OrganisationId uint   `json:"organisation_id"`
    Uri            string `json:"uri"`
    ParentType          string              `json:"-"`
    ParentID            uint                `json:"-"`
}