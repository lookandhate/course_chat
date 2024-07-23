package config

// GRPCConfig config for GRPC server.
type GRPCConfig struct {
	Port int `yaml:"port" env-default:"50052" env:"GRPC_PORT"`
}
