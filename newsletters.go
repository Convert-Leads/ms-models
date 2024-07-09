package models

// NewsletterElement struct
type NewsletterElement struct {
	QModel
	ParentID       uint                         `json:"-"`
	ParentType     string                       `json:"-" gorm:"type:varchar(50);polymorphic:Parent"`
	Type           string                       `json:"t,omitempty"`
	Data           string                       `json:"d,omitempty"`
	Order          int                          `json:"o,omitempty"`
	Media          *Media                       `gorm:"polymorphic:Parent;polymorphicValue:NewsletterElements" json:"m,omitempty"`
	Decorations    NewsletterElementDecoration  `json:"dec,omitempty" gorm:"foreignKey:NewsletterElementID"`
	Children       []NewsletterElement          `json:"ch,omitempty" gorm:"foreignKey:ParentID;polymorphicValue:NewsletterElements"`
	CTA            *CallToAction                `json:"cta,omitempty" gorm:"polymorphic:Parent;polymorphicValue:NewsletterElements"`
}

// NewsletterElementDecoration struct
type NewsletterElementDecoration struct {
	QModel
	NewsletterElementID uint   `json:"-"`
	BgColor             string `json:"bg,omitempty"`
	FontSize            int    `json:"fs,omitempty"`
	FontWeight          string `json:"fw,omitempty"`
	FontFamily          string `json:"ff,omitempty"`
	FontColour          string `json:"fc,omitempty"`
	Alignment           string `json:"a,omitempty"`
	Padding             string `json:"p,omitempty"`
	Margin              string `json:"m,omitempty"`
	Height              string `json:"h,omitempty"`
	Width               string `json:"w,omitempty"`
}

// NewsletterTemplate struct
type NewsletterTemplate struct {
	Model
	Title       string              `json:"t,omitempty"`
	Description string              `json:"d,omitempty"`
	Thumbnail   *Media              `gorm:"polymorphic:Parent;polymorphicValue:NewsletterTemplates" json:"th,omitempty"`
	Elements    []NewsletterElement `gorm:"polymorphic:Parent;polymorphicValue:NewsletterTemplates" json:"e,omitempty"`
}

// Newsletter struct
type Newsletter struct {
	Model
	OrganisationID uint                `json:"-"`
	TemplateID     uint                `json:"t_id,omitempty"`
	ContentID      uint                `json:"-"`
	Title          string              `json:"t,omitempty"`
	Description    string              `json:"d,omitempty"`
	ThumbnailID    *uint               `json:"-"`
	Thumbnail      *Media              `gorm:"polymorphic:Parent;polymorphicValue:NewsletterTemplates;foreignKey:ThumbnailID" json:"th,omitempty"`
	Elements       []NewsletterElement `gorm:"polymorphic:Parent;polymorphicValue:Newsletters" json:"e,omitempty"`
}