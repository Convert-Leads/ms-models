package models

type ContentType struct {
    QModel
    Name           string `json:"name"`
    Description    string `json:"description"`
}
