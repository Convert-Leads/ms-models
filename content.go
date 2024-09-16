package models

import (
	"time"
)

type Content struct {
	Model
	OrganisationId     uint                 `json:"-"`
	ContentTypeId      uint                 `json:"t"`
	ContentType        ContentType          `gorm:"foreignKey:ContentTypeId" json:"ct,omitempty"`
	AvailableDatetime  *time.Time           `json:"ad,omitempty"`
	ExpiryDatetime     *time.Time           `json:"ed,omitempty"`
	Pinned             bool                 `json:"p"`
	Media              *Media               `gorm:"polymorphic:Parent;polymorphicValue:contentmedia" json:"m,omitempty"`
	Thumbnail          *Media               `gorm:"polymorphic:Parent;polymorphicValue:contentthumbnail" json:"th,omitempty"`
	UserInteractions   []UserInteraction    `gorm:"polymorphic:Parent;polymorphicValue:contents" json:"ui,omitempty"`
	Newsletter         *Newsletter          `gorm:"foreignKey:ContentID;constraint:OnDelete:CASCADE" json:"nl,omitempty"`
	CreatedBy          uint                 `json:"cb,omitempty"`
	UpdatedBy          uint                 `json:"ub,omitempty"`
	Caption            string               `json:"c,omitempty"`
	CTA                *CallToAction        `json:"cta,omitempty" gorm:"polymorphic:Parent;polymorphicValue:contents"`
	Tags               []MetadataTag        `gorm:"many2many:content_tags;" json:"tags,omitempty"`
	AllowedUsers       []User               `gorm:"many2many:content_allowed_users;" json:"au,omitempty"`
	ContentInteraction []ContentInteraction `json:"-" gorm:"polymorphic:Parent;polymorphicValue:contents"`
	Runtime            int                  `json:"r,omitempty"`
    LikeCount          int                  `json:"lc" gorm:"-"`
	Liked              bool                 `json:"l" gorm:"-"`
	Bookmarked         bool                 `json:"b" gorm:"-"`
	CommentCount       int                  `json:"cc" gorm:"-"`
	WatchCount         int                  `json:"wc" gorm:"-"`
	HasAccess          bool                 `json:"ha" gorm:"-"`
	Runtime            *int                 `json:"r,omitempty" gorm:"default:null"` // Change to pointer type
}
