package config

import (
	"net"
	"strconv"
)

type HTTPConfig struct {
	Port int    `yaml:"port" env-default:"8081" env:"HTTP_PORT"`
	Host string `yaml:"host" env-default:"localhost" env:"HTTP_HOST"`
}

func (h HTTPConfig) Address() string {
	return net.JoinHostPort(h.Host, strconv.Itoa(h.Port))
}
