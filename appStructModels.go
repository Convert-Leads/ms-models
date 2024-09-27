package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type LoginInfo struct {
	Token string `json:"token" binding:"required"`
}

type FakeLoginInfo struct {
	Email string `json:"email"`
}

type CustomClaims struct {
	Uid    uint     `json:"userId"`
	Oid    uint     `json:"orgId"`
	Rights int      `json:"rights"`
	Rol    []string `json:"roles"`
	jwt.StandardClaims
}

type FirebaseTokenClaims struct {
	IssuedAtTime   int64  `json:"iat"`
	Issuer         string `json:"iss"`
	Audience       string `json:"aud"`
	ExpiresAt      int64  `json:"exp"`
	Subject        string `json:"sub"`
	Email          string `json:"email"`
	EmailVerified  bool   `json:"email_verified"`
	Name           string `json:"name"`
	Picture        string `json:"picture"`
	UserID         string `json:"user_id"`
	SignInProvider string `json:"firebase.sign_in_provider"`
	// Add other fields as necessary
}

type CommentRequest struct {
	Body string `json:"comment" binding:"required"`
}

type CommentResponse struct {
	ID          uint              `json:"i"`
	UserName    string            `json:"un"`
	UserProfile string            `json:"up"`
	Comment     string            `json:"c"`
	Timestamp   time.Time         `json:"t"`
	Likes       int               `json:"l"`
	Replies     []CommentResponse `json:"r"`
	Edited      bool              `json:"e"`
	Liked       bool              `json:"lk"`
}

type ResponseInteractionUserView struct {
	LikeCount    int `json:"like_count"`
	CommentCount int `json:"comment_count"`
	Comments     []struct {
		User    string `json:"user"`
		Comment string `json:"comment"`
	} `json:"comments"`
}

type Comment struct {
	InteractionDetails string `gorm:"column:interaction_details"`
	UserName           string `gorm:"column:user_name"`
}
type CommentsInteractionUserView struct {
	Comment []Comment `json:"cuv"`
}

type VaultContent struct {
	Recent   []Content    `json:"recent"`
	LongForm []Content    `json:"long_form"`
	Courses  []Collection `json:"courses"`
	Podcasts []Collection `json:"podcasts"`
	Other    []Content    `json:"other"`
}

type RequestDataSubscription struct {
	Name                 string `json:"name"`
	Description          string `json:"description"`
	AppLogo              uint   `json:"app_logo"`
	AppName              string `json:"app_name"`
	WelcomeScreen        uint   `json:"welcome_screen"`
	ContactAddress       string `json:"address"`
	ContactEmail         string `json:"contact_email"`
	ContactPhone         string `json:"contact_phone"`
	Email                string `json:"email"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	PaymentMethod        string `json:"paymentMethod"`
	SubscriptionName     string `json:"subscriptionName"`
	SubscriptionPrice    int64  `json:"subscriptionPrice"` // assuming price in smallest currency unit (e.g., cents)
	SubscriptionInterval string `json:"subscriptionInterval"`
	SubscriptionCurrency string `json:"subscriptionCurrency"`
}

type AssociationData struct {
	AllowedContentTypes []ContentType `json:"allowed_content_types"`
	AllowedMetadataTags []MetadataTag `json:"allowed_metadata_tags"`
}

type TeamsResponse struct {
	Users []User `json:"users"`
	Team  bool   `json:"team"`
}
