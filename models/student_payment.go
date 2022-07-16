package models

import "time"

type StudentPayment struct {
	ID                int64             `gorm:"primary_key;auto_increment"  json:"id"`
	StudentID         int               `json:"student_id"`
	Students          Students          `gorm:"foreignKey:StudentID"`
	PaymentCategoryID int               `json:"payment_category_id"`
	PaymentCategories PaymentCategories `gorm:"foreignKey:PaymentCategoryID"`
	FeeTypeID         int               `json:"fee_type_id"`
	FeeType           FeeType           `gorm:"foreignKey:FeeTypeID"`
	CurrencyID        int               `json:"currency_id"`
	Currencies        Currencies        `gorm:"foreignKey:CurrencyID"`
	Description       string            `json:"description"`
	AmountPaid        float64           `json:"amount_paid"`
	CreatedAt         time.Time         `json:"created_at,omitempty"`
	UpdatedAt         time.Time         `json:"updated_at,omitempty"`
}

func (studentPayment *StudentPayment) TableName() string {
	return "StudentPayment"
}

type RecordStudentPayment struct {
	StudentID         int     `form:"student_id" binding:"required"`
	Description       string  `form:"description" binding:"required"`
	AmountPaid        float64 `form:"amount_paid" binding:"required"`
	FeeTypeID         int     `form:"fee_type_id"`
	PaymentCategoryID int     `form:"payment_category_id"`
	CurrencyID        int     `form:"currency_id"`
}

type UpdateStudentPayment struct {
	ID                int64  `form:"id" binding:"required"`
	Description       string `form:"description"`
	FeeTypeID         int    `form:"fee_type_id"`
	PaymentCategoryID int    `form:"payment_category_id"`
	CurrencyID        int    `form:"currency_id"`
}

func (studentPayment *StudentPayment) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = studentPayment.ID
	resp["description"] = studentPayment.Description
	resp["amount_paid"] = studentPayment.AmountPaid
	resp["student_id"] = studentPayment.StudentID
	resp["fee_type_id"] = studentPayment.FeeTypeID
	resp["payment_category_id"] = studentPayment.PaymentCategoryID
	resp["currency_id"] = studentPayment.CurrencyID

	return resp
}
