package service

import (
	"context"

	"com.mailnau.api/common"
	"com.mailnau.api/common/utils"
	"com.mailnau.api/config"
	"com.mailnau.api/role-menu-action/domain"
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

func (s *service) GetAllActions(ctx context.Context) (common.GeneralResponse, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, s.f(utils.GetFN(s.GetAllActions)))
	defer span.Finish()

	actions, err := s.repo.FindAllActions(ctx)
	if err != nil {
		return common.GeneralResponse{}, err
	}

	var actionsDTO []domain.MenuOrActionDTO
	for _, action := range actions {
		actionsDTO = append(actionsDTO, domain.MenuOrActionDTO(action))
	}

	resp := common.GeneralResponse{Status: "success", Data: actionsDTO}

	return resp, nil
}

func (s *service) GetAllMenus(ctx context.Context) (common.GeneralResponse, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, s.f(utils.GetFN(s.GetAllMenus)))
	defer span.Finish()

	actions, err := s.repo.FindAllMenus(ctx)
	if err != nil {
		return common.GeneralResponse{}, err
	}

	var menusDTO []domain.MenuOrActionDTO
	for _, menu := range actions {
		menusDTO = append(menusDTO, domain.MenuOrActionDTO(menu))
	}

	resp := common.GeneralResponse{Status: "success", Data: menusDTO}

	return resp, nil
}
