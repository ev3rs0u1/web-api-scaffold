package handler

import (
	"github.com/gofiber/fiber/v2"
	"web-api-scaffold/internal/app/model"
	"web-api-scaffold/internal/app/service"
	"web-api-scaffold/internal/pkg/errno"
	"web-api-scaffold/internal/pkg/responder"
	"web-api-scaffold/internal/pkg/validutil"
)

type DeviceBindHandler struct {
	input struct {
		Token     string `form:"token" json:"token" validate:"len=32"`
		Signature string `form:"sign"  json:"sign"  validate:"len=32"`
	}
	UserService service.UserService
}

func NewDeviceBindHandler(userService service.UserService) *DeviceBindHandler {
	return &DeviceBindHandler{UserService: userService}
}

func (h *DeviceBindHandler) Input() interface{} {
	return &h.input
}

func (h *DeviceBindHandler) Validate() (err error) {
	if !validutil.ValidateTokenSignature(h.input.Token, h.input.Signature) {
		err = errno.CodeInvalidSignature
	}
	return
}

func (h *DeviceBindHandler) Handler(ctx *fiber.Ctx) error {
	var err error
	var user *model.User

	if user, err = h.UserService.
		WithContext(ctx.Context()).
		BindDevice(h.input.Token); err != nil {
		return err
	}

	return responder.Succeed().With(user)
}
