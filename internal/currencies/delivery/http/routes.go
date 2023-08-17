package http

import (
	"github.com/labstack/echo/v4"

	"github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/currencies"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/middleware"
)

// MapCurrenciesRoutes Map currency routes
func MapCurrenciesRoutes(currencyGroup *echo.Group, h currencies.Handlers, mw *middleware.MiddlewareManager) {
	currencyGroup.GET("", h.GetCurrencies())
}
