package mysql

import (
	"context"
	"fmt"

	"com.mailnau.api/common/utils"
	"com.mailnau.api/config"
	"com.mailnau.api/role-menu-action/domain"
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
	var actions []domain.Action
	if err := r.DB.
		Find(&actions).
		Error; err != nil {
		msg := fmt.Errorf("cannot find all actions: error=%s", err)
		fmt.Println(msg)
		return nil, err
	}
	return actions, nil
}

func (r *repository) FindAllMenus(ctx context.Context) ([]domain.Menu, error) {
	var menus []domain.Menu
	if err := r.DB.
		Find(&menus).
		Error; err != nil {
		msg := fmt.Errorf("cannot find all menus: error=%s", err)
		fmt.Println(msg)
		return nil, err
	}
	return menus, nil
}
