package db

import (
	"com.mailnau.api/common/utils"
	"com.mailnau.api/config"
	"com.mailnau.api/role/domain"
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"gopkg.in/jinzhu/gorm.v1"
)

type repository struct {
	cfg config.Config
	f   utils.LogFormatter
	*gorm.DB
}

func (r *repository) FindRoleByID(ctx context.Context, id int) (domain.Role, error) {
	var result domain.Role
	db := r.DB
	if err := db.Where(&domain.Role{ID: id}).
		Find(&result).Error; err != nil {
		msg := fmt.Errorf("cannot find role: error=%s, role_id=%d", err, id)
		log.Ctx(ctx).Error().Str("module", "mysql").Msg(msg.Error())
		return result, err
	}

	return result, nil
}

func (r *repository) FindRoleMenuByRoleID(ctx context.Context, roleID int) ([]domain.RoleMenu, error) {
	var results []domain.RoleMenu
	db := r.DB
	if err := db.Where(&domain.RoleMenu{RoleID: int64(roleID)}).
		Find(&results).Error; err != nil {
		msg := fmt.Errorf("cannot find role_menu: error=%s, role_id=%d", err, roleID)
		log.Ctx(ctx).Error().Str("module", "mysql").Msg(msg.Error())
		return results, err
	}

	return results, nil
}

func NewRepository(cfg config.Config, DB *gorm.DB) domain.Repository {
	f := utils.NewLogFormatter("user.repository")
	return &repository{cfg: cfg, DB: DB, f: f}
}

//
//func (r *repository) StoreUser(ctx context.Context, model domain.User) error {
//	db := r.DB
//	err := db.Create(&model).Error
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (r *repository) FindUserByEmail(ctx context.Context, email string) (domain.User, error) {
//
//	var result domain.User
//	db := r.DB
//	if err := db.Where(&domain.User{Email: email}).
//		Find(&result).Error; err != nil {
//		msg := fmt.Errorf("cannot find user: error=%s, username=%s", err, email)
//		fmt.Println(msg)
//		return result, err
//	}
//
//	return result, nil
//}
//
//func (r *repository) FindUserByUsernameAndPassword(ctx context.Context, email, password string) (*domain.User, error) {
//
//	db := r.DB
//	var result domain.User
//	if err := db.Where(&domain.User{Email: email, Password: password}).
//		Find(&result).Error; err != nil {
//		msg := fmt.Errorf("cannot find user: error=%s, username=%s", err, email)
//		fmt.Println(msg)
//		return nil, err
//	}
//
//	return &result, nil
//}
