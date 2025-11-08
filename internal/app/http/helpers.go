package http

import (
	"errors"
	"github.concur.com/I573758/example-golang-webapi/internal/app/helpers"
	"github.concur.com/I573758/example-golang-webapi/internal/core/server"
	"github.concur.com/I573758/example-golang-webapi/internal/core/utils"
	"net/http"
)

type DefaultResponse struct {
	Message string `json:"message"`
}

func NewDefaultResponse(message string) DefaultResponse {
	return DefaultResponse{Message: message}
}

func stringResponses(messages ...string) []DefaultResponse {
	var response []DefaultResponse
	for _, message := range messages {
		response = append(response, DefaultResponse{Message: message})
	}
	return response
}

// NotFound returns a 404 Not Found response
func NotFound(ctx server.Context) error {
	return ctx.JSON(http.StatusNotFound, stringResponses("Not found"))
}

// InternalServerError returns a 500 Internal Server Error response
func InternalServerError(ctx server.Context, err error) error {
	return ctx.JSON(http.StatusInternalServerError, stringResponses(err.Error()))
}

// BadRequest returns a 400 Bad Request response
func BadRequest[T any](ctx server.Context, data T) error {
	if err, ok := utils.IsError(data); ok {
		return ctx.JSON(http.StatusBadRequest, stringResponses(err.Error()))
	}
	return ctx.JSON(http.StatusBadRequest, data)
}

// Created returns a 201 Created response
func Created[T any](ctx server.Context, data T) error {
	return ctx.JSON(http.StatusCreated, data)
}

// NoContent returns a 204 No Content response
func NoContent(ctx server.Context) error {
	return ctx.NoContent(http.StatusNoContent)
}

func Ok[T any](ctx server.Context, data T) error {
	return ctx.JSON(http.StatusOK, data)
}

// Unauthorized returns a 401 Unauthorized response
func Unauthorized(ctx server.Context) error {
	return ctx.JSON(http.StatusUnauthorized, stringResponses("Unauthorized"))
}

// Forbidden returns a 403 Forbidden response
func Forbidden(ctx server.Context) error {
	return ctx.JSON(http.StatusForbidden, stringResponses("Forbidden"))
}

func NotModified(ctx server.Context) error {
	return ctx.JSON(http.StatusNotModified, stringResponses("Not modified"))
}

func ResolveError(ctx server.Context, err error) error {
	if err == nil {
		return InternalServerError(ctx, err)
	}

	switch {
	case errors.Is(err, helpers.EntityNotFoundError()):
		return NotFound(ctx)
	default:
		return InternalServerError(ctx, err)
	}
}
