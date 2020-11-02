package models

import (
	"errors"
	"time"
)

type SalesOrder struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Qty         uint64 `gorm: json:"qty"`
	Price       uint64 `gorm: json:"price"`
	Total       uint64 `gorm: json:"total"`
	NamaProduct string `gorm:"text;not null;" json:"product_name"`
	Category    string `json:"category"`

	Author    User      `json:"author"`
	AuthorID  uint32    `gorm:"not null" json:"author_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (s *SalesOrder) BeforeSave() error {

	if s.Category != "WhatsApp" {
		return errors.New("Invalid Category")
	}
	return nil
}
