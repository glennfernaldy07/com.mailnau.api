package service

import (
	"com.mailnau.api/attendance/domain"
	"com.mailnau.api/common"
	comdb "com.mailnau.api/common/db"
	cerr "com.mailnau.api/common/errors"
	"com.mailnau.api/common/utils"
	"com.mailnau.api/config"
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"gopkg.in/jinzhu/gorm.v1"
	"net/http"
	"time"
)

type service struct {
	cfg       config.Config
	repo      domain.Repository
	cacheRepo domain.CacheRepository
	f         utils.LogFormatter
}

func NewService(cfg config.Config, repo domain.Repository, cacheRepo domain.CacheRepository) domain.Service {
	f := utils.NewLogFormatter("attendance.service")
	return &service{cfg: cfg, repo: repo, cacheRepo: cacheRepo, f: f}
}

func (s *service) DoCheckIn(ctx context.Context, req domain.AttendanceRequest, claims jwt.MapClaims) (common.GeneralResponse, error) {
	resp := domain.AttendanceResponse{}

	//Decode qr content base 64
	var content domain.QrContentValue
	if err := utils.DecodeFromBase64(&content, req.QrContent); err != nil {
		return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusInternalServerError, err.Error(), err)
	}
	if !s.validate(req, content, claims) {
		err := fmt.Errorf("invalid content")
		return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusUnauthorized, err.Error(), err)
	}

	resp = domain.AttendanceResponse{
		RequestType: req.RequestType,
		UserID:      content.UserID,
		Timestamp:   content.Timestamp,
		Location:    content.Location,
	}
	//Get attendance table by user uuid
	attendanceModel, err := s.GetAttendanceByUserID(ctx, content.UserID)
	//if not found then insert
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//insert attendance log table
		newAttendanceModel := domain.Attendance{
			UserID: content.UserID,
			Status: domain.StatusIn,
			Base: comdb.Base{
				CreatedAt: time.Now(),
				CreatedBy: "SYSTEM",
				UpdatedAt: time.Now(),
				UpdatedBy: "SYSTEM",
			},
		}
		if err := s.InsertAttendance(ctx, newAttendanceModel); err != nil {
			return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusInternalServerError, err.Error(), err)
		}
		attendanceLogModel := domain.AttendanceLog{
			AttendanceID: newAttendanceModel.ID.String(),
			Action:       domain.StatusIn,
			Location:     content.Location,
			Base:         comdb.Base{},
		}
		if errInsert := s.InsertAttendanceLog(ctx, attendanceLogModel); errInsert != nil {
			return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusInternalServerError, errInsert.Error(), errInsert)
		}
		return common.GeneralResponse{Status: "200", Message: common.SuccessMessage, Data: resp}, nil
	} else if err != nil {
		return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusInternalServerError, err.Error(), err)
	}
	//update attendance table status to checked in
	attendanceModel.Status = domain.StatusIn
	attendanceModel.UpdatedAt = time.Now()
	errUpdate := s.UpdateAttendanceByAttendanceID(ctx, *attendanceModel)
	if errUpdate != nil {
		return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusInternalServerError, errUpdate.Error(), errUpdate)
	}
	//INSERT LOG
	attendanceLogModel := domain.AttendanceLog{
		AttendanceID: attendanceModel.ID.String(),
		Action:       domain.StatusIn,
		Base: comdb.Base{
			CreatedAt: time.Now(),
			CreatedBy: "SYSTEM",
			UpdatedAt: time.Now(),
			UpdatedBy: "SYSTEM",
		},
	}
	if err := s.InsertAttendanceLog(ctx, attendanceLogModel); err != nil {
		return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusInternalServerError, err.Error(), err)
	}
	// return token
	return common.GeneralResponse{Status: "200", Message: common.SuccessMessage, Data: resp}, nil
}

func (s *service) DoCheckOut(ctx context.Context, req domain.AttendanceRequest, claims jwt.MapClaims) (common.GeneralResponse, error) {
	resp := domain.AttendanceResponse{}

	//Decode qr content base 64
	var content domain.QrContentValue
	if err := utils.DecodeFromBase64(&content, req.QrContent); err != nil {
		return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusInternalServerError, err.Error(), err)
	}
	if !s.validate(req, content, claims) {
		err := fmt.Errorf("invalid content")
		return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusUnauthorized, err.Error(), err)
	}

	resp = domain.AttendanceResponse{
		RequestType: req.RequestType,
		UserID:      content.UserID,
		Timestamp:   content.Timestamp,
		Location:    content.Location,
	}
	//Get attendance table by user uuid
	attendanceModel, err := s.GetAttendanceByUserID(ctx, content.UserID)
	//if not found then insert
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//insert attendance log table
		newAttendanceModel := domain.Attendance{
			UserID: content.UserID,
			Status: domain.StatusOut,
			Base: comdb.Base{
				CreatedAt: time.Now(),
				CreatedBy: "SYSTEM",
				UpdatedAt: time.Now(),
				UpdatedBy: "SYSTEM",
			},
		}
		if err := s.InsertAttendance(ctx, newAttendanceModel); err != nil {
			return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusInternalServerError, err.Error(), err)
		}
		attendanceLogModel := domain.AttendanceLog{
			AttendanceID: newAttendanceModel.ID.String(),
			Action:       domain.StatusOut,
			Location:     content.Location,
			Base:         comdb.Base{},
		}
		if errInsert := s.InsertAttendanceLog(ctx, attendanceLogModel); errInsert != nil {
			return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusInternalServerError, errInsert.Error(), errInsert)
		}
		return common.GeneralResponse{Status: "200", Message: common.SuccessMessage, Data: resp}, nil
	} else if err != nil {
		return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusInternalServerError, err.Error(), err)
	}
	//update attendance table status to checked in
	attendanceModel.Status = domain.StatusOut
	attendanceModel.UpdatedAt = time.Now()
	errUpdate := s.UpdateAttendanceByAttendanceID(ctx, *attendanceModel)
	if errUpdate != nil {
		return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusInternalServerError, errUpdate.Error(), errUpdate)
	}
	//INSERT LOG
	attendanceLogModel := domain.AttendanceLog{
		AttendanceID: attendanceModel.ID.String(),
		Action:       domain.StatusOut,
		Base: comdb.Base{
			CreatedAt: time.Now(),
			CreatedBy: "SYSTEM",
			UpdatedAt: time.Now(),
			UpdatedBy: "SYSTEM",
		},
	}
	if err := s.InsertAttendanceLog(ctx, attendanceLogModel); err != nil {
		return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusInternalServerError, err.Error(), err)
	}
	// return token
	return common.GeneralResponse{Status: "200", Message: common.SuccessMessage, Data: resp}, nil
}

func (s *service) GetAttendanceByUserID(ctx context.Context, userID string) (*domain.Attendance, error) {
	attendanceModel, err := s.repo.FindAttendanceByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return attendanceModel, nil
}

func (s *service) InsertAttendance(ctx context.Context, model domain.Attendance) error {
	if err := s.repo.StoreAttendance(ctx, model); err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateAttendanceByAttendanceID(ctx context.Context, model domain.Attendance) error {
	if err := s.repo.UpdateAttendanceByAttendanceID(ctx, model); err != nil {
		return err
	}
	return nil
}

func (s *service) InsertAttendanceLog(ctx context.Context, model domain.AttendanceLog) error {
	if err := s.repo.StoreAttendanceLog(ctx, model); err != nil {
		return err
	}
	return nil
}

func (s *service) validate(req domain.AttendanceRequest, content domain.QrContentValue, claims jwt.MapClaims) bool {
	//validate data
	if req.UserID != content.UserID {
		return false
	}
	userIDClaims := claims["key"].(string)
	if userIDClaims == "" {
		return false
	}
	if userIDClaims != content.UserID {
		return false
	}
	return true
}
