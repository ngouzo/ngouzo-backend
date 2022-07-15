package models

import "time"

type Classes struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (class *Classes) TableName() string {
	return "Classes"
}

type RecordClasses struct {
	Name string `form:"name" binding:"required"`
}

type UpdateClasses struct {
	ID   int64  `form:"id" binding:"required"`
	Name string `form:"name" binding:"required"`
}

func (class *Classes) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = class.ID
	resp["name"] = class.Name

	return resp
}
