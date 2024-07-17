package models



type Role struct {
    QModel
    Name        string `json:"n"`
    Description string `json:"d"`
}
