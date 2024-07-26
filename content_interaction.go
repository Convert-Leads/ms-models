package models

type ContentInteraction struct {
	QModel
	OrganisationId  uint   `json:"-"`
	UserId          uint   `json:"user_id"`
	User            User   `json:"u"`
	ParentType      string `json:"parent_type"`
	ParentID        uint   `json:"parent_id"`
	InteractionType string `json:"it"`
	InteractionData string `json:"ida"`
}
