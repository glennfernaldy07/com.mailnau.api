package service

import (
	"com.mailnau.api/common/utils"
	"com.mailnau.api/config"
	"com.mailnau.api/user/domain"
	"context"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type service struct {
	cfg  config.Config
	repo domain.Repository
	f    utils.LogFormatter
}

func NewService(cfg config.Config, repo domain.Repository) domain.Service {
	f := utils.NewLogFormatter("user.service")
	return &service{cfg: cfg, repo: repo, f: f}
}

func (s *service) GetUserByUsernameAndPassword(ctx context.Context, username, password string) (*domain.UserResponse, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, s.f(utils.GetFN(s.GetUserByUsernameAndPassword)))
	defer span.Finish()
	resp := domain.UserResponse{}
	userModel, err := s.repo.FindUserByUsernameAndPassword(ctx, username, password)
	if err != nil {
		return nil, err
	}
	if userModel.ID > 0 {
		resp.ResponseMessage = "Success"
		resp.ResponseCode = "00"
	}

	return &resp, nil
}
