package db

import (
	"com.mailnau.api/common/utils"
	"com.mailnau.api/config"
	"com.mailnau.api/user/domain"
	"context"
	"fmt"
	"gopkg.in/jinzhu/gorm.v1"
)

type repository struct {
	cfg config.Config
	f   utils.LogFormatter
	*gorm.DB
}

func NewRepository(cfg config.Config, DB *gorm.DB) domain.Repository {
	f := utils.NewLogFormatter("user.repository")
	return &repository{cfg: cfg, DB: DB, f: f}
}

func (r *repository) StoreUserRole(ctx context.Context, role domain.UserRole) error {
	db := r.DB
	err := db.Create(&role).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) StoreUser(ctx context.Context, model domain.User) (domain.User, error) {
	db := r.DB
	err := db.Create(&model).Error
	if err != nil {
		return domain.User{}, err
	}

	return model, nil
}

func (r *repository) FindUserByEmail(ctx context.Context, email string) (domain.User, error) {

	var result domain.User
	db := r.DB
	if err := db.Where(&domain.User{Email: email}).
		Find(&result).Error; err != nil {
		msg := fmt.Errorf("cannot find user: error=%s, username=%s", err, email)
		fmt.Println(msg)
		return result, err
	}

	return result, nil
}

func (r *repository) FindUserByUsernameAndPassword(ctx context.Context, email, password string) (*domain.User, error) {

	db := r.DB
	var result domain.User
	if err := db.Where(&domain.User{Email: email, Password: password}).
		Find(&result).Error; err != nil {
		msg := fmt.Errorf("cannot find user: error=%s, username=%s", err, email)
		fmt.Println(msg)
		return nil, err
	}

	return &result, nil
}
