package models

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	Caption					string            `json:"caption"`
	AvailableDatetime  *time.Time `json:"available_datetime"`
  ExpiryDatetime     *time.Time `json:"expiry_datetime"`
	Pinned						 bool       `json:"pinned"` // max 3
	Media					*Media            `gorm:"polymorphic:Parent;polymorphicValue:posts" json:"content"`
	UserId  uint              `json:"user_id"`
	GroupId          uint              `json:"group_id"`
	User             User              `gorm:"foreignKey:UserId" json:"user"`
	Group            Group             `gorm:"foreignKey:GroupId" json:"group"`
	UserInteractions []UserInteraction `gorm:"polymorphic:Parent;polymorphicValue:posts" json:"user_interactions"`
	Poll							*Poll              `gorm:"foreignKey:PostId" json:"poll"`
	Mentions					[]User          `gorm:"polymorphic:Parent;polymorphicValue:posts" json:"mentions"`
}

type Poll struct {
	gorm.Model
	Question string `json:"question"`
	PostId uint `json:"post_id"`
	Post   Post `gorm:"foreignKey:PostId" json:"post"`
	Options []PollOption `json:"options"`
}

type PollOption struct {
	gorm.Model
	PollId uint `json:"poll_id"`
	Poll   Poll `gorm:"foreignKey:PollId" json:"poll"`
	Option string `json:"option"`
	Votes []PollVote `json:"votes"`
}

type PollVote struct {
	gorm.Model
	PollOptionId uint `json:"poll_option_id"`
	PollOption   PollOption `gorm:"foreignKey:PollOptionId" json:"poll_option"`
	UserId       uint `json:"user_id"`
	User         User `gorm:"foreignKey:UserId" json:"user"`
}