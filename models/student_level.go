package models

import "time"

type StudentLevel struct {
	ID                 int64     `gorm:"primary_key;auto_increment"  json:"id"`
	StudentID          int       `json:"student_id"`
	Students           Students  `gorm:"foreignKey:StudentID"`
	CurrentLevel       string    `json:"current_level"`
	StuddingRangeLevel string    `json:"studding_range_year"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
}

func (studentLevel *StudentLevel) TableName() string {
	return "StudentLevel"
}

type RecordStudentLevel struct {
	StudentID          int    `form:"student_id" binding:"required"`
	CurrentLevel       string `form:"current_level" binding:"required"`
	StuddingRangeLevel string `form:"studding_range_year" binding:"required"`
}

type UpdateStudentLevel struct {
	ID                 int64  `form:"id" binding:"required"`
	CurrentLevel       string `form:"current_level"`
	StuddingRangeLevel string `form:"studding_range_year"`
}

func (studentLevel *StudentLevel) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = studentLevel.ID
	resp["current_level"] = studentLevel.CurrentLevel
	resp["studding_range_year"] = studentLevel.StuddingRangeLevel
	resp["student_id"] = studentLevel.StudentID

	return resp
}
