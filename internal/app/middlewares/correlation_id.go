package middlewares

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.concur.com/I573758/example-golang-webapi/internal/app/helpers/constants"
	"github.concur.com/I573758/example-golang-webapi/internal/core/utils"
)

const correlationIdHeaderKey = "X-Correlation-Id"

func CorrelationId() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			correlationId := c.Request().Header.Get(correlationIdHeaderKey)
			if correlationId == "" {
				year := time.Now().Year()
				month := time.Now().Month()
				day := time.Now().Day()
				correlationId = fmt.Sprintf("%d%02d%02d-%s", year, month, day, utils.Ulid())
			}

			ctx := c.Request().Context()
			ctx = context.WithValue(ctx, constants.CorrelationIdKey{}, correlationId)

			c.SetRequest(c.Request().WithContext(ctx))

			c.Response().Header().Set(correlationIdHeaderKey, correlationId)
			return next(c)
		}
	}
}
