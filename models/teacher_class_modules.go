package models

import "time"

type TeacherClassModules struct {
	ID        int64        `gorm:"primary_key;auto_increment" json:"id"`
	ModuleID  int          `json:"module_id"`
	Module    ClassModules `gorm:"foreignKey:ModuleID"`
	TeacherID int          `json:"teacher_id"`
	Teacher   User         `gorm:"foreignKey:User"`
	CreatedAt time.Time    `json:"created_at,omitempty"`
	UpdatedAt time.Time    `json:"updated_at,omitempty"`
}

func (teacherClassModule *TeacherClassModules) TableName() string {
	return "TeacherClassModules"
}

type RecordTeacherClassModules struct {
	ModuleID  int `form:"module_id" binding:"required"`
	TeacherID int `form:"teacher_id" binding:"required"`
}

type UpdateTeacherClassModules struct {
	ModuleID int `form:"module_id" binding:"required"`
}

func (teacherClassModule *TeacherClassModules) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = teacherClassModule.ID
	resp["module_id"] = teacherClassModule.ModuleID
	resp["teacher_id"] = teacherClassModule.TeacherID

	return resp
}
