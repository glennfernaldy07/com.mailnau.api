package config

import (
	"context"
	"errors"
	"google.golang.org/appengine/log"
	"sync"

	"github.com/fsnotify/fsnotify"
	_viper "github.com/spf13/viper"
)

const (
	rootKey = "data."
)

type Viper interface {
	Config

	Gray() Config
	Certs() Config
}

type viper struct {
	viper *_viper.Viper
	*sync.Mutex
}

func NewViper() (Config, error) {
	var configFileName, configFolder string
	ctx := context.Background()
	v := _viper.New()
	v.SetConfigType("json")

	configFileName = ".env"
	configFolder = "."

	v.SetConfigName(configFileName)
	v.AddConfigPath(configFolder)

	// Read config
	if err := v.ReadInConfig(); err != nil {
		var errNotFound *_viper.ConfigFileNotFoundError
		if ok := errors.As(err, &errNotFound); !ok {
			return nil, errNotFound
		}
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Infof(ctx, "vault file changed: %s", e.Name)
	})

	return &viper{
		v,
		&sync.Mutex{},
	}, nil
}

func (v *viper) GetInt(key string) int64 {
	return v.viper.GetInt64(rootKey + key)
}

func (v *viper) GetString(key string) string {
	return v.viper.GetString(rootKey + key)
}

func (v *viper) GetFloat64(key string) float64 {
	return v.viper.GetFloat64(rootKey + key)
}

func (v *viper) GetBool(key string) bool {
	return v.viper.GetBool(rootKey + key)
}
