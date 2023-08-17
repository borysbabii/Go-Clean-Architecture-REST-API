package http

import (
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/currencies"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"

	"github.com/borysbabii/Go-Clean-Architecture-REST-API/config"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/pkg/httpErrors"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/pkg/logger"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/pkg/utils"
)

// Currency handlers
type currenciesHandlers struct {
	cfg        *config.Config
	currencyUC currencies.UseCase
	logger     logger.Logger
}

// NewCurrenciesHandlers Currency handlers constructor
func NewCurrenciesHandlers(cfg *config.Config, currencyUC currencies.UseCase, logger logger.Logger) currencies.Handlers {
	return &currenciesHandlers{cfg: cfg, currencyUC: currencyUC, logger: logger}
}

// GetCurrencies godoc
// @Summary Get all currencies
// @Description get all currencies
// @Tags Currency
// @Accept json
// @Produce json
// @Param page query int false "page number" Format(page)
// @Param size query int false "number of elements per page" Format(size)
// @Param orderBy query int false "filter name" Format(orderBy)
// @Success 200 {object} models.CurrenciesList
// @Router /currencies [get]
func (h currenciesHandlers) GetCurrencies() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "currenciesHandlers.GetCurrencies")
		defer span.Finish()

		pq, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		currencyList, err := h.currencyUC.GetCurrencies(ctx, pq)

		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, currencyList)
	}
}
