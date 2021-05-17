package handler

import (
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"strings"
	"web-api-scaffold/internal/app/service"
	"web-api-scaffold/internal/pkg/constant"
	"web-api-scaffold/internal/pkg/errno"
	"web-api-scaffold/internal/pkg/hasher"
	"web-api-scaffold/internal/pkg/responder"
)

type FileCreateHandler struct {
	input struct {
		FileHash   string `form:"file_hash"   validate:"len=32"`
		ChunkHash  string `form:"chunk_hash"  validate:"len=64"`
		ChunkIndex uint32 `form:"chunk_index" validate:"min=1"`
	}
	FileService service.FileService
}

func NewFileCreateHandler(fileService service.FileService) *FileCreateHandler {
	return &FileCreateHandler{FileService: fileService}
}

func (h *FileCreateHandler) Input() interface{} {
	return &h.input
}

func (h *FileCreateHandler) Validate() (err error) {
	h.input.FileHash = strings.ToLower(h.input.FileHash)
	h.input.ChunkHash = strings.ToLower(h.input.ChunkHash)
	return
}

func (h *FileCreateHandler) Handler(ctx *fiber.Ctx) error {
	var err error
	var chunk *multipart.FileHeader

	if chunk, err = ctx.FormFile("chunk"); err != nil {
		return responder.BuildBody(errno.CodeChunkUploadError).With(err)
	}

	if chunk.Size == 0 || chunk.Size > constant.MaxUploadChunkSize {
		return errno.CodeFileSizeExceedLimit
	}

	var file multipart.File

	if file, err = chunk.Open(); err != nil {
		return err
	}
	defer file.Close()

	if h.input.ChunkHash != hasher.CalculateSHA256HashByReader(file) {
		return errno.CodeInvalidChunkHash
	}

	return responder.Succeed().With(ctx.Locals(constant.CtxKeyUserID))
}
