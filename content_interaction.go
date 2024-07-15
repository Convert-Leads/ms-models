package models

type ContentInteraction struct {
	QModel
	OrganisationId         uint `json:"-"`
	UserId                 uint `json:"-"`
	User 								 User `json:"u"`
	ParentType             string `json:"-"`
	ParentID               uint `json:"-"`
	InteractionType        string `json:"it"`
	InteractionData				string `json:"ida"`
}
