package helpers

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path"
	"time"

	"github.com/vitordm/go-boilerplate-webapi/internal/app/helpers/constants"
	"github.com/vitordm/go-boilerplate-webapi/internal/core/logger"
	"github.com/vitordm/go-boilerplate-webapi/internal/core/utils"
)

func BeforeHandle(ctx context.Context) []slog.Attr {
	var result []slog.Attr

	correlationId := ctx.Value(constants.CorrelationIdKey{})
	if correlationId != nil && correlationId != "" {
		result = append(result, slog.String("correlation_id", correlationId.(string)))
	}
	return result
}

func NewSLogJsonCommandLine() *slog.Logger {
	handlerOptions := slog.HandlerOptions{
		AddSource: true,
	}
	handler := slog.NewJSONHandler(os.Stdout, &handlerOptions)
	l := slog.New(logger.NewHandler(handler, BeforeHandle))
	return l
}

func NewSLogJsonCommandLineAndFile() *slog.Logger {
	handlerOptions := slog.HandlerOptions{
		AddSource: true,
	}

	logPath := utils.GetEnvOrDefault(constants.ROOT_PATH_KEY, "/app")
	logPath = path.Join(logPath, "temp", "log")

	fileWriter := NewLogFileWriter(logPath)

	logOutput := io.MultiWriter(fileWriter, os.Stdout)

	handler := slog.NewJSONHandler(logOutput, &handlerOptions)
	l := slog.New(logger.NewHandler(handler, BeforeHandle))
	return l
}

func NewSLogTextCommandLine() *slog.Logger {
	handlerOptions := slog.HandlerOptions{
		//AddSource: true,
	}
	//handler := slog.NewJSONHandler(os.Stdout, &handlerOptions)
	handler := slog.NewTextHandler(os.Stdout, &handlerOptions)
	l := slog.New(logger.NewHandler(handler, BeforeHandle))
	return l
}

func BuildLogger() *slog.Logger {
	typeLogger := utils.GetEnvOrDefault(constants.LOGGER_TYPE_KEY, constants.LOGGER_TYPE_DEFAULT)

	if typeLogger == constants.LOGGER_TYPE_JSON {
		return NewSLogJsonCommandLine()
	}

	if typeLogger == constants.LOGGER_TYPE_FILE {
		return NewSLogJsonCommandLineAndFile()
	}

	return NewSLogTextCommandLine()

}

type LogFileWriter struct {
	logFolder string
}

func NewLogFileWriter(logFolder string) *LogFileWriter {
	return &LogFileWriter{
		logFolder: logFolder,
	}

}

func (w *LogFileWriter) Write(p []byte) (n int, err error) {

	filename := path.Join(w.logFolder, "log.log")

	if _, err := os.Stat(w.logFolder); os.IsNotExist(err) {
		os.MkdirAll(w.logFolder, os.ModePerm)
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	defer f.Close()

	n, err = f.WriteString(string(p))

	fileInfo, nErr := os.Stat(filename)
	if nErr != nil {
		return n, err
	}

	if fileInfo.Size() > (1024 * 1024 * 10) {
		newPath := fmt.Sprintf("log_%s.log", time.Now().Format("2006-01-02 15:04:05"))
		newPath = path.Join(w.logFolder, newPath)
		os.Rename(filename, newPath)
	}

	return n, err
}
