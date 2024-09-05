package service

import (
	"com.mailnau.api/common/utils"
	"com.mailnau.api/config"
	"com.mailnau.api/role/domain"
	"context"
	"github.com/rs/zerolog/log"
)

type service struct {
	cfg       config.Config
	repo      domain.Repository
	cacheRepo domain.CacheRepository
	f         utils.LogFormatter
}

func (s *service) GetListMenuByRoleID(ctx context.Context, roleID int) ([]string, error) {
	var menus []string
	var err error
	if menus, err = s.cacheRepo.GetListMenuByRoleID(ctx, roleID); err == nil {
		return []string{}, nil
	}

	roleMenu, err := s.repo.FindRoleMenuByRoleID(ctx, roleID)
	if err != nil {
		return []string{}, nil
	}
	for _, menu := range roleMenu {
		menus = append(menus, menu.MenuID)
	}

	if err := s.cacheRepo.StoreListMenuByRoleID(ctx, roleID, menus); err != nil {
		log.Ctx(ctx).Err(err).Msg("store role menu")
	}
	return menus, nil
}

func (s *service) GetRoleByID(ctx context.Context, id int) (domain.Role, error) {
	if role, err := s.cacheRepo.GetRoleByID(ctx, id); err == nil {
		return role, nil
	}

	role, err := s.repo.FindRoleByID(ctx, id)
	if err != nil {
		return domain.Role{}, nil
	}

	if err := s.cacheRepo.StoreRoleByID(ctx, id, role); err != nil {
		log.Ctx(ctx).Err(err).Msg("store role")
	}
	return role, nil
}

func NewService(cfg config.Config, repo domain.Repository, cacheRepo domain.CacheRepository) domain.Service {
	f := utils.NewLogFormatter("role.service")
	return &service{cfg: cfg, repo: repo, cacheRepo: cacheRepo, f: f}
}
