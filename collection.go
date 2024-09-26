package models

// Collection represents a group of chapters, like a book or a course.
type Collection struct {
	Model
	OrganisationID uint                    `json:"oid,omitempty"` // Foreign Key
	Title          string                  `json:"t,omitempty"`
	Type           string                  `json:"tp,omitempty"`
	Description    string                  `json:"d,omitempty"`
	LikeCount      uint                    `json:"lc"`                                                          // New field
	ViewCount      uint                    `json:"vc"`                                                          // New field for view count
	Interactions   []CollectionInteraction `json:"int" gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE;"` // Related interactions
	Image          *Media                  `json:"img,omitempty" gorm:"polymorphic:Parent;polymorphicValue:collections"`
	Chapters       []Chapter               `json:"ch,omitempty" gorm:"foreignKey:CollectionID;references:ID"` // Explicit one-to-many relationship
	Files          []Content               `json:"f,omitempty" gorm:"many2many:collection_files;"`            // Changed from ExtraFiles to Files and json tag from "ef" to "f"
}

// Chapter represents individual chapters, sections, or units within a collection.
type Chapter struct {
	QModel
	CollectionID      uint             `json:"cid,omitempty"` // Foreign Key
	Title             string           `json:"t,omitempty"`
	Description       string           `json:"d,omitempty"`
	Image             *Media           `json:"img,omitempty" gorm:"polymorphic:Parent;polymorphicValue:chapters"`
	OrderInCollection int              `json:"oic,omitempty"`
	Contents          []ChapterContent `json:"c,omitempty" gorm:"foreignKey:ChapterID;references:ID"` // Explicit one-to-many relationship
}

// ChapterContent represents the actual content or items within each chapter.
type ChapterContent struct {
	QModel
	ChapterID          uint                 `json:"cid,omitempty"` // Foreign Key
	Title              string               `json:"t,omitempty"`
	ContentId          uint                 `json:"ctid,omitempty"`
	Content            Content              `json:"cnt,omitempty" gorm:"foreignKey:ContentId;references:ID"`
	OrderInChapter     int                  `json:"oic,omitempty"`
	ContentInteraction []ContentInteraction `json:"ci,omitempty" gorm:"polymorphic:Parent;polymorphicValue:chaptercontents"`
	HasAccess          bool                 `json:"has_access" gorm:"-"`
}
