package models

// Collection represents a group of chapters, like a book or a course.
type Collection struct {
	Model
	OrganisationID uint                    `json:"oid,omitempty"` // Foreign Key
	Title          string                  `json:"t,omitempty"`
	Type           string                  `json:"tp,omitempty"`
	Description    string                  `json:"d,omitempty" gorm:"type:text"`
	LikeCount      uint                    `json:"lc"`                                                          // New field
	ViewCount      uint                    `json:"vc"`                                                          // New field for view count
	Interactions   []CollectionInteraction `json:"int" gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE;"` // Related interactions
	Image          *Media                  `json:"img,omitempty" gorm:"polymorphic:Parent;polymorphicValue:collections"`
	ImageID        *uint                   `json:"-"`                                                         // Add this field to track the image ID
	Chapters       []Chapter               `json:"ch,omitempty" gorm:"foreignKey:CollectionID;references:ID"` // Explicit one-to-many relationship
	// Remove the Files field from here
}

// Chapter represents individual chapters, sections, or units within a collection.
type Chapter struct {
	QModel
	CollectionID      uint             `json:"cid,omitempty"` // Foreign Key
	Title             string           `json:"t,omitempty"`
	Description       string           `json:"d,omitempty" gorm:"type:text"`
	Image             *Media           `json:"img,omitempty" gorm:"polymorphic:Parent;polymorphicValue:chapters"`
	OrderInCollection int              `json:"oic,omitempty"`
	Files             []Content        `json:"f,omitempty" gorm:"many2many:chapter_files;"`           // Add this line
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
