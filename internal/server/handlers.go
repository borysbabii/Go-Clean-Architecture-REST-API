package server

import (
	"net/http"
	"strings"

	"github.com/borysbabii/Go-Clean-Architecture-REST-API/docs"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/pkg/csrf"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	// _ "github.com/borysbabii/Go-Clean-Architecture-REST-API/docs"
	authHttp "github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/auth/delivery/http"
	authRepository "github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/auth/repository"
	authUseCase "github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/auth/usecase"
	commentsHttp "github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/comments/delivery/http"
	commentsRepository "github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/comments/repository"
	commentsUseCase "github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/comments/usecase"
	currenciesHttp "github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/currencies/delivery/http"
	currenciesRepository "github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/currencies/repository"
	currenciesUseCase "github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/currencies/usecase"
	apiMiddlewares "github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/middleware"
	newsHttp "github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/news/delivery/http"
	newsRepository "github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/news/repository"
	newsUseCase "github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/news/usecase"
	sessionRepository "github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/session/repository"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/session/usecase"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/pkg/metric"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/pkg/utils"
)

// Map Server Handlers
func (s *Server) MapHandlers(e *echo.Echo) error {
	metrics, err := metric.CreateMetrics(s.cfg.Metrics.URL, s.cfg.Metrics.ServiceName)
	if err != nil {
		s.logger.Errorf("CreateMetrics Error: %s", err)
	}
	s.logger.Info(
		"Metrics available URL: %s, ServiceName: %s",
		s.cfg.Metrics.URL,
		s.cfg.Metrics.ServiceName,
	)

	// Init repositories
	aRepo := authRepository.NewAuthRepository(s.db)
	nRepo := newsRepository.NewNewsRepository(s.db)
	cRepo := commentsRepository.NewCommentsRepository(s.db)
	curRepo := currenciesRepository.NewCurrenciesRepository(s.db)
	sRepo := sessionRepository.NewSessionRepository(s.redisClient, s.cfg)
	aAWSRepo := authRepository.NewAuthAWSRepository(s.awsClient)
	authRedisRepo := authRepository.NewAuthRedisRepo(s.redisClient)
	newsRedisRepo := newsRepository.NewNewsRedisRepo(s.redisClient)

	// Init useCases
	authUC := authUseCase.NewAuthUseCase(s.cfg, aRepo, authRedisRepo, aAWSRepo, s.logger)
	newsUC := newsUseCase.NewNewsUseCase(s.cfg, nRepo, newsRedisRepo, s.logger)
	commUC := commentsUseCase.NewCommentsUseCase(s.cfg, cRepo, s.logger)
	curUC := currenciesUseCase.NewCurrenciesUseCase(s.cfg, curRepo, s.logger)
	sessUC := usecase.NewSessionUseCase(sRepo, s.cfg)

	// Init handlers
	authHandlers := authHttp.NewAuthHandlers(s.cfg, authUC, sessUC, s.logger)
	newsHandlers := newsHttp.NewNewsHandlers(s.cfg, newsUC, s.logger)
	commHandlers := commentsHttp.NewCommentsHandlers(s.cfg, commUC, s.logger)
	curHandlers := currenciesHttp.NewCurrenciesHandlers(s.cfg, curUC, s.logger)

	mw := apiMiddlewares.NewMiddlewareManager(sessUC, authUC, s.cfg, []string{"*"}, s.logger)

	e.Use(mw.RequestLoggerMiddleware)

	docs.SwaggerInfo.Title = "Go example REST API"
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	if s.cfg.Server.SSL {
		e.Pre(middleware.HTTPSRedirect())
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestID, csrf.CSRFHeader},
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10, // 1 KB
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	e.Use(middleware.RequestID())
	e.Use(mw.MetricsMiddleware(metrics))

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("2M"))
	if s.cfg.Server.Debug {
		e.Use(mw.DebugMiddleware)
	}

	v1 := e.Group("/api/v1")

	health := v1.Group("/health")
	authGroup := v1.Group("/auth")
	newsGroup := v1.Group("/news")
	commGroup := v1.Group("/comments")
	currenciesGroup := v1.Group("/currencies")

	authHttp.MapAuthRoutes(authGroup, authHandlers, mw)
	newsHttp.MapNewsRoutes(newsGroup, newsHandlers, mw)
	commentsHttp.MapCommentsRoutes(commGroup, commHandlers, mw)
	currenciesHttp.MapCurrenciesRoutes(currenciesGroup, curHandlers, mw)

	health.GET("", func(c echo.Context) error {
		s.logger.Infof("Health check RequestID: %s", utils.GetRequestID(c))
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	return nil
}
