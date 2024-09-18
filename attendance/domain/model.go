package domain

import (
	comdb "com.mailnau.api/common/db"
)

const (
	StatusIn  = "IN"
	StatusOut = "OUT"
)

type Attendance struct {
	UserID string `gorm:"column:user_id"`
	Status string `gorm:"column:status"` //in or out default out

	comdb.Base
}

func (Attendance) TableName() string {
	return "attendance"
}

type AttendanceLog struct {
	AttendanceID string `gorm:"column:attendance_id"`
	Action       string `gorm:"column:action"`
	Location     string `gorm:"column:location"`
	Image        string `gorm:"column:image"`
	comdb.Base
}

func (AttendanceLog) TableName() string {
	return "attendance_logs"
}

type AttendanceRequest struct {
	QrContent   string `json:"qr_content" validate:"required"`
	RequestType string `json:"request_type" validate:"required"` //in or out
	UserID      string
}

type AttendanceResponse struct {
	RequestType string `json:"request_type"`
	UserID      string `json:"user_id"`
	Timestamp   string `json:"timestamp"`
	Location    string `json:"location"`
}

type QrContentValue struct {
	UserID    string `json:"user_id" validate:"required"` // ganti jadi qr content
	Timestamp string `json:"timestamp" validate:"required"`
	Location  string `json:"location,omitempty"`
}
