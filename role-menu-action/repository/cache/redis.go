package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"com.mailnau.api/common/utils"
	"com.mailnau.api/config"
	"com.mailnau.api/role-menu-action/domain"
	"github.com/go-redis/redis/v8"
)

const (
	expiredSecond = 1800

	actionListKey = "action_list"
	menuListKey   = "menu_list"
)

type repository struct {
	client redis.UniversalClient
	cfg    config.Config
	f      utils.LogFormatter
}

func (r *repository) GetListAction(ctx context.Context) ([]domain.Action, error) {
	res, err := r.client.Get(ctx, actionListKey).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			fmt.Printf("cannot get from redis: key=%s", actionListKey)
		}
		return nil, err
	}
	var listAction []domain.Action
	err = json.Unmarshal([]byte(res), &listAction)
	if err != nil {
		return nil, err
	}
	return listAction, nil
}

func (r *repository) GetListMenu(ctx context.Context) ([]domain.Menu, error) {
	res, err := r.client.Get(ctx, menuListKey).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			fmt.Printf("cannot get from redis: key=%s", menuListKey)
		}
		return nil, err
	}
	var listMenu []domain.Menu
	err = json.Unmarshal([]byte(res), &listMenu)
	if err != nil {
		return nil, err
	}
	return listMenu, nil
}

func (r *repository) StoreListAction(ctx context.Context, listAction []domain.Action) error {
	data, err := json.Marshal(listAction)
	if err != nil {
		return err
	}
	_, err = r.client.SetNX(ctx, actionListKey, data, time.Duration(expiredSecond)*time.Second).Result()
	if err != nil {
		fmt.Printf("cannot run SetNX: data=%+v", listAction)
		return err
	}
	return nil
}

func (r *repository) StoreListMenu(ctx context.Context, listMenu []domain.Menu) error {
	data, err := json.Marshal(listMenu)
	if err != nil {
		return err
	}
	_, err = r.client.SetNX(ctx, menuListKey, data, time.Duration(expiredSecond)*time.Second).Result()
	if err != nil {
		fmt.Printf("cannot run SetNX: data=%+v", listMenu)
		return err
	}
	return nil
}

func NewRepository(cfg config.Config) domain.CacheRepository {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.GetString(config.RedisAddress),
		DB:   int(cfg.GetInt(config.RedisMasterDB)),
	})

	f := utils.NewLogFormatter("role-menu-action.repository.cache")
	return &repository{rdb, cfg, f}
}
