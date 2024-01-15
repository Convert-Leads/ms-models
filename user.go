package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    OrganisationId   uint   `json:"organisation_id"`
    First_name       string `json:"first_name"`
    Last_name        string `json:"last_name"`
    Contact_address  string `json:"contact_address"`
    Contact_phone    string `json:"contact_phone"`
    Contact_email    string `json:"contact_email"`
    Contact_twitter  string `json:"contact_twitter"`
    Contact_facebook string `json:"contact_facebook"`
    Contact_youtube  string `json:"contact_youtube"`
    Contact_instagram string `json:"contact_instagram"`
    Contact_linkedin string `json:"contact_linkedin"`
    Roles   []Role  `json:"roles"`
}
