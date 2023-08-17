package middleware

import (
	"time"

	"github.com/labstack/echo/v4"

	"github.com/borysbabii/Go-Clean-Architecture-REST-API/pkg/utils"
)

// Request logger middleware
func (mw *MiddlewareManager) RequestLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		start := time.Now()
		err := next(ctx)

		req := ctx.Request()
		res := ctx.Response()
		status := res.Status
		size := res.Size
		s := time.Since(start).String()
		requestID := utils.GetRequestID(ctx)
		clientIP := ctx.RealIP()
		browser := req.UserAgent()

		mw.logger.Infof("RequestID: %s, Method: %s, URI: %s, Status: %v, Size: %v, Time: %s, ClientIP: %s, Browser: %s",
			requestID, req.Method, req.URL, status, size, s, clientIP, browser,
		)
		return err
	}
}
