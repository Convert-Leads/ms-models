package models

type CallToAction struct {
	QModel
	Text            string     `json:"t,omitempty"`
	Type            string     `json:"tp,omitempty"`
	ParentID        uint       `json:"pid,omitempty"`
	ParentType      string     `json:"pt,omitempty"`
	TargetContentID *uint      `json:"tcid,omitempty"`
	Content         *Content   `gorm:"foreignKey:TargetContentID" json:"cnt,omitempty"`
	CollectionID    *uint      `json:"cid,omitempty"`
	Collection      *Collection `gorm:"foreignKey:CollectionID" json:"clc,omitempty"`
	URL             string     `json:"url,omitempty"`
	ContentInteraction          []ContentInteraction    `json:"ci,omitempty" gorm:"polymorphic:Parent;polymorphicValue:callstoaction"`
}