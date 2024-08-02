package mysql

import (
	"com.mailnau.api/common/utils"
	"com.mailnau.api/config"
	"com.mailnau.api/user/domain"
	"context"
	"fmt"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
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

func (r *repository) FindUserByUsernameAndPassword(ctx context.Context, username, password string) (*domain.User, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, r.f(utils.GetFN(r.FindUserByUsernameAndPassword)))
	defer span.Finish()

	db := r.DB
	var result domain.User
	if err := db.Where(&domain.User{Username: username, Password: password}).
		Find(&result).Error; err != nil {
		msg := fmt.Errorf("cannot find user: error=%s, username=%s", err, username)
		fmt.Println(msg)
		return nil, err
	}

	return &result, nil
}
