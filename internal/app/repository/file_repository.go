package repository

import (
	"context"
	"gorm.io/gorm"
	"web-api-scaffold/internal/app/model"
)

type FileRepository interface {
	WithContext(ctx context.Context) FileRepository
	Create(file *model.File) (*model.File, error)
	GetFolderID(hash string) (id uint32)
	IsHashOwner(uid uint32, hash string, typ model.FileType) bool
	HasNameInFolder(uid uint32, name, hash string, types ...model.FileType) bool
}

type fileRepositoryImpl struct {
	database *gorm.DB
}

func NewFileRepository(database *gorm.DB) FileRepository {
	return &fileRepositoryImpl{database}
}

func (repo *fileRepositoryImpl) Create(file *model.File) (*model.File, error) {
	if err := repo.database.Create(file).Error; err != nil {
		return nil, err
	}
	return file, nil
}

func (repo *fileRepositoryImpl) GetFolderID(hash string) (id uint32) {
	if hash != "" {
		repo.database.
			Model(&model.File{}).
			Select("id").
			Where("hash = ? AND type = ?", hash, model.FileTypeFolder).
			Scan(&id)
	}
	return
}

func (repo *fileRepositoryImpl) HasNameInFolder(uid uint32, name, hash string, types ...model.FileType) bool {
	return repo.database.
		Where("uid = ? AND pid = ? AND name = ? AND type IN (?)",
			uid, repo.GetFolderID(hash), name, types).
		First(&model.File{}).RowsAffected == 1
}

func (repo *fileRepositoryImpl) IsHashOwner(uid uint32, hash string, typ model.FileType) bool {
	return repo.database.
		Where("uid = ? AND hash = ? AND type = ?",
			uid, hash, typ).
		First(&model.File{}).RowsAffected == 1 || (typ == model.FileTypeFolder && hash == "")
}

func (repo *fileRepositoryImpl) WithContext(ctx context.Context) FileRepository {
	repo.database = repo.database.WithContext(ctx)
	return repo
}
