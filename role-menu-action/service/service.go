package service

import (
	"com.mailnau.api/common"
	"com.mailnau.api/common/utils"
	"com.mailnau.api/config"
	"com.mailnau.api/role-menu-action/domain"
	"context"
	"github.com/rs/zerolog/log"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type service struct {
	cfg       config.Config
	repo      domain.Repository
	cacheRepo domain.CacheRepository
	f         utils.LogFormatter
}

func NewService(cfg config.Config, repo domain.Repository, cacheRepo domain.CacheRepository) domain.Service {
	f := utils.NewLogFormatter("role-menu-action.service")
	return &service{cfg: cfg, repo: repo, f: f, cacheRepo: cacheRepo}
}

func (s *service) GetRoleMenuActionByRoleID(ctx context.Context, roleID int) ([]domain.RoleMenuAction, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, s.f(utils.GetFN(s.GetRoleMenuActionByRoleID)))
	defer span.Finish()
	var roleMenuActions []domain.RoleMenuAction
	var err error

	roleMenuActions, err = s.cacheRepo.GetRoleMenuActionByRoleID(ctx, roleID)
	if err == nil {
		return roleMenuActions, nil
	}

	roleMenuActions, err = s.repo.FindRoleMenuActionByRoleID(ctx, roleID)
	if err != nil {
		return roleMenuActions, err
	}

	if errStore := s.cacheRepo.StoreRoleMenuAction(ctx, roleMenuActions, roleID); errStore != nil {
		log.Ctx(ctx).Err(err).Msg("error store roleMenuAction")
	}

	return roleMenuActions, nil
}

func (s *service) GetAllActions(ctx context.Context) (common.GeneralResponse, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, s.f(utils.GetFN(s.GetAllActions)))
	defer span.Finish()

	actions, err := s.cacheRepo.GetListAction(ctx)
	if err == nil {
		return common.GeneralResponse{Status: "success", Data: actions}, nil
	}

	actions, err = s.repo.FindAllActions(ctx)
	if err != nil {
		return common.GeneralResponse{}, err
	}

	var actionsDTO []domain.MenuOrActionDTO
	for _, action := range actions {
		actionsDTO = append(actionsDTO, domain.MenuOrActionDTO(action))
	}

	if err = s.cacheRepo.StoreListAction(ctx, actions); err != nil {
		return common.GeneralResponse{}, err
	}

	resp := common.GeneralResponse{Status: "success", Data: actionsDTO}

	return resp, nil
}

func (s *service) GetAllMenus(ctx context.Context) (common.GeneralResponse, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, s.f(utils.GetFN(s.GetAllMenus)))
	defer span.Finish()

	menus, err := s.cacheRepo.GetListMenu(ctx)
	if err == nil {
		return common.GeneralResponse{Status: "success", Data: menus}, nil
	}

	menus, err = s.repo.FindAllMenus(ctx)
	if err != nil {
		return common.GeneralResponse{}, err
	}

	var menusDTO []domain.MenuOrActionDTO
	for _, menu := range menus {
		menusDTO = append(menusDTO, domain.MenuOrActionDTO(menu))
	}

	if err = s.cacheRepo.StoreListMenu(ctx, menus); err != nil {
		return common.GeneralResponse{}, err
	}

	resp := common.GeneralResponse{Status: "success", Data: menusDTO}

	return resp, nil
}
