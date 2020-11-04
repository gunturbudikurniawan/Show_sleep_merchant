package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Onlinesales1 struct {
	ID             uint32    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create_dtm"`
	SalesId        string    `gorm:"size:50;" json:"sales_id"`
	UserId         string    `gorm:"size:50;" json:"user_id"`
	OutletId       string    `gorm:"size:50;" json:"outlet_id"`
	CustomerId     string    `gorm:"size:50;" json:"customer_id"`
	Customer       string    `gorm:"size:50;" json:"customer"`
	Products       string    `gorm:"not null;" json:"products,"`
	Subtotal       uint64    `gorm:json:"subtotal,"`
	TotalDiskon    uint64    `gorm:json:"total_diskon,"`
	TotalTax       uint64    `gorm:json:"total_tax,"`
	TotalBill      uint64    `gorm:json:"total_bill,"`
	PaymentMethod  string    `gorm:"size:50;not null;" json:"payment_method"`
	PaymentAccount string    `gorm:"size:50;not null;" json:"payment_account"`
	PaymentDueDate string    `gorm:"size:50;not null;" json:"payment_due_date"`
	TotalPayment   uint64    `gorm:json:"total_payment,"`
	Expedition     string    `gorm:"size:50;not null;" json:"expedition"`
	Service        string    `gorm:"size:50;not null;" json:"service"`
	Weight         uint64    `gorm:json:"weight,"`
	DeliveryCost   uint64    `gorm:json:"delivery_cost,"`
	Notes          string    `gorm:"size:100;" json:"notes"`
	TotalBuyCost   uint64    `gorm:json:"total_buy_cost,"`
	PaymentDate    string    `gorm:"size:100;" json:"payment_date"`
	RewardId       string    `gorm:"size:50;not null;" json:"reward_id"`
	PointsRedeem   uint64    `gorm:json:"points_redeem,"`
	OrderStatus    string    `gorm:"size:50;not null;" json:"order_status"`
	ShipmentNumber string    `gorm:"size:50;not null;" json:"shipment_number"`
}

func (k *Onlinesales1) Prepare() {

	k.PaymentAccount = html.EscapeString(strings.TrimSpace(k.PaymentAccount))
	k.CreatedAt = time.Now()
}

func (k *Onlinesales1) Validate() map[string]string {

	var err error

	var errorMessages = make(map[string]string)

	if k.CustomerId == "" {
		err = errors.New("Required Customer")
		errorMessages["Required_Customer"] = err.Error()
	}
	if k.OutletId == "" {
		err = errors.New("Required Outlet")
		errorMessages["Required_Outlet"] = err.Error()
	}
	return errorMessages
}

func (k *Onlinesales1) SaveOnlineSales(db *gorm.DB) (*Onlinesales1, error) {
	var err error
	err = db.Debug().Model(&Onlinesales1{}).Create(&k).Error
	if err != nil {
		return &Onlinesales1{}, err
	}

	if k.UserId == "" {
		err = db.Debug().Model(&Onlinesales1{}).Where("user_id = ?", k.UserId).Error
		if err != nil {
			return &Onlinesales1{}, err
		}
	}
	return k, nil
}
