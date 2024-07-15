package models

type ContentInteraction struct {
	QModel
	UserId                 uint `json:"userId"`
	ParentType             string `json:"parentType"`
	ParentID               uint `json:"parentId"`
	InteractionType        string `json:"interactionType"`
	InteractionData				string `json:"interactionData"`
}
