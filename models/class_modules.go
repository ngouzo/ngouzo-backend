package models

import "time"

type ClassModules struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Grades    string    `gorm:"grades"`
	Hours     float64   `gorm:"hours"`
	ClassID   int       `json:"class_id"`
	Class     Classes   `gorm:"foreignKey:ClassID"`
	ModuleID  int       `json:"module_id"`
	Module    Classes   `gorm:"foreignKey:ModuleID"`
	SchoolID  int       `json:"school_id"`
	School    User      `gorm:"foreignKey:User"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (classModule *ClassModules) TableName() string {
	return "ClassModules"
}

type RecordClassModules struct {
	Hours    string  `form:"hours"`
	Grades   float64 `form:"grades"`
	ClassID  int     `form:"class_id" binding:"required"`
	ModuleID int     `form:"module_id" binding:"required"`
	SchoolID int     `form:"school_id" binding:"required"`
}

type UpdateClassModules struct {
	Hours    string  `form:"hours"`
	Grades   float64 `form:"grades"`
	ClassID  int     `form:"class_id" binding:"required"`
	ModuleID int     `form:"module_id" binding:"required"`
}

func (classModule *ClassModules) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = classModule.ID
	resp["grades"] = classModule.Grades
	resp["hours"] = classModule.Hours
	resp["class_id"] = classModule.ClassID
	resp["module_id"] = classModule.ModuleID
	resp["school_id"] = classModule.SchoolID

	return resp
}
