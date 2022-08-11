package models

import "time"

type EmployeePaymentHistory struct {
	ID          int64      `gorm:"primary_key;auto_increment"  json:"id"`
	EmployeeID  int        `json:"employee_id"`
	Employee    User       `gorm:"foreignKey:User"`
	Amount      float64    `json:"amount"`
	Month       string     `json:"month"`
	Year        string     `json:"year"`
	Description string     `json:"description"`
	CurrencyID  int        `json:"currency_id"`
	Currencies  Currencies `gorm:"foreignKey:CurrencyID"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty"`
}

func (employeePaymentHistory *EmployeePaymentHistory) TableName() string {
	return "EmployeePaymentHistory"
}

type RecordEmployeePaymentHistory struct {
	EmployeeID  int     `form:"employee_id" binding:"required"`
	Description string  `form:"description" binding:"required"`
	Amount      float64 `form:"amount" binding:"required"`
	Month       string  `json:"month"`
	Year        string  `json:"year"`
	CurrencyID  int     `form:"currency_id"`
}

type UpdateEmployeePaymentHistory struct {
	ID          int64  `form:"id" binding:"required"`
	Description string `form:"description" binding:"required"`
	Month       string `json:"month"`
	Year        string `json:"year"`
	CurrencyID  int    `form:"currency_id"`
}

func (employeePaymentHistory *EmployeePaymentHistory) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = employeePaymentHistory.ID
	resp["description"] = employeePaymentHistory.Description
	resp["amount"] = employeePaymentHistory.Amount
	resp["employee_id"] = employeePaymentHistory.EmployeeID
	resp["month"] = employeePaymentHistory.Month
	resp["year"] = employeePaymentHistory.Year
	resp["currency_id"] = employeePaymentHistory.CurrencyID

	return resp
}
