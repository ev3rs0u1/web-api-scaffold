package logger

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	gl "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"io"
	"os"
	"path/filepath"
	"time"
	"web-api-scaffold/internal/pkg/config"
	"web-api-scaffold/internal/pkg/constant"
	"web-api-scaffold/internal/pkg/logger/lumberjack"
)

const dbLogMsgTag = "DATABASE"

type DatabaseLogger struct {
	log *Logger
}

func NewGORMLogger() gl.Interface {
	var out io.Writer = os.Stdout

	if config.IsReleaseLogMode() {
		out = &lumberjack.Logger{
			Filename:   config.Instance().Logfile.Database,
			MaxSize:    10,
			MaxBackups: 20,
			MaxAge:     30,
			Compress:   false,
		}
	}

	return &DatabaseLogger{
		log: New(out),
	}
}

func (l *DatabaseLogger) LogMode(gl.LogLevel) gl.Interface {
	return l
}

func (l *DatabaseLogger) Info(_ context.Context, msg string, data ...interface{}) {
	l.log.Info().
		Str("msg", msg).
		Interface("data", data).
		Msg(dbLogMsgTag)
}

func (l *DatabaseLogger) Warn(_ context.Context, msg string, data ...interface{}) {
	l.log.Warn().
		Str("msg", msg).
		Interface("data", data).
		Msg(dbLogMsgTag)
}

func (l *DatabaseLogger) Error(_ context.Context, msg string, data ...interface{}) {
}

func (l *DatabaseLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	slowThreshold := 200 * time.Millisecond
	elapsed := time.Since(begin).Round(time.Millisecond)

	sql, rows := fc()
	if rows == -1 {
		rows = 0
	}

	caller := utils.FileWithLineNum()
	if len(caller) > 0 {
		if cwd, err := os.Getwd(); err == nil {
			if rel, err := filepath.Rel(cwd, caller); err == nil {
				caller = fmt.Sprintf("< %s >", rel)
			}
		}
	}

	if tid, ok := ctx.Value(constant.CtxKeyRequestID).(string); ok {
		switch {
		case err != nil && (!errors.Is(err, gorm.ErrRecordNotFound)):
			l.log.Err(err).
				Str("trace-id", tid).
				Str("statement", sql).
				Str("latency", elapsed.String()).
				Int64("rows", rows).
				Str("caller-src", caller).
				Msg(dbLogMsgTag)

		case elapsed > slowThreshold:
			l.log.Warn().
				Str("trace-id", tid).
				Str("statement", sql).
				Str("latency", fmt.Sprintf("SLOW SQL (%s) >= %s",
					elapsed.String(), slowThreshold.String())).
				Int64("rows", rows).
				Str("caller-src", caller).
				Msg(dbLogMsgTag)

		default:
			l.log.Trace().
				Str("trace-id", tid).
				Str("statement", sql).
				Str("latency", elapsed.String()).
				Int64("rows", rows).
				Msg(dbLogMsgTag)
		}
	}
}
