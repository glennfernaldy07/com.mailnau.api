package mysql

import (
	"context"

	errs "com.mailnau.api/common/errors"
	"com.mailnau.api/common/utils"
	"com.mailnau.api/config"
	rma_domain "com.mailnau.api/role-menu-action/domain"
	"com.mailnau.api/role/domain"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/jinzhu/gorm.v1"
)

type repository struct {
	cfg config.Config
	f   utils.LogFormatter
	*gorm.DB
}

func NewRepository(cfg config.Config, DB *gorm.DB) domain.Repository {
	f := utils.NewLogFormatter("role.repository")
	return &repository{cfg: cfg, DB: DB, f: f}
}

func (r *repository) CreateRole(ctx context.Context, name, echelon string, menuIDs, actionIDs []int64) (int64, error) {
	span, _ := tracer.StartSpanFromContext(ctx, r.f(utils.GetFN(r.CreateRole)))
	defer span.Finish()

	tx := r.DB.Begin()

	role := domain.Role{
		Name:    name,
		Echelon: echelon,
	}
	if err := tx.Create(&role).Error; err != nil {
		tx.Rollback()
		return -1, errs.NewBadRequestError("Gagal membuat role", err)
	}

	// Create role-menu-action associations
	for _, menuID := range menuIDs {
		for _, actionID := range actionIDs {
			rma := rma_domain.RoleMenuAction{
				RoleID:   role.ID,
				MenuID:   menuID,
				ActionID: actionID,
			}
			if err := tx.Create(&rma).Error; err != nil {
				tx.Rollback()
				return -1, errs.NewInternalError(err)
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		return -1, errs.NewInternalError(err)
	}

	return role.ID, nil
}

func (r *repository) FindRolesWithMenuAndActions(ctx context.Context, limit int, offset int) ([]domain.Role, error) {
	span, _ := tracer.StartSpanFromContext(ctx, r.f(utils.GetFN(r.FindRolesWithMenuAndActions)))
	defer span.Finish()

	var roles []domain.Role
	if err := r.DB.
		Preload("Menus").
		Preload("Actions").
		Limit(limit).Offset(offset).
		Find(&roles).
		Error; err != nil {
		return nil, errs.NewInternalError(err)
	}
	return roles, nil
}

func (r *repository) CountRolesRecords(ctx context.Context) (int64, error) {
	span, _ := tracer.StartSpanFromContext(ctx, r.f(utils.GetFN(r.CountRolesRecords)))
	defer span.Finish()

	var totalRecords int64
	if err := r.DB.Model(&domain.Role{}).Count(&totalRecords).Error; err != nil {
		return -1, errs.NewInternalError(err)
	}
	return totalRecords, nil
}
