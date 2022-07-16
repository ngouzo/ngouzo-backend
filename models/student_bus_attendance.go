package models

import "time"

type StudentBusAttendance struct {
	ID            int64     `gorm:"primary_key;auto_increment"  json:"id"`
	StudentID     int       `json:"student_id"`
	Students      Students  `gorm:"foreignKey:StudentID"`
	PickStudent   bool      `json:"pick_student"`
	ArriveStudent string    `json:"arrive_school"`
	ArrivedAt     time.Time `json:"arrived_at"`
	PickedAt      time.Time `json:"picked_at"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

func (studentBusAttendance *StudentBusAttendance) TableName() string {
	return "StudentBusAttendance"
}

type RecordStudentBusAttendance struct {
	StudentID     int       `form:"student_id" binding:"required"`
	PickStudent   bool      `form:"pick_student" binding:"required"`
	ArriveStudent bool      `form:"arrive_school" binding:"required"`
	PickedAt      time.Time `json:"picked_at"`
}

type UpdateStudentBusAttendance struct {
	ID            int64     `form:"id" binding:"required"`
	PickStudent   bool      `form:"pick_student"`
	ArriveStudent bool      `form:"arrive_school"`
	ArrivedAt     time.Time `json:"arrived_at"`
	PickedAt      time.Time `json:"picked_at"`
}

func (studentBusAttendance *StudentBusAttendance) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = studentBusAttendance.ID
	resp["pick_student"] = studentBusAttendance.PickStudent
	resp["arrive_school"] = studentBusAttendance.ArriveStudent
	resp["student_id"] = studentBusAttendance.StudentID
	resp["arrived_at"] = studentBusAttendance.ArrivedAt
	resp["pick_at"] = studentBusAttendance.PickedAt

	return resp
}
