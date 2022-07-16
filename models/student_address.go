package models

import "time"

type StudentAddress struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	StudentID int       `json:"student_id"`
	Students  Students  `gorm:"foreignKey:StudentID"`
	Country   string    `gorm:"country"`
	State     string    `gorm:"state"`
	City      string    `gorm:"city"`
	Street    string    `gorm:"street"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (studentAddress *StudentAddress) TableName() string {
	return "StudentAddress"
}

type RecordStudentAddress struct {
	StudentID int    `form:"student_id" binding:"required"`
	Country   string `form:"country"`
	State     string `form:"state"`
	City      string `form:"city"`
	Street    string `form:"street"`
}

type UpdateStudentAddress struct {
	ID      int    `form:"id" binding:"required"`
	Country string `form:"country"`
	State   string `form:"state"`
	City    string `form:"city"`
	Street  string `form:"street"`
}

func (studentAddress *StudentAddress) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = studentAddress.ID
	resp["student_id"] = studentAddress.StudentID
	resp["country"] = studentAddress.Country
	resp["state"] = studentAddress.State
	resp["city"] = studentAddress.City

	return resp
}
