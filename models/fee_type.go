package models

import "time"

type FeeType struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (class *FeeType) TableName() string {
	return "FeeType"
}

type RecordFeeType struct {
	Name string `form:"name" binding:"required"`
}

type UpdateFeeType struct {
	ID   int64  `form:"id" binding:"required"`
	Name string `form:"name" binding:"required"`
}

func (feeType *FeeType) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = feeType.ID
	resp["name"] = feeType.Name

	return resp
}
