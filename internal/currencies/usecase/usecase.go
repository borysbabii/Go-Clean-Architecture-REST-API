package usecase

import (
	"context"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/config"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/currencies"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/models"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/pkg/logger"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/pkg/utils"
	"github.com/opentracing/opentracing-go"
)

const (
	basePrefix    = "api-currency:"
	cacheDuration = 3600
)

// Currency UseCase
type currencyUC struct {
	cfg          *config.Config
	currencyRepo currencies.Repository
	logger       logger.Logger
}

// NewCurrenciesUseCase Currencies UseCase constructor
func NewCurrenciesUseCase(cfg *config.Config, currencyRepo currencies.Repository, logger logger.Logger) currencies.UseCase {
	return &currencyUC{cfg: cfg, currencyRepo: currencyRepo, logger: logger}
}

// GetCurrencies returns list of currencies
func (u *currencyUC) GetCurrencies(ctx context.Context, query *utils.PaginationQuery) (*models.CurrenciesList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "currencyUC.GetCurrencies")
	defer span.Finish()

	return u.currencyRepo.GetCurrencies(ctx, query)
}
