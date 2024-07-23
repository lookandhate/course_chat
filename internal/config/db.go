package config

import "fmt"

// DatabaseConfig base interface for database configs with DSN retriever.
type DatabaseConfig interface {
	GetDSN() string
}

// PostgresConfig config for Postgres Database.
type PostgresConfig struct {
	Host     string `yaml:"host" env:"DB_HOST" env-default:"localhost"`
	Port     int    `yaml:"port" env:"POSTGRES_PORT" env-default:"5432"`
	User     string `yaml:"user" env:"POSTGRES_USER" env-default:"postgres"`
	DBName   string `yaml:"db_name" env:"POSTGRES_DB_NAME" env-default:"postgres"`
	Password string `yaml:"password" env:"POSTGRES_PASSWORD"`
}

// GetDSN return PG connection dsn.
func (db *PostgresConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.DBName)
}
