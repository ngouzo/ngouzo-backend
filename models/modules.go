package models

import "time"

type Modules struct {
	ID        int64     `gorm:"primary_key;auto_increment"  json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (module *Modules) TableName() string {
	return "module"
}

type RecordModules struct {
	Name string `form:"name" binding:"required"`
}

type UpdateModules struct {
	ID   int64  `form:"id" binding:"required"`
	Name string `form:"name"`
}

func (module *Modules) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = module.ID
	resp["name"] = module.Name

	return resp
}
