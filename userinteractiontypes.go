type UserInteractionTypes struct {
    gorm.Model
    Name        string `json:"name"`
    Description string `json:"description"`
}
