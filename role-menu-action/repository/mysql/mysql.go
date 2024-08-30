package mysql

import (
	"context"

	errs "com.mailnau.api/common/errors"
	"com.mailnau.api/common/utils"
	"com.mailnau.api/config"
	"com.mailnau.api/role-menu-action/domain"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/jinzhu/gorm.v1"
)

type repository struct {
	cfg config.Config
	f   utils.LogFormatter
	*gorm.DB
}

func NewRepository(cfg config.Config, DB *gorm.DB) domain.Repository {
	f := utils.NewLogFormatter("role-menu-action.repository")
	return &repository{cfg: cfg, DB: DB, f: f}
}

func (r *repository) FindAllActions(ctx context.Context) ([]domain.Action, error) {
	span, _ := tracer.StartSpanFromContext(ctx, r.f(utils.GetFN(r.FindAllActions)))
	defer span.Finish()

	var actions []domain.Action
	if err := r.DB.
		Find(&actions).
		Error; err != nil {
		return nil, errs.NewInternalError(err)
	}
	return actions, nil
}

func (r *repository) FindAllMenus(ctx context.Context) ([]domain.Menu, error) {
	span, _ := tracer.StartSpanFromContext(ctx, r.f(utils.GetFN(r.FindAllMenus)))
	defer span.Finish()

	var menus []domain.Menu
	if err := r.DB.
		Find(&menus).
		Error; err != nil {
		return nil, errs.NewInternalError(err)
	}
	return menus, nil
}
