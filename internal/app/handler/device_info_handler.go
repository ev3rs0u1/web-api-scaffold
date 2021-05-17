package handler

import (
	"github.com/gofiber/fiber/v2"
	"web-api-scaffold/internal/pkg/devio"
	"web-api-scaffold/internal/pkg/responder"
)

type DeviceInfoHandler struct {
	input struct{}
}

func (h *DeviceInfoHandler) Validate() (err error) {
	return
}

func NewDeviceInfoHandler() *DeviceInfoHandler {
	return &DeviceInfoHandler{}
}

func (h *DeviceInfoHandler) Input() interface{} {
	return &h.input
}

func (h *DeviceInfoHandler) Handler(_ *fiber.Ctx) error {
	return responder.Succeed().With(devio.Information())
}
