package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Caption           string            `json:"caption" gorm:"type:text"`
	AvailableDatetime *time.Time        `json:"available_datetime"`
	ExpiryDatetime    *time.Time        `json:"expiry_datetime"`
	Pinned            bool              `json:"pinned"` // max 3
	Media             *Media            `gorm:"polymorphic:Parent;polymorphicValue:posts" json:"content"`
	UserId            uint              `json:"user_id"`
	GroupId           uint              `json:"group_id"`
	User              User              `gorm:"foreignKey:UserId" json:"user"`
	Group             Group             `gorm:"foreignKey:GroupId" json:"group"`
	UserInteractions  []UserInteraction `gorm:"polymorphic:Parent;polymorphicValue:posts" json:"user_interactions"`
	Poll              *Poll             `gorm:"foreignKey:PostId" json:"poll"`
	Mentions          []User            `gorm:"many2many:post_user_mentions;" json:"mentions"`
	LikeCount         int               `json:"lc" gorm:"-"`
	Liked             bool              `json:"l" gorm:"-"`
	Bookmarked        bool              `json:"b" gorm:"-"`
	CommentCount      int               `json:"cc" gorm:"-"`
}

type Poll struct {
	gorm.Model
	Question string       `json:"question" gorm:"type:text"`
	PostId   uint         `json:"post_id"`
	Post     Post         `gorm:"foreignKey:PostId" json:"post"`
	Options  []PollOption `json:"options"`
}

type PollOption struct {
	gorm.Model
	PollId uint       `json:"poll_id"`
	Poll   Poll       `gorm:"foreignKey:PollId" json:"poll"`
	Option string     `json:"option"`
	Votes  []PollVote `json:"votes"`
}

type PollVote struct {
	gorm.Model
	PollOptionId uint       `json:"poll_option_id"`
	PollOption   PollOption `gorm:"foreignKey:PollOptionId" json:"poll_option"`
	UserId       uint       `json:"user_id"`
	User         User       `gorm:"foreignKey:UserId" json:"user"`
}
