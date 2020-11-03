package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/gunturbudikurniawan/Show_sleep_merchant/api/security"
	"github.com/jinzhu/gorm"
)

type Subscribers1 struct {
	ID             uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserId         string    `gorm:"size:50;" json:"user_id"`
	OwnerName      string    `gorm:"size:100;" json:"owner_name"`
	FcmToken       string    `gorm:"size:200;" json:"fcm_token"`
	IdcardName     string    `gorm:"size:50;" json:"idcard_name"`
	IdcardNumber   string    `gorm:"size:50;" json:"idcard_number"`
	BankHolderName string    `gorm:"size:256;" json:"bank_holder_name"`
	BankName       string    `gorm:"size:256;" json:"bank_name"`
	BankAccount    string    `gorm:"size:256;" json:"bank_account"`
	ReferralCode   string    `gorm:"size:256;" json:"referral_code"`
	Email          string    `gorm:"size:100;" json:"email"`
	SecretPassword string    `gorm:"size:100;" json:"secret_password,omitempty"`
	IdcardImage    string    `gorm: json:"idcard_image"`
	Created_Date   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create_dtm"`
}

func (m *Subscribers1) BeforeSave() error {
	hashedPassword, err := security.Hash(m.SecretPassword)
	if err != nil {
		return err
	}
	m.SecretPassword = string(hashedPassword)
	return nil
}

func (m *Subscribers1) Prepare() {

	m.Email = html.EscapeString(strings.TrimSpace(m.Email))
	m.FcmToken = html.EscapeString(strings.TrimSpace(m.FcmToken))
	m.IdcardName = html.EscapeString(strings.TrimSpace(m.IdcardName))
	m.IdcardNumber = html.EscapeString(strings.TrimSpace(m.IdcardNumber))
	m.OwnerName = html.EscapeString(strings.TrimSpace(m.OwnerName))
	m.BankAccount = html.EscapeString(strings.TrimSpace(m.BankAccount))
	m.Created_Date = time.Now()
}
func (m *Subscribers1) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)
	var err error

	switch strings.ToLower(action) {
	case "update":
		if m.Email == "" {
			err = errors.New("Required Email")
			errorMessages["Required_email"] = err.Error()
		}
		if m.Email != "" {
			if err = checkmail.ValidateFormat(m.Email); err != nil {
				err = errors.New("Invalid Email")
				errorMessages["Invalid_email"] = err.Error()
			}
		}

	case "login":
		if m.SecretPassword == "" {
			err = errors.New("Required Password")
			errorMessages["Required_password"] = err.Error()
		}
		if m.Email == "" {
			err = errors.New("Required Email")
			errorMessages["Required_email"] = err.Error()
		}
		if m.Email != "" {
			if err = checkmail.ValidateFormat(m.Email); err != nil {
				err = errors.New("Invalid Email")
				errorMessages["Invalid_email"] = err.Error()
			}
		}
	case "forgotpassword":
		if m.Email == "" {
			err = errors.New("Required Email")
			errorMessages["Required_email"] = err.Error()
		}
		if m.Email != "" {
			if err = checkmail.ValidateFormat(m.Email); err != nil {
				err = errors.New("Invalid Email")
				errorMessages["Invalid_email"] = err.Error()
			}
		}
	default:
		if m.OwnerName == "" {
			err = errors.New("Required Owner Name")
			errorMessages["Required Owner Name"] = err.Error()
		}
		if m.SecretPassword == "" {
			err = errors.New("Required Password")
			errorMessages["Required_password"] = err.Error()
		}
		if m.SecretPassword != "" && len(m.SecretPassword) < 6 {
			err = errors.New("Password should be atleast 6 characters")
			errorMessages["Invalid_password"] = err.Error()
		}
		if m.Email == "" {
			err = errors.New("Required Email")
			errorMessages["Required_email"] = err.Error()

		}
		if m.Email != "" {
			if err = checkmail.ValidateFormat(m.Email); err != nil {
				err = errors.New("Invalid Email")
				errorMessages["Invalid_email"] = err.Error()
			}
		}
	}
	return errorMessages
}
func (m *Subscribers1) SaveUser(db *gorm.DB) (*Subscribers1, error) {

	var err error
	err = db.Debug().Create(&m).Error
	if err != nil {
		return &Subscribers1{}, err
	}
	return m, nil
}
