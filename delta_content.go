package models

import (
    "gorm.io/gorm"
)

type DeltaContent struct {
	gorm.Model
	ContentId  uint   `json:"content_id"`  // Foreign key to reference Content
	Data       string `json:"data"`        // String to store the Delta format JSON
	CreatedBy  uint   `json:"created_by"`  // ID of the user who created the record
	UpdatedBy  uint   `json:"updated_by"`  // ID of the user who last updated the record
}
