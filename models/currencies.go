package models

import "time"

type Currencies struct {
	ID          int64     `gorm:"primary_key;auto_increment"  json:"id"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func (currency *Currencies) TableName() string {
	return "currency"
}

type RecordCurrencies struct {
	Code        string `form:"code" binding:"required"`
	Description string `form:"description" binding:"required"`
}

type UpdateCurrencies struct {
	ID          int64  `form:"id" binding:"required"`
	Code        string `form:"code"`
	Description string `form:"description"`
}

func (currency *Currencies) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = currency.ID
	resp["code"] = currency.Code
	resp["description"] = currency.Description

	return resp
}
