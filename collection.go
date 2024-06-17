package models

import (
    "gorm.io/gorm"
)

// Collection represents a group of chapters, like a book or a course.
type Collection struct {
	gorm.Model
	OrganisationID uint      `json:"organisationId"` // Foreign Key
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageID		 uint      `json:"imageId"` // Foreign Key
	Image			 Media		 `json:"image" gorm:"foreignKey:ImageID;references:ID"`
	Chapters    []Chapter `json:"chapters" gorm:"foreignKey:CollectionID;references:ID"` // Explicit one-to-many relationship
}

// Chapter represents individual chapters, sections, or units within a collection.
type Chapter struct {
	gorm.Model
	CollectionID      uint      `json:"collectionId"` // Foreign Key
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	ImageID		 uint      `json:"imageId"` // Foreign Key
	Image			 Media		 `json:"image" gorm:"foreignKey:ImageID;references:ID"`
	OrderInCollection int       `json:"orderInCollection"`
	Contents          []ChapterContent `json:"contents" gorm:"foreignKey:ChapterID;references:ID"` // Explicit one-to-many relationship
}

// Content represents the actual content or items within each chapter.
type ChapterContent struct {
	gorm.Model
	ChapterID      uint   `json:"chapterId"` // Foreign Key
	Title          string `json:"title"`
	ContentId		uint   `json:"contentId"`
	Content			 Content `json:"content" gorm:"foreignKey:ContentId;references:ID"` 
	OrderInChapter int    `json:"orderInChapter"`
}

