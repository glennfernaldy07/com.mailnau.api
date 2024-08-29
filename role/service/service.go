package service

import (
	"context"

	"com.mailnau.api/common"
	"com.mailnau.api/common/utils"
	"com.mailnau.api/config"
	"com.mailnau.api/role/domain"
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

func (s *service) AddRole(
	ctx context.Context,
	name, echelon string,
	menuIDs, actionIDs []int64,
) (*common.BaseResponse[domain.CreateRoleResponse], error) {
	span, ctx := tracer.StartSpanFromContext(ctx, s.f(utils.GetFN(s.AddRole)))
	defer span.Finish()

	roleID, err := s.repo.CreateRole(ctx, name, echelon, menuIDs, actionIDs)
	if err != nil {
		return nil, err
	}

	resp := common.NewBaseResponse(
		"Role baru berhasil ditambahkan",
		domain.CreateRoleResponse{RoleId: roleID},
		nil,
	)

	return resp, nil
}
