package config

import (
	"github.com/kialkuz/task-manager/internal/api"
	"github.com/kialkuz/task-manager/pkg/logger"
)

type ServerConfig struct {
	Logger logger.LoggerInterface
	Routes []api.RouteRegister
}
