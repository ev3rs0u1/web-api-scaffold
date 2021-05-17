package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"strings"
	"time"
	"web-api-scaffold/internal/pkg/config"
)

const (
	colorBlack = iota + 30
	colorRed
	colorGreen
	colorYellow
	colorBlue
	colorMagenta
	colorCyan
	colorWhite

	colorBold     = 1
	colorDarkGray = 90
)

type Logger struct {
	*zerolog.Logger
}

func New(out io.Writer) *Logger {
	writer := zerolog.ConsoleWriter{Out: out}

	colorize := func(s interface{}, c int, bold bool) string {
		if config.IsReleaseLogMode() {
			return fmt.Sprintf("%v", s)
		}

		if bold {
			return fmt.Sprintf("\u001B[%dm\x1b[%dm%v\x1b[0m\u001B[0m",
				colorBold, c, s)
		}

		return fmt.Sprintf("\x1b[%dm%v\x1b[0m", c, s)
	}

	writer.FormatTimestamp = func(i interface{}) string {
		t := time.Now().Format("2006-01-02 15:04:05.000")
		return colorize(t, colorDarkGray, false)
	}

	writer.FormatLevel = func(i interface{}) string {
		var l string

		if ll, ok := i.(string); ok {
			switch ll {
			case "trace":
				l = colorize("[TRC]", colorMagenta, false)
			case "debug":
				l = colorize("[DBG]", colorYellow, false)
			case "info":
				l = colorize("[INF]", colorGreen, false)
			case "warn":
				l = colorize("[WRN]", colorRed, false)
			case "error":
				l = colorize("[ERR]", colorRed, true)
			case "fatal":
				l = colorize("[FTL]", colorRed, true)
			case "panic":
				l = colorize("[PNC]", colorRed, true)
			default:
				l = colorize("[???]", colorBold, true)
			}
		} else {
			if i == nil {
				l = colorize("[???]", colorBold, true)
			} else {
				l = strings.ToUpper(fmt.Sprintf("[%s]", i))[0:3]
			}
		}

		return l
	}

	writer.FormatCaller = func(i interface{}) string {
		return ""
	}

	writer.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("(%s)",
			colorize(fmt.Sprintf("%s", i), colorWhite, true))
	}

	writer.FormatFieldName = func(i interface{}) string {
		return colorize(fmt.Sprintf("%s", i), colorCyan, false) +
			colorize("=", colorBold, true)
	}

	writer.FormatFieldValue = func(i interface{}) string {
		return colorize(fmt.Sprintf("%s", i), colorWhite, false)
	}

	writer.FormatErrFieldName = func(i interface{}) string {
		return colorize(fmt.Sprintf("%s", i), colorRed, false) +
			colorize("=", colorBold, true)
	}

	writer.FormatErrFieldValue = func(i interface{}) string {
		return colorize(fmt.Sprintf("%s", i), colorRed, true)
	}

	logger := zerolog.New(writer).
		With().
		Timestamp().
		Logger()

	return &Logger{&logger}
}

//func Instance() *Logger {
//	//jack :=lumberjack.Logger{
//	//	Filename:   config.Instance().Logfile.Server,
//	//	MaxSize:    7,
//	//	MaxBackups: 30,
//	//	MaxAge:     30,
//	//	Compress:   false,
//	//}
//	//fmt.Println("once +++++++++++++++++++++++++++++++++++++")
//	once.Do(func() {
//		fmt.Println("once +++++++++++++++++++++++++++++++++++++")
//		log = New(os.Stdout)
//	})
//	return log
//}
