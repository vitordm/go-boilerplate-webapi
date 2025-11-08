package logger

import (
	"context"
	"log/slog"
)

type BeforeHandleFunc = func(context.Context) []slog.Attr

type Handler struct {
	handler      slog.Handler
	beforeHandle BeforeHandleFunc
}

func NewHandler(subHandler slog.Handler, beforeHandle BeforeHandleFunc) slog.Handler {
	return Handler{
		handler:      subHandler,
		beforeHandle: beforeHandle,
	}
}

func (h Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h Handler) Handle(ctx context.Context, record slog.Record) error {

	if h.beforeHandle != nil {
		extraLogs := h.beforeHandle(ctx)
		for _, log := range extraLogs {
			record.AddAttrs(log)
		}
	}
	return h.handler.Handle(ctx, record)
}

func (h Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewHandler(h.handler.WithAttrs(attrs), h.beforeHandle)
}

func (h Handler) WithGroup(name string) slog.Handler {
	return h.handler.WithGroup(name)
}
