package logging

import (
	"context"
	"fmt"
	"github.com/vitordm/go-boilerplate-webapi/internal/shared/constants"
	"github.com/vitordm/go-boilerplate-webapi/internal/shared/utils"
	"io"
	"log/slog"
	"os"
	"path"
	"time"
)

func BeforeHandle(ctx context.Context) []slog.Attr {
	if correlationID, ok := ctx.Value(constants.CorrelationIdKey{}).(string); ok && correlationID != "" {
		return []slog.Attr{slog.String("correlation_id", correlationID)}
	}
	return nil
}

func BuildLogger() *slog.Logger {
	switch utils.GetEnvOrDefault(constants.LOGGER_TYPE_KEY, constants.LOGGER_TYPE_DEFAULT) {
	case constants.LOGGER_TYPE_JSON:
		return slog.New(NewHandler(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}), BeforeHandle))
	case constants.LOGGER_TYPE_FILE:
		logPath := path.Join(utils.GetEnvOrDefault(constants.ROOT_PATH_KEY, "/app"), "temp", "log")
		return slog.New(NewHandler(slog.NewJSONHandler(io.MultiWriter(NewLogFileWriter(logPath), os.Stdout), &slog.HandlerOptions{AddSource: true}), BeforeHandle))
	default:
		return slog.New(NewHandler(slog.NewTextHandler(os.Stdout, nil), BeforeHandle))
	}
}

type LogFileWriter struct{ logFolder string }

func NewLogFileWriter(logFolder string) *LogFileWriter { return &LogFileWriter{logFolder: logFolder} }
func (w *LogFileWriter) Write(p []byte) (int, error) {
	if err := os.MkdirAll(w.logFolder, os.ModePerm); err != nil {
		return 0, err
	}
	filename := path.Join(w.logFolder, "log.log")
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return 0, fmt.Errorf("open log file: %w", err)
	}
	defer f.Close()
	n, err := f.Write(p)
	if err != nil {
		return n, err
	}
	if info, statErr := os.Stat(filename); statErr == nil && info.Size() > 10*1024*1024 {
		_ = os.Rename(filename, path.Join(w.logFolder, fmt.Sprintf("log_%s.log", time.Now().Format("2006-01-02 15-04-05"))))
	}
	return n, nil
}
