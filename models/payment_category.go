package models

import "time"

type PaymentCategories struct {
	ID         int64      `gorm:"primary_key;auto_increment" json:"id"`
	Name       string     `gorm:"name"`
	Amount     float64    `gorm:"amount"`
	CurrencyID int        `json:"currency_id"`
	Currency   Currencies `gorm:"foreignKey:CurrencyID"`
	SchoolID   int        `json:"school_id"`
	School     User       `gorm:"foreignKey:User"`
	CreatedAt  time.Time  `json:"created_at,omitempty"`
	UpdatedAt  time.Time  `json:"updated_at,omitempty"`
}

func (paymentCategory *PaymentCategories) TableName() string {
	return "paymentCategories"
}

type RecordPaymentCategories struct {
	Name       string  `form:"name"`
	Amount     float64 `form:"amount"`
	CurrencyID int     `form:"currency_id" binding:"required"`
	SchoolID   int     `form:"school_id" binding:"required"`
}

type UpdatePaymentCategories struct {
	ID         int64   `form:"id"`
	Amount     float64 `form:"amount"`
	CurrencyID int     `form:"currency_id" binding:"required"`
}

func (paymentCategory *PaymentCategories) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = paymentCategory.ID
	resp["name"] = paymentCategory.Name
	resp["currency_id"] = paymentCategory.CurrencyID
	resp["school_id"] = paymentCategory.SchoolID

	return resp
}
