package config

import "fmt"

// GRPCConfig config for GRPC server.
type GRPCConfig struct {
	Port int `yaml:"port" env-default:"50052" env:"GRPC_PORT"`
}

func (receiver GRPCConfig) Address() string {
	return fmt.Sprintf(":%d", receiver.Port)
}
