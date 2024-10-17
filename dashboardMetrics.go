package models

type DashboardMetrics struct {
	Type             string `json:"type" gorm:"column:type"`
	Last1Day         int    `json:"l1d" gorm:"column:l1d"`
	Last2Days        int    `json:"l2d" gorm:"column:l2d"`
	Last7Days        int    `json:"l7d" gorm:"column:l7d"`
	Last14Days       int    `json:"l14d" gorm:"column:l14d"`
	Last30Days       int    `json:"l30d" gorm:"column:l30d"`
	Last60Days       int    `json:"l60d" gorm:"column:l60d"`
	Difference1Day   int    `json:"diff1d" gorm:"column:diff1d"`
	Difference7Days  int    `json:"diff7d" gorm:"column:diff7d"`
	Difference30Days int    `json:"diff30d" gorm:"column:diff30d"`
}
