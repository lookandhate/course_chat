package config

import (
	"net"
	"strconv"
)

type SwaggerConfig struct {
	Port int    `yaml:"port" env-default:"8081" env:"SWAGGER_PORT"`
	Host string `yaml:"host" env-default:"localhost" env:"SWAGGER_HOST"`
}

func (h SwaggerConfig) Address() string {
	return net.JoinHostPort(h.Host, strconv.Itoa(h.Port))
}
