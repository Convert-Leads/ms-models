package models

import "gorm.io/gorm"

type Conversions struct {
	gorm.Model
	UserID        uint  `json:"uid"` // Foreign key for User
	Converted     bool  `json:"cnv" gorm:"default:false"`
	PriceOptionID *uint `json:"pid"`            // Foreign key for PriceOption
	ContentID     *uint `json:"cid,omitempty"`  // Foreign key for Content
	CollectionID  *uint `json:"clid,omitempty"` // Foreign key for Collection
	GroupID       *uint `json:"gid,omitempty"`  // Foreign key for Group
}
