package models

import (
	"time"

	"github.com/google/uuid"
)

// Currency base model
type Currency struct {
	CurrencyID uuid.UUID `json:"category_id" db:"category_id" validate:"omitempty,uuid"`
	Name       string    `json:"name" db:"name" validate:"required,gte=2"`
	Code       string    `json:"code" db:"code" validate:"required,gte=2"`
	Symbol     string    `json:"symbol" db:"symbol" validate:"required,gte=2"`
	CreatedAt  time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// CurrenciesList All Currencies response
type CurrenciesList struct {
	TotalCount int         `json:"total_count"`
	TotalPages int         `json:"total_pages"`
	Page       int         `json:"page"`
	Size       int         `json:"size"`
	HasMore    bool        `json:"has_more"`
	Categories []*Currency `json:"currencies"`
}
