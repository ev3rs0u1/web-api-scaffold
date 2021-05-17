package handler

import (
	"github.com/gofiber/fiber/v2"
	"web-api-scaffold/internal/app/model"
	"web-api-scaffold/internal/app/service"
	"web-api-scaffold/internal/pkg/errno"
	"web-api-scaffold/internal/pkg/responder"
	"web-api-scaffold/internal/pkg/validutil"
)

type DeviceListHandler struct {
	input struct {
		Token     string `query:"token" validate:"len=32"`
		Signature string `query:"sign"  validate:"len=32"`
	}
	UserService service.UserService
}

func NewDeviceListHandler(userService service.UserService) *DeviceListHandler {
	return &DeviceListHandler{UserService: userService}
}

func (h *DeviceListHandler) Input() interface{} {
	return &h.input
}

func (h *DeviceListHandler) Validate() (err error) {
	if !validutil.ValidateTokenSignature(h.input.Token, h.input.Signature) {
		err = errno.CodeInvalidSignature
	}
	return
}

func (h *DeviceListHandler) Handler(ctx *fiber.Ctx) error {
	var err error
	var users []*model.User

	if users, err = h.UserService.
		WithContext(ctx.Context()).
		FetchAllUsersUsedSpace(); err != nil {
		return err
	}

	return responder.Succeed().With(users)
}
