package models

import "gorm.io/gorm"

type Visit struct {
	gorm.Model
	Referrer string `json:"referrer"`
	Url      string `json:"url"`
	Timestamp string `json:"timestamp"`
	UserAgent string `json:"userAgent"`
	ScreenResolution string `json:"screenResolution"`
	Timezone string `json:"timezone"`
	Language string `json:"language"`
	Platform string `json:"platform"`
	VisitorId string `json:"visitorId"`
	Pathname string `json:"pathname"`
	Host string `json:"host"`
	VisitorIP string `json:"visitorIP"`
}

type SignUpDetails struct {
	gorm.Model
	Referrer string `json:"referrer"`
	Url      string `json:"url"`
	Timestamp string `json:"timestamp"`
	UserAgent string `json:"userAgent"`
	AppName string `json:"appName"`
	FullName string `json:"fullName"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	ButtonType string `json:"buttonType"`
	VisitorId string `json:"visitorId"`
	Pathname string `json:"pathname"`
	Host string `json:"host"`
	VisitorIP string `json:"visitorIP"`
}
