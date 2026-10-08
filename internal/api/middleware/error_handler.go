package middlewares

import "github.com/labstack/echo/v4"

func ErrorHandler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)

			if err != nil {
				c.Logger().Error(err)
				return c.JSON(500, map[string]interface{}{
					"message": "Internal server error",
					"error":   err.Error(),
				})
			}
			return nil
		}
	}
}
