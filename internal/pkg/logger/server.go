package logger

import (
	"io"
	"os"
	"web-api-scaffold/internal/pkg/config"
	"web-api-scaffold/internal/pkg/logger/lumberjack"
)

func NewServerLogger() *Logger {
	var out io.Writer = os.Stdout

	if config.IsReleaseLogMode() {
		out = &lumberjack.Logger{
			Filename:   config.Instance().Logfile.Server,
			MaxSize:    10,
			MaxBackups: 20,
			MaxAge:     30,
			Compress:   false,
		}
	}

	return New(out)
}
