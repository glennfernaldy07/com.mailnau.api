package cache

import (
	"com.mailnau.api/attendance/domain"
	"com.mailnau.api/common/utils"
	"com.mailnau.api/config"
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type repository struct {
	client redis.UniversalClient
	cfg    config.Config
	f      utils.LogFormatter
}

func NewRepository(cfg config.Config) domain.CacheRepository {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.GetString(config.RedisAddress),
		DB:   int(cfg.GetInt(config.RedisMasterDB)),
	})

	f := utils.NewLogFormatter("user.repository.cache")
	return &repository{rdb, cfg, f}
}

func (r *repository) StoreAccessToken(ctx context.Context, userID string, token string) error {
	key := token

	expiry := r.cfg.GetInt(config.TokenExpTimeSecond)

	_, err := r.client.SetNX(ctx, key, 1, time.Duration(expiry)*time.Second).Result()
	if err != nil {
		fmt.Printf("cannot run SetNX: userID=%s, data=%s", userID, token)

		return err
	}

	return nil
}

func (r *repository) GetAccessToken(ctx context.Context, accessToken string) (bool, error) {

	key := accessToken

	_, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			fmt.Printf("cannot get from redis: key=%s", key)
		}

		return false, err
	}

	return true, nil
}
