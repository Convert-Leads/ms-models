type Media struct {
    gorm.Model
    OrganisationId int    `json:"organisation_id"`
    Uri            string `json:"uri"`
}
