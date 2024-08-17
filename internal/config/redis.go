package config

import (
	"net"
)

// RedisConfig config for redis client.
type RedisConfig struct {
	Host        string `yaml:"host" env:"REDIS_HOST" env-default:"localhost"`
	Port        string `yaml:"port" env:"REDIS_PORT" env-default:"63790"`
	MaxIdle     int    `yaml:"max_idle" env:"REDIS_MAX_IDLE" env-default:"5"`
	IdleTimeout int    `yaml:"idle_timeout " env:"REDIS_IDLE_TIMEOUT" env-default:"5"`
}

func (c RedisConfig) Address() string {
	return net.JoinHostPort(c.Host, c.Port)
}
