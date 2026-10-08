package middlewares

import (
	"context"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vitordm/go-boilerplate-webapi/internal/app/helpers/constants"
)

func Logger(logger *slog.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:        true,
		LogLatency:       true,
		LogURI:           true,
		LogError:         true,
		LogContentLength: true,
		HandleError:      true, // forwards error to the global error handler, so it can decide appropriate status code

		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			correlationId := c.Request().Context().Value(constants.CorrelationIdKey{}).(string)

			if v.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("correlation_id", correlationId),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
					slog.String("correlation_id", correlationId),
				)
			}
			return nil
		},
		Skipper: func(c echo.Context) bool {
			path := c.Request().URL.Path
			if path == "/health" {
				return true
			}
			return false
		},
	})
}
