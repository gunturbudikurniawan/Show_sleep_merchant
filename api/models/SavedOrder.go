package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type SavedOrder struct {
	ID              uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserId          string    `gorm:"size:50;not null;" json:"user_id"`
	OutletId        string    `gorm:"size:50;not null;" json:"outlet_id"`
	Saved_orders_id string    `gorm:"size:50;" json:"saved_orders_id"`
	Name            string    `gorm:"size:100;" json:"name"`
	Phone           string    `gorm:"size:20;null;" json:"phone"`
	TableId         string    `gorm:"size:20;null;" json:"table_id"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create_dtm"`
}

func (s *SavedOrder) Prepare() {
	s.Name = html.EscapeString(strings.TrimSpace(s.Name))
	s.CreatedAt = time.Now()
}

func (s *SavedOrder) Validate() map[string]string {

	var err error

	var errorMessages = make(map[string]string)

	if s.Name == "" {
		err = errors.New("Required Name")
		errorMessages["Required_Name"] = err.Error()

	}
	if s.OutletId == "" {
		err = errors.New("Required Content")
		errorMessages["Required_content"] = err.Error()

	}
	return errorMessages
}

func (s *SavedOrder) SaveOrder(db *gorm.DB) (*SavedOrder, error) {
	var err error
	err = db.Debug().Model(&SavedOrder{}).Create(&s).Error
	if err != nil {
		return &SavedOrder{}, err
	}
	if s.UserId == "" {
		err = db.Debug().Model(&Subscribers1{}).Where("user_id = ?", s.UserId).Error
		if err != nil {
			return &SavedOrder{}, err
		}
	}
	return s, nil
}
