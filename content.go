package models

import (
    "gorm.io/gorm"
    "time"
)

type Content struct {
    gorm.Model
    OrganisationId     uint             `json:"-"`
    ContentTypeId      uint             `json:"t"`
    ContentType        ContentType      `gorm:"foreignKey:ContentTypeId" json:"ct"`
    AvailableDatetime  *time.Time       `json:"ad,omitempty"`
    ExpiryDatetime     *time.Time       `json:"ed,omitempty"`
    Pinned             bool             `json:"p"`
    MediaID            *uint            `json:"-"`
    Media              *Media           `gorm:"polymorphic:Parent;polymorphicValue:contents;foreignKey:MediaID" json:"m,omitempty"`
    ThumbnailID        *uint            `json:"-"`
    Thumbnail          *Media           `gorm:"polymorphic:Parent;polymorphicValue:contents;foreignKey:ThumbnailID" json:"th,omitempty"`
    UserInteractions   []UserInteraction `gorm:"polymorphic:Parent;polymorphicValue:contents" json:"ui,omitempty"`
    DeltaContent       *DeltaContent    `gorm:"foreignKey:ContentId" json:"dc,omitempty"`
    Newsletter         *Newsletter      `gorm:"foreignKey:ContentID" json:"nl,omitempty"`
    CreatedBy          uint             `json:"cb"`
    UpdatedBy          uint             `json:"ub"`
    Caption            string           `json:"c,omitempty"`
    CTA                *CallToAction    `json:"cta,omitempty" gorm:"polymorphic:Parent;polymorphicValue:contents"`
    Tags               []MetadataTag    `gorm:"many2many:content_tags;" json:"tags,omitempty"`
    AllowedUsers       []User           `gorm:"many2many:content_allowed_users;" json:"au,omitempty"`
    LikeCount          int              `json:"lc" gorm:"-"`
    Liked              bool             `json:"l" gorm:"-"`
    Bookmarked         bool             `json:"b" gorm:"-"`
    CommentCount       int              `json:"cc" gorm:"-"`
}
