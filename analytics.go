package models

// NewsletterOverview holds the summary statistics for all newsletters.
type NewsletterOverview struct {
	AverageOpenRate  int     `json:"averageOpenRate"`
	TotalRevenue     float64 `json:"totalRevenue"`
	TotalConversions int     `json:"totalConversions"`
}

// NewsletterDetail holds the statistics for a single newsletter campaign.
type NewsletterDetail struct {
	ID            string  `json:"id"`
	Title         string  `json:"title"`
	OpenRate      int     `json:"openRate"`
	Revenue       float64 `json:"revenue"`
	CtaClicks     int     `json:"ctaClicks"`
	UnopenedCount int     `json:"unopenedCount"`
}

// ClipOverview holds the summary statistics for all clips.
type ClipOverview struct {
	TotalReelViews   int     `json:"totalReelViews"`
	TotalRevenue     float64 `json:"totalRevenue"`
	TotalConversions int     `json:"totalConversions"`
}

// ClipDetail holds the statistics and metadata for a single clip.
type ClipDetail struct {
	ID        string  `json:"id"`
	Caption   string  `json:"caption"`
	Revenue   float64 `json:"revenue"`
	CtaClicks int     `json:"ctaClicks"`
	Comments  int     `json:"comments"`
	Likes     int     `json:"likes"`
	Views     int     `json:"views"`
	Bookmarks int     `json:"bookmarks"`
}

// CollectionOverview holds the summary statistics for all collections.
type CollectionOverview struct {
	AverageWatchTime int64   `json:"averageWatchTimeSeconds"`
	TotalRevenue     float64 `json:"totalRevenue"`
	TotalConversions int     `json:"totalConversions"`
}

// CollectionDetail holds the statistics and metadata for a single collection.
type CollectionDetail struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	WatchTime int64   `json:"watchTime"`
	Revenue   float64 `json:"revenue"`
	Likes     int     `json:"likes"`
	Views     int     `json:"views"`
}

// CircleOverview holds the summary statistics for all circles.
type CircleOverview struct {
	TotalMembers     int     `json:"totalMembers"`
	TotalRevenue     float64 `json:"totalRevenue"`
	TotalConversions int     `json:"totalConversions"`
}

// CircleDetail holds the statistics and metadata for a single circle.
type CircleDetail struct {
	ID                   string  `json:"id"`
	Name                 string  `json:"name"`
	SubscriptionRequired bool    `json:"subscriptionRequired"`
	Members              int     `json:"members"`
	Revenue              float64 `json:"revenue"`
	TotalPosts           int     `json:"totalPosts"`
	TotalComments        int     `json:"totalComments"`
	TotalLikes           int     `json:"totalLikes"`
}
