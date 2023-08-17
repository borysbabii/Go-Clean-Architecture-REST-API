package currencies

import "github.com/labstack/echo/v4"

// Handlers Currencies HTTP handlers interface
type Handlers interface {
	GetCurrencies() echo.HandlerFunc
}
