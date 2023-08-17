package middleware

import (
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/config"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/auth"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/session"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/pkg/logger"
)

// Middleware manager
type MiddlewareManager struct {
	sessUC  session.UCSession
	authUC  auth.UseCase
	cfg     *config.Config
	origins []string
	logger  logger.Logger
}

// Middleware manager constructor
func NewMiddlewareManager(sessUC session.UCSession, authUC auth.UseCase, cfg *config.Config, origins []string, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{sessUC: sessUC, authUC: authUC, cfg: cfg, origins: origins, logger: logger}
}
