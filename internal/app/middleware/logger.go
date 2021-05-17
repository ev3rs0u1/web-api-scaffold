package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"time"
	"web-api-scaffold/internal/pkg/constant"
	"web-api-scaffold/internal/pkg/logger"
)

func Logger(l *logger.Logger) fiber.Handler {
	return func(ctx *fiber.Ctx) (err error) {
		begin := time.Now()

		if err = ctx.Next(); err != nil {
			ctx.Status(fiber.StatusInternalServerError)
		}

		end := time.Now()
		res := ctx.Response()
		event := &zerolog.Event{}

		switch {
		case res.StatusCode() == 200:
			event = l.Info()
		case res.StatusCode() >= 400:
			event = l.Warn()
		case res.StatusCode() >= 500:
			event = l.Error().Err(err)
		default:
			event = l.Trace()
		}

		latency := end.Sub(begin).Round(time.Millisecond)

		reqBody := []byte("(BODY TOO LARGE)")
		resBody := []byte("(BODY TOO LARGE)")

		if ctx.Request().Header.ContentLength() <= constant.MaxLogContentLength {
			reqBody = ctx.Request().Body()
		}

		if ctx.Response().Header.ContentLength() <= constant.MaxLogContentLength {
			resBody = ctx.Response().Body()
		}

		event.
			Str("method", ctx.Method()).
			Str("path", ctx.Path()).
			Str("remote-addr", ctx.IP()).
			Str("user-agent", ctx.Get(fiber.HeaderUserAgent)).
			Str("latency", latency.String()).
			Int("status", res.StatusCode()).
			Bytes("req-body", reqBody).
			Bytes("res-body", resBody).
			Interface("req-id", ctx.Locals(constant.CtxKeyRequestID)).
			Msg("REQUEST")

		return
	}
}
