package config

type Config interface {
	GetInt(key string) int64
	GetString(key string) string
	GetFloat64(key string) float64
	GetBool(key string) bool
}
