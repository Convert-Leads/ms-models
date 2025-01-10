package models

type FirstTimeExperience struct {
	QModel
	UserId      uint `json:"user_id"`
	FirstLogin  bool `json:"fl,omitempty" gorm:"default:true"`
	Reels       bool `json:"r,omitempty" gorm:"default:true"`
	Newsletters bool `json:"n,omitempty" gorm:"default:true"`
	Collections bool `json:"c,omitempty" gorm:"default:true"`
	Circles     bool `json:"g,omitempty" gorm:"default:true"`
}
