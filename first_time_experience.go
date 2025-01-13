package models

type FirstTimeExperience struct {
	QModel
	UserId      uint `json:"user_id"`
	FirstLogin  bool `json:"fl" gorm:"default:true"`
	Reels       bool `json:"r" gorm:"default:true"`
	Newsletters bool `json:"n" gorm:"default:true"`
	Collections bool `json:"c" gorm:"default:true"`
	Circles     bool `json:"g" gorm:"default:true"`
}
