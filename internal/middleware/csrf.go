package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/borysbabii/Go-Clean-Architecture-REST-API/pkg/csrf"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/pkg/httpErrors"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/pkg/utils"
)

// CSRF Middleware
func (mw *MiddlewareManager) CSRF(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if !mw.cfg.Server.CSRF {
			return next(ctx)
		}

		token := ctx.Request().Header.Get(csrf.CSRFHeader)
		if token == "" {
			mw.logger.Errorf("CSRF Middleware get CSRF header, Token: %s, Error: %s, RequestId: %s",
				token,
				"empty CSRF token",
				utils.GetRequestID(ctx),
			)
			return ctx.JSON(http.StatusForbidden, httpErrors.NewRestError(http.StatusForbidden, "Invalid CSRF Token", "no CSRF Token"))
		}

		sid, ok := ctx.Get("sid").(string)
		if !csrf.ValidateToken(token, sid, mw.logger) || !ok {
			mw.logger.Errorf("CSRF Middleware csrf.ValidateToken Token: %s, Error: %s, RequestId: %s",
				token,
				"empty token",
				utils.GetRequestID(ctx),
			)
			return ctx.JSON(http.StatusForbidden, httpErrors.NewRestError(http.StatusForbidden, "Invalid CSRF Token", "no CSRF Token"))
		}

		return next(ctx)
	}
}
