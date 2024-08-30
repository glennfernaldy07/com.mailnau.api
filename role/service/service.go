package service

import (
	"context"

	"com.mailnau.api/common"
	"com.mailnau.api/common/utils"
	"com.mailnau.api/config"
	rmad "com.mailnau.api/role-menu-action/domain"
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

func (s *service) GetRolesWithMenuAndActions(ctx context.Context, limit, page int) (*common.BaseResponse[[]domain.RoleDTO], error) {
	span, ctx := tracer.StartSpanFromContext(ctx, s.f(utils.GetFN(s.GetRolesWithMenuAndActions)))
	defer span.Finish()

	offset := (page - 1) * limit

	totalRecords, err := s.repo.CountRolesRecords(ctx)
	if err != nil {
		return nil, err
	}

	roles, err := s.repo.FindRolesWithMenuAndActions(ctx, limit, offset)
	if err != nil {
		return nil, err
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
			RoleName: role.Name,
			Menus:    menus,
			Echelon:  role.Echelon,
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

	resp := common.NewBaseResponse(
		"Berhasil",
		result,
		pagination,
	)

	return resp, nil
}
