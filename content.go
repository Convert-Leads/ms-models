package models

import (
    "gorm.io/gorm"
    "time"
)

type Content struct {
    gorm.Model
    OrganisationId     uint       `json:"organisation_id"`
    ContentTypeId      uint       `json:"content_type_id"`
    ContentType        ContentType `gorm:"foreignKey:ContentTypeId" json:"content_type"`
    AvailableDatetime  *time.Time `json:"available_datetime"`
    ExpiryDatetime     *time.Time `json:"expiry_datetime"`
    Pinned             bool       `json:"pinned"`
    MediaID            *uint       
    Media              *Media     `gorm:"foreignKey:MediaID" json:"media"` 
    ThumbnailID        *uint       
    Thumbnail          *Media     `gorm:"foreignKey:ThumbnailID" json:"thumbnail"` 
    UserInteractions   []UserInteraction `gorm:"foreignKey:ContentId" json:"user_interactions"` 
    DeltaContent       *DeltaContent `gorm:"foreignKey:ContentId" json:"delta_content"` 
    CreatedBy          uint       `json:"created_by"`  
    UpdatedBy          uint       `json:"updated_by"`  
    Caption            string     `json:"caption"` 
    CTA                *CallToAction     `json:"cta"`
    Tags               []MetadataTag `gorm:"many2many:content_tags;" json:"tags"`
    AllowedUsers       []User `gorm:"many2many:content_allowed_users;" json:"allowed_users"`
    LikeCount          int        `json:"like_count" gorm:"-"`
    Liked              bool       `json:"liked" gorm:"-"`
    Bookmarked         bool       `json:"bookmarked" gorm:"-"`
    CommentCount       int        `json:"comment_count" gorm:"-"`
}
