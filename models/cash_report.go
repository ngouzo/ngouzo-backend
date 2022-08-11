package models

import "time"

type CashReport struct {
	ID          int64      `gorm:"primary_key;auto_increment" json:"id"`
	Description string     `gorm:"description"`
	Debit       float64    `gorm:"debit"`
	Credit      float64    `gorm:"credit"`
	CurrencyID  int        `json:"currency_id"`
	Currency    Currencies `gorm:"foreignKey:CurrencyID"`
	SchoolID    int        `json:"school_id"`
	School      User       `gorm:"foreignKey:User"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty"`
}

func (cashReport *CashReport) TableDescription() string {
	return "CashReport"
}

type RecordCashReport struct {
	Description string  `form:"description"`
	Debit       float64 `form:"debit"`
	Credit      float64 `form:"credit"`
	CurrencyID  int     `form:"currency_id" binding:"required"`
	SchoolID    int     `form:"school_id" binding:"required"`
}

type UpdateCashReport struct {
	ID          int64  `form:"id"`
	Description string `form:"description"`
	CurrencyID  int    `form:"currency_id" binding:"required"`
}

func (cashReport *CashReport) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = cashReport.ID
	resp["debit"] = cashReport.Debit
	resp["credit"] = cashReport.Credit
	resp["description"] = cashReport.Description
	resp["currency_id"] = cashReport.CurrencyID
	resp["school_id"] = cashReport.SchoolID

	return resp
}
