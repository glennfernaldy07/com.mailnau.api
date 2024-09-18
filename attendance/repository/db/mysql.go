package db

import (
	"com.mailnau.api/attendance/domain"
	"com.mailnau.api/common/utils"
	"com.mailnau.api/config"
	"context"
	"fmt"
	"gopkg.in/jinzhu/gorm.v1"
)

type repository struct {
	cfg config.Config
	f   utils.LogFormatter
	*gorm.DB
}

func (r *repository) FindAttendanceByUserID(ctx context.Context, userID string) (*domain.Attendance, error) {
	var result domain.Attendance
	db := r.DB
	if err := db.Where(&domain.Attendance{UserID: userID}).
		Find(&result).Error; err != nil {
		msg := fmt.Errorf("cannot find user: error=%s, user_id=%s", err, userID)
		fmt.Println(msg)
		return &result, err
	}

	return &result, nil
}

func (r *repository) StoreAttendance(ctx context.Context, model domain.Attendance) error {
	db := r.DB
	err := db.Create(&model).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateAttendanceByAttendanceID(ctx context.Context, model domain.Attendance) error {
	db := r.DB
	updatedModel := domain.Attendance{
		Status: model.Status,
	}
	row := db.Model(&domain.Attendance{}).
		Where("id = ?", model.ID).
		Updates(updatedModel)
	if row.Error != nil {
		return row.Error
	}
	return nil
}

func (r *repository) StoreAttendanceLog(ctx context.Context, model domain.AttendanceLog) error {
	db := r.DB
	err := db.Create(&model).Error
	if err != nil {
		return err
	}

	return nil
}

func NewRepository(cfg config.Config, DB *gorm.DB) domain.Repository {
	f := utils.NewLogFormatter("user.repository")
	return &repository{cfg: cfg, DB: DB, f: f}
}
