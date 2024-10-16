package models

type DashboardMetrics struct {
	Type       string `gorm:"column:type"`
	Last1Day   int    `gorm:"column:ld"`
	Last7Days  int    `gorm:"column:lsd"`
	Last30Days int    `gorm:"column:ltd"`
}
