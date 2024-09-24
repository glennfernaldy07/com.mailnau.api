package config

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"google.golang.org/appengine/log"

	"github.com/fsnotify/fsnotify"
	infisical "github.com/infisical/go-sdk"
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

const (
	INFISICAL_CLIENT_ID     = "INFISICAL_CLIENT_ID"
	INFISICAL_CLIENT_SECRET = "INFISICAL_CLIENT_SECRET"
	INFISICAL_URL           = "INFISICAL_URL"
	INFISCAL_PROJECT_ID     = "INFISICAL_PROJECT_ID"
	INFISICAL_PATH          = "INFISICAL_PATH"
	INFISICAL_ENVIRONMENT   = "INFISICAL_ENVIRONMENT"
	INFISICAL_SECRET_KEY    = "INFISICAL_API_SECRET_KEY"
	// 5 minutes
	REFRESH_DURATION_INTERVAL = time.Minute * 5
	CONFIG_TYPE               = "json"
	CONFIG_FILE_NAME          = "config.json"
)

func fetchSecrets(env *_viper.Viper) bool {
	v := _viper.New()

	v.AutomaticEnv()

	client := infisical.NewInfisicalClient(infisical.Config{
		SiteUrl: v.GetString(INFISICAL_URL), // Optional, default is https://app.infisical.com
	})

	_, err := client.Auth().UniversalAuthLogin(v.GetString(INFISICAL_CLIENT_ID), v.GetString(INFISICAL_CLIENT_SECRET))

	if err != nil {
		fmt.Println("[ENV]	Authentication failed: ", err)
		return false
	}

	secret, err := client.Secrets().Retrieve(infisical.RetrieveSecretOptions{
		SecretKey:   v.GetString(INFISICAL_SECRET_KEY),
		Environment: v.GetString(INFISICAL_ENVIRONMENT),
		ProjectID:   v.GetString(INFISCAL_PROJECT_ID),
		SecretPath:  v.GetString(INFISICAL_PATH),
	})

	if err != nil {
		fmt.Println("[ENV]	Error: ", err)
		return false
	}

	env.SetConfigType(CONFIG_TYPE)
	err = env.ReadConfig(bytes.NewReader([]byte(secret.SecretValue)))

	return err == nil
}

func autoRefresh(env *viper) {
	for {
		env.Mutex.Lock()

		if fetchSecrets(env.viper) {
			fmt.Println("[ENV]	Secrets Refreshed")
		} else {
			fmt.Println("Secrets Refresh Failed")
		}

		env.Mutex.Unlock()

		time.Sleep(REFRESH_DURATION_INTERVAL)
	}
}

func NewViper() (Config, error) {
	var configFileName, configFolder string
	ctx := context.Background()
	v := _viper.New()

	if fetchSecrets(v) {
		fmt.Println("[ENV] Secrets Fetched, using secrets from vault")

		env := &viper{
			v,
			&sync.Mutex{},
		}

		go autoRefresh(env)

		return env, nil
	}

	fmt.Println("[ENV] Secrets Fetch Failed, fallback to local config")

	v.SetConfigType(CONFIG_TYPE)

	configFileName = "config.json"
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

	env := &viper{
		v,
		&sync.Mutex{},
	}

	return env, nil
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
