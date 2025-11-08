package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.concur.com/I573758/example-golang-webapi/internal/core/utils"
)

func EasterEggMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			phrases := []string{
				"follow the white rabbit, Neo",
				"Believe you can and you're halfway there.",
				"Your only limit is your mind.",
				"Don't watch the clock; do what it does. Keep going.",
				"Success is not final, failure is not fatal: It is the courage to continue that counts.",
				"The harder you work for something, the greater you'll feel when you achieve it.",
				"Dream it. Wish it. Do it.",
				"Success doesn't just find you. You have to go out and get it.",
				"Don't stop when you're tired. Stop when you're done.",
				"Great things never come from comfort zones.",
				"Dream big and dare to fail.",
			}

			phrase := phrases[utils.RandomInt(0, len(phrases))]

			c.Response().Header().Set("X-Easter-Egg", phrase)
			return next(c)
		}
	}
}
