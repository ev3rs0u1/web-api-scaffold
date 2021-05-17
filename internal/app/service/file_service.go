package service

import (
	"context"
	"web-api-scaffold/internal/app/model"
	"web-api-scaffold/internal/app/repository"
	"web-api-scaffold/internal/pkg/errno"
	"web-api-scaffold/internal/pkg/fileutil"
	"web-api-scaffold/internal/pkg/hasher"
)

type FileService interface {
	WithContext(ctx context.Context) FileService
	InitFileInfo(uid uint32, name, hash string) (*model.File, error)
}

type fileServiceImpl struct {
	FileRepository repository.FileRepository
}

func (service *fileServiceImpl) InitFileInfo(uid uint32, name, hash string) (*model.File, error) {
	if !service.FileRepository.IsHashOwner(uid, hash, model.FileTypeFolder) {
		return nil, errno.CodeFolderIdNotFound
	}

	if service.FileRepository.HasNameInFolder(uid, name, hash, model.FileTypeFile) {
		name = fileutil.NewCopyName(name)
	}

	file := &model.File{
		UID:  uid,
		PID:  service.FileRepository.GetFolderID(hash),
		Hash: hasher.GenerateUniqueHexUUID(),
		Name: name,
		Type: model.FileTypeFile,
	}

	return service.FileRepository.Create(file)
}

func NewFileService(fileRepository repository.FileRepository) FileService {
	return &fileServiceImpl{fileRepository}
}

func (service *fileServiceImpl) WithContext(ctx context.Context) FileService {
	service.FileRepository = service.FileRepository.WithContext(ctx)
	return service
}
