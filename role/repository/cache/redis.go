package cache

import (
	"com.mailnau.api/common/utils"
	"com.mailnau.api/config"
	"com.mailnau.api/role/domain"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"strings"
	"time"
)

const (
	roleByIDKey       = "role_by_id_"
	roleExpiredSecond = 1800

	roleMenuByRoleIDKey = "role_menu_by_role_id_"
)

type repository struct {
	client redis.UniversalClient
	cfg    config.Config
	f      utils.LogFormatter
}

func (r *repository) StoreRoleByID(ctx context.Context, id int, role domain.Role) error {
	key := roleByIDKey + strconv.Itoa(id)

	b, err := json.Marshal(role)
	if err != nil {
		fmt.Printf("cannot marshal json: role_id=%d, dt=%+v", id, role)
		return err
	}

	_, err = r.client.SetNX(ctx, key, string(b), time.Duration(roleExpiredSecond)*time.Second).Result()
	if err != nil {
		fmt.Printf("cannot run SetNX: roleID=%d, data=%+v", id, role)
		return err
	}

	return nil
}

func (r *repository) GetRoleByID(ctx context.Context, id int) (domain.Role, error) {

	key := roleByIDKey + strconv.Itoa(id)

	res, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			fmt.Printf("cannot get from redis: key=%s", key)
		}

		return domain.Role{}, err
	}

	var dt *domain.Role
	err = json.Unmarshal([]byte(res), &dt)
	if err != nil {
		return domain.Role{}, err
	}

	return *dt, nil
}

func (r *repository) StoreListMenuByRoleID(ctx context.Context, roleID int, listMenu []string) error {
	key := roleMenuByRoleIDKey + strconv.Itoa(roleID)

	_, err := r.client.SetNX(ctx, key, listMenu, time.Duration(roleExpiredSecond)*time.Second).Result()
	if err != nil {
		fmt.Printf("cannot run SetNX: roleID=%d, data=%+v", roleID, listMenu)
		return err
	}

	return nil
}

func (r *repository) GetListMenuByRoleID(ctx context.Context, id int) ([]string, error) {

	key := roleMenuByRoleIDKey + strconv.Itoa(id)

	res, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			fmt.Printf("cannot get from redis: key=%s", key)
		}

		return []string{}, err
	}
	resp := strings.Split(res, ",")
	return resp, nil
}

func NewRepository(cfg config.Config) domain.CacheRepository {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.GetString(config.RedisAddress),
		DB:   int(cfg.GetInt(config.RedisMasterDB)),
	})

	f := utils.NewLogFormatter("role.repository.cache")
	return &repository{rdb, cfg, f}
}
