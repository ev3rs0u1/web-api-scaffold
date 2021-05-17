package handler

import (
	"github.com/gofiber/fiber/v2"
	"web-api-scaffold/internal/app/model"
	"web-api-scaffold/internal/app/service"
	"web-api-scaffold/internal/pkg/errno"
	"web-api-scaffold/internal/pkg/requestor"
	"web-api-scaffold/internal/pkg/responder"
)

type FileInitHandler struct {
	input struct {
		Name       string `form:"name"        json:"name"        validate:"min=1,max=128"`
		FolderHash string `form:"folder_hash" json:"folder_hash" validate:"omitempty"`
	}
	FileService service.FileService
}

func NewFileInitHandler(fileService service.FileService) *FileInitHandler {
	return &FileInitHandler{FileService: fileService}
}

func (h *FileInitHandler) Input() interface{} {
	return &h.input
}

func (h *FileInitHandler) Validate() (err error) {
	return
}

func (h *FileInitHandler) Handler(ctx *fiber.Ctx) error {
	var err error
	var file *model.File

	if file, err = h.FileService.
		WithContext(ctx.Context()).
		InitFileInfo(requestor.GetUserID(ctx), h.input.Name, h.input.FolderHash); err != nil {
		return responder.BuildBody(errno.CodeFileInitializeError).With(err)
	}

	return responder.Succeed().With(file)
}
