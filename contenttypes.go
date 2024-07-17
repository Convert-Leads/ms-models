package models

type ContentType struct {
    QModel
    Name           string `json:"n"`
    Description    string `json:"-"`
}
