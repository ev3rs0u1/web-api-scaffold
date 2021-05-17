package requestor

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"web-api-scaffold/internal/pkg/constant"
)

func GetUserID(ctx *fiber.Ctx) uint32 {
	if v, ok := ctx.Locals(constant.CtxKeyUserID).(string); ok {
		if id, err := strconv.ParseUint(v, 10, 64); err == nil {
			return uint32(id)
		}
	}
	return 0
}
