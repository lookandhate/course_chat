package config

import "fmt"

// GRPCConfig config for GRPC server.
type GRPCConfig struct {
	Port int `yaml:"port" env-default:"50052" env:"GRPC_PORT"`
}

func (c GRPCConfig) Address() string {
	return fmt.Sprintf(":%d", c.Port)
}
