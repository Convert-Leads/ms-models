package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	OrganisationId      uint                 `json:"o,omitempty"`
	FirstName           string               `json:"fn,omitempty"`
	LastName            string               `json:"ln,omitempty"`
	ContactAddress      string               `json:"cad,omitempty"`
	ContactPhone        string               `json:"cp,omitempty"`
	ContactEmail        string               `json:"ce,omitempty"`
	ContactTwitter      string               `json:"ct,omitempty"`
	ContactFacebook     string               `json:"cf,omitempty"`
	ContactYoutube      string               `json:"cy,omitempty"`
	ContactInstagram    string               `json:"ci,omitempty"`
	ContactLinkedin     string               `json:"cl,omitempty"`
	Registered          bool                 `json:"r,omitempty" gorm:"default:false"`
	Roles               []Role               `json:"ro,omitempty" gorm:"many2many:user_roles;"`
	Bookmarks           []UserBookmark       `json:"b,omitempty" gorm:"one2many:user_bookmarks;"`
	Groups              []Group              `json:"g,omitempty" gorm:"many2many:user_groups;"`
	Posts               []Post               `json:"p,omitempty" gorm:"foreignKey:UserId"`
	ProfileImage        string               `json:"pi,omitempty"`
	contentInteractions []ContentInteraction `json:"int,omitempty" gorm:"foreignkey:UserId"`
	UserSubscription    []UserSubscription   `json:"subs,omitempty" gorm:"foreignkey:UserId"`
	FCMToken            string               `json:"-"`
}
