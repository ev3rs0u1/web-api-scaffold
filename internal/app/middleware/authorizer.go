package middleware

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"web-api-scaffold/internal/app/model"
	"web-api-scaffold/internal/pkg/constant"
	"web-api-scaffold/internal/pkg/errno"
	"web-api-scaffold/internal/pkg/responder"
)

func Authorizer(database *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) (err error) {
		var token string
		if token = ctx.Get(constant.HeaderKeyToken); token == "" {
			if token = ctx.Query(constant.HeaderKeyToken); token == "" {
				return ctx.JSON(responder.BuildBody(errno.CodeTokenNotRequired))
			}
		}

		var user model.User
		if database.
			WithContext(ctx.Context()).
			Where("token = ?", token).
			First(&user).RowsAffected == 0 {
			return ctx.JSON(responder.BuildBody(errno.CodeTokenNotRequired))
		}

		ctx.Locals(constant.CtxKeyUserID, user.ID)

		return ctx.Next()
	}
}
