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
