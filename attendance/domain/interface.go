package domain

import (
	"com.mailnau.api/common"
	"context"
	"github.com/golang-jwt/jwt/v4"
)

type Service interface {
	DoCheckIn(ctx context.Context, req AttendanceRequest, claims jwt.MapClaims) (common.GeneralResponse, error)
	DoCheckOut(ctx context.Context, req AttendanceRequest, claims jwt.MapClaims) (common.GeneralResponse, error)

	GetAttendanceByUserID(ctx context.Context, userID string) (*Attendance, error)
	InsertAttendance(ctx context.Context, model Attendance) error
	UpdateAttendanceByAttendanceID(ctx context.Context, model Attendance) error
	InsertAttendanceLog(ctx context.Context, model AttendanceLog) error
}

type Repository interface {
	FindAttendanceByUserID(ctx context.Context, userID string) (*Attendance, error)
	StoreAttendance(ctx context.Context, model Attendance) error
	UpdateAttendanceByAttendanceID(ctx context.Context, model Attendance) error
	StoreAttendanceLog(ctx context.Context, model AttendanceLog) error
	//
	//StoreUserRole(ctx context.Context, role UserRole) error
}

type CacheRepository interface {
	StoreAccessToken(ctx context.Context, userID string, token string) error
	//GetAccessToken(ctx context.Context, accessToken string) (bool, error)
}
