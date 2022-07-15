package models

import "time"

type TeacherClass struct {
	ID        int64        `gorm:"primary_key;auto_increment" json:"id"`
	ClassID   int          `json:"module_id"`
	Class     ClassModules `gorm:"foreignKey:ClassID"`
	TeacherID int          `json:"teacher_id"`
	Teacher   User         `gorm:"foreignKey:User"`
	CreatedAt time.Time    `json:"created_at,omitempty"`
	UpdatedAt time.Time    `json:"updated_at,omitempty"`
}

func (teacherClassModule *TeacherClass) TableName() string {
	return "TeacherClass"
}

type RecordTeacherClass struct {
	ClassID   int `form:"class_id" binding:"required"`
	TeacherID int `form:"teacher_id" binding:"required"`
}

type UpdateTeacherClass struct {
	ClassID int `form:"class_id" binding:"required"`
}

func (teacherClassModule *TeacherClass) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = teacherClassModule.ID
	resp["class_id"] = teacherClassModule.ClassID
	resp["teacher_id"] = teacherClassModule.TeacherID

	return resp
}
