package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Sales1 struct {
	ID             uint32    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create_dtm"`
	SalesId        string    `gorm:"size:50;" json:"sales_id"`
	UserId         string    `gorm:"size:50;" json:"user_id"`
	OutletId       string    `gorm:"size:50;" json:"outlet_id"`
	SalesType      string    `gorm:"size:50;" json:"sales_type"`
	CustomerId     string    `gorm:"size:50;" json:"customer_id"`
	Products       string    `gorm:"not null;" json:"products,"`
	Subtotal       uint64    `gorm:json:"subtotal,"`
	TotalDiskon    uint64    `gorm:json:"total_diskon,"`
	TotalBill      uint64    `gorm:json:"total_bill,"`
	PaymentMethod  string    `gorm:"size:50;not null;" json:"payment_method"`
	PaymentDueDate string    `gorm:"size:50;not null;" json:"payment_due_date"`
	TotalPayment   uint64    `gorm:json:"total_payment,"`
	Exchange       uint64    `gorm:json:"exchange,"`
	Notes          string    `gorm:"size:100;" json:"notes"`
	TotalBuyCost   uint64    `gorm:json:"total_buy_cost,"`
	PaymentDate    string    `gorm:"size:100;" json:"payment_date"`
	TotalTax       uint64    `gorm:json:"total_tax,"`
	RewardId       string    `gorm:"size:50;not null;" json:"reward_id"`
	PointsRedeem   uint64    `gorm:json:"points_redeem,"`
}

func (w *Sales1) Prepare() {
	w.SalesType = html.EscapeString(strings.TrimSpace(w.SalesType))
	w.CreatedAt = time.Now()
}

func (w *Sales1) Validate() map[string]string {

	var err error

	var errorMessages = make(map[string]string)

	if w.CustomerId == "" {
		err = errors.New("Required Customer")
		errorMessages["Required_Customer"] = err.Error()
	}
	if w.OutletId == "" {
		err = errors.New("Required Outlet")
		errorMessages["Required_Outlet"] = err.Error()
	}
	return errorMessages
}

func (w *Sales1) SaveSales(db *gorm.DB) (*Sales1, error) {
	var err error
	err = db.Debug().Model(&Sales1{}).Create(&w).Error
	if err != nil {
		return &Sales1{}, err
	}
	if w.UserId == "" {
		err = db.Debug().Model(&Sales1{}).Where("user_id = ?", w.UserId).Error
		if err != nil {
			return &Sales1{}, err
		}
	}
	return w, nil
}
