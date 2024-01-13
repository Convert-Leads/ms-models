package models

import (
    "gorm.io/gorm"
)

type Organisation struct {
    gorm.Model
    Name           string `json:"name"`
    Description    string `json:"description"`
    App_logo       int    `json:"app_logo"`
    App_name       string `json:"app_name"`
    Welcome_screen int    `json:"welcome_screen"`
    Address        string `json:"address"`
    Contact_email  string `json:"contact_email"`
    Contact_phone  string `json:"contact_phone"`
}
