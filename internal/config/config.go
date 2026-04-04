package config

import (
	"fmt"

	"github.com/kialkuz/task-manager/internal/infrastructure/env"
)

type Config struct {
	DB   DBConfig
	HTTP HTTPConfig
}

func Load() *Config {
	dbURL := fmt.Sprintf("%s://%s:%s@db:%s/%s?sslmode=%s",
		env.GetEnv("DB_TYPE", ""),
		env.GetEnv("DB_USER", ""),
		env.GetEnv("DB_PASSWORD", ""),
		env.GetEnv("DB_PORT", ""),
		env.GetEnv("DB_NAME", ""),
		env.GetEnv("DB_SSLMODE", ""),
	)

	return &Config{
		DB: DBConfig{
			DatabaseURI: dbURL,
			DBType:      env.GetEnv("DB_TYPE", ""),
		},
		HTTP: HTTPConfig{
			Port: env.GetEnv("SERVER_PORT", ""),
		},
	}
}
