package models

import "time"

type StudentRelative struct {
	ID             int64     `gorm:"primary_key;auto_increment" json:"id"`
	StudentID      int       `json:"student_id"`
	Students       Students  `gorm:"foreignKey:StudentID"`
	FatherFullName string    `gorm:"father_full_name"`
	MotherFullName string    `gorm:"mother_full_name"`
	Email          string    `gorm:"email"`
	PhoneNumber    string    `gorm:"phone_number"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

func (studentRelative *StudentRelative) TableName() string {
	return "StudentRelative"
}

type RecordStudentRelative struct {
	StudentID      int    `form:"student_id" binding:"required"`
	FatherFullName string `form:"father_full_name"`
	MotherFullName string `form:"mother_full_name"`
	Email          string `form:"email" binding:"required"`
	PhoneNumber    string `form:"phone_number"`
}

type UpdateStudentRelative struct {
	ID             int    `form:"id" binding:"required"`
	FatherFullName string `form:"father_full_name"`
	MotherFullName string `form:"mother_full_name"`
	Email          string `form:"email" binding:"required"`
	PhoneNumber    string `form:"phone_number"`
}

func (studentRelative *StudentRelative) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = studentRelative.ID
	resp["student_id"] = studentRelative.StudentID
	resp["father_full_name"] = studentRelative.FatherFullName
	resp["mother_full_name"] = studentRelative.MotherFullName
	resp["email"] = studentRelative.Email
	resp["phone_number"] = studentRelative.PhoneNumber

	return resp
}
