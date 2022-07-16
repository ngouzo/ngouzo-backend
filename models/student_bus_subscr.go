package models

import "time"

type StudentBusSubScription struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	StudentID int       `json:"student_id"`
	Students  Students  `gorm:"foreignKey:StudentID"`
	SchoolID  int       `json:"school_id"`
	School    User      `gorm:"foreignKey:User"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (studentBusSubScription *StudentBusSubScription) TableName() string {
	return "StudentBusSubScription"
}

type RecordStudentBusSubScription struct {
	StudentID int `form:"student_id" binding:"required"`
	SchoolID  int `form:"school_id" binding:"required"`
}

type UpdateStudentBusSubScription struct {
	StudentID int `form:"student_id" binding:"required"`
}

func (studentBusSubScription *StudentBusSubScription) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = studentBusSubScription.ID
	resp["student_id"] = studentBusSubScription.StudentID
	resp["school_id"] = studentBusSubScription.SchoolID

	return resp
}
