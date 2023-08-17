package models

import (
	"time"

	"github.com/google/uuid"
)

// Category base model
type Category struct {
	CategoryID uuid.UUID  `json:"category_id" db:"category_id" validate:"omitempty,uuid"`
	UserID     uuid.UUID  `json:"user_id" db:"user_id" validate:"omitempty,uuid"`
	ParentID   uuid.UUID  `json:"parent_id" db:"parent_id" validate:"omitempty,uuid"`
	Name       string     `json:"name" db:"name" validate:"required,gte=2"`
	CreatedAt  time.Time  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// CategoriesList All Categories response
type CategoriesList struct {
	TotalCount int         `json:"total_count"`
	TotalPages int         `json:"total_pages"`
	Page       int         `json:"page"`
	Size       int         `json:"size"`
	HasMore    bool        `json:"has_more"`
	Categories []*Category `json:"categories"`
}
