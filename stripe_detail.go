package models



type StripeDetail struct {
	QModel
	OrganisationId uint                  `json:"oi"`
	CustomerID     string                `json:"ci"`  // Stripe Customer ID
	SubscriptionID string                `json:"si"`  // Stripe Subscription ID
	Active         string                `json:"a"`   // Indicates if the subscription is active
	Payments       []PaymentTransaction  `json:"p" gorm:"foreignKey:StripeDetailID"`
}
