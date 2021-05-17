package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	uuid "github.com/satori/go.uuid"
	"web-api-scaffold/internal/pkg/constant"
)

func RequestID() fiber.Handler {
	return requestid.New(requestid.Config{
		Header: fiber.HeaderXRequestID,
		Generator: func() string {
			return uuid.NewV4().String()
		},
		ContextKey: constant.CtxKeyRequestID,
	})
}
