package service

import (
	"context"
	"strings"

	"net/http"

	"com.mailnau.api/common"
	cerr "com.mailnau.api/common/errors"
	"com.mailnau.api/common/utils"
	"com.mailnau.api/config"
	rmad "com.mailnau.api/role-menu-action/domain"
	"com.mailnau.api/role/domain"
	"github.com/rs/zerolog/log"
)

type service struct {
	cfg       config.Config
	repo      domain.Repository
	cacheRepo domain.CacheRepository
	f         utils.LogFormatter
}

func (s *service) AddRole(ctx context.Context, name string, echelon string, menuIDs []string, actionIDs []string) (common.GeneralResponse, error) {
	roleID, err := s.repo.CreateRole(ctx, name, echelon, menuIDs, actionIDs)
	if err != nil {
		if strings.HasPrefix(err.Error(), "Error 1062") {
			return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusBadRequest, "Role sudah pernah dibuat", err)
		}
		return common.GeneralResponse{}, err
	}
	data := map[string]int{}
	data["roleId"] = roleID
	resp := common.GeneralResponse{Status: "success", Message: "Role berhasil ditambahkan", Data: data}
	return resp, nil
}

func (s *service) GetRolesWithMenuAndActions(ctx context.Context, limit int, page int) (common.GeneralResponse, error) {
	offset := (page - 1) * limit

	totalRecords, err := s.repo.CountRolesRecords(ctx)
	if err != nil {
		return common.GeneralResponse{}, err
	}

	roles, err := s.repo.FindRolesWithMenuAndActions(ctx, limit, offset)
	if err != nil {
		return common.GeneralResponse{}, err
	}

	var result []domain.RoleDTO
	for _, role := range roles {
		var menus []rmad.MenuOrActionDTO
		for _, menu := range role.Menus {
			menus = append(menus, rmad.MenuOrActionDTO(menu))
		}

		var actions []rmad.MenuOrActionDTO
		for _, action := range role.Actions {
			actions = append(actions, rmad.MenuOrActionDTO(action))
		}

		roleDTO := domain.RoleDTO{
			ID:       role.ID,
			RoleName: role.RoleName,
			Menus:    menus,
			Echelon:  role.Eselon,
			Actions:  actions,
		}

		result = append(result, roleDTO)
	}

	totalPages := int((totalRecords + int64(limit) - 1) / int64(limit))

	pagination := common.PageMetaResponse{
		Current:      page,
		TotalPages:   totalPages,
		PerPage:      limit,
		TotalRecords: totalRecords,
	}

	resp := common.GeneralResponse{
		Status: "success",
		Data:   result,
		Meta:   pagination,
	}

	return resp, nil
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
